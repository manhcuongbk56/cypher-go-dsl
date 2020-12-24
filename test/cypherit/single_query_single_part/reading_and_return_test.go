package single_query_single_part

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

var bikeNode = cypher.NewNode("Bike").NamedByString("b")
var userNode = cypher.NewNode("User").NamedByString("u")

func TestUnrelatedNode(t *testing.T) {
	statement, err := cypher.MatchElements(bikeNode, userNode, cypher.NewNode("U").NamedByString("o")).
		ReturningByNamed(bikeNode, userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
	}
	query := cypher.NewRenderer().Render(statement)
	expect := "MATCH (b:`Bike`), (u:`User`), (o:`U`) RETURN b, u"
	if query != expect {
		t.Errorf("%s is incorrect \n %s", query, expect)
	}
}

func TestAsteriskShouldWork(t *testing.T) {
	statement, err := cypher.MatchElements(bikeNode, userNode, cypher.NewNode("U").NamedByString("o")).
		Returning(cypher.CypherAsterisk()).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
	}
	query := cypher.NewRenderer().Render(statement)
	expect := "MATCH (b:`Bike`), (u:`User`), (o:`U`) RETURN *"
	if query != expect {
		t.Errorf("%s is incorrect \n %s", query, expect)
	}
}

func TestAliasedExpressionInReturn(t *testing.T) {
	unnamedNode := cypher.NewNode("ANode")
	namedNode := cypher.NewNode("AnotherNode").NamedByString("o")
	statement, err := cypher.MatchElements(unnamedNode, namedNode).
		Returning(namedNode.As("theOtherNode")).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
	}
	query := cypher.NewRenderer().Render(statement)
	expect := "MATCH (:`ANode`), (o:`AnotherNode`) RETURN o AS theOtherNode"
	if query != expect {
		t.Errorf("\n%s is incorrect \n%s", query, expect)
	}
}

func TestSimpleRelationship(t *testing.T) {
	statement, err := cypher.MatchElements(userNode.RelationshipTo(bikeNode, "OWNS")).
		ReturningByNamed(bikeNode, userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
	}
	query := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`)-[:`OWNS`]->(b:`Bike`) RETURN b, u"
	if query != expect {
		t.Errorf("\n%s is incorrect \n%s", query, expect)
	}
}

func TestMultipleRelationshipTypes(t *testing.T) {
	statement, err := cypher.MatchElements(userNode.RelationshipTo(bikeNode, "OWNS", "RIDES")).
		ReturningByNamed(bikeNode, userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
	}
	query := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`)-[:`OWNS`|`RIDES`]->(b:`Bike`) RETURN b, u"
	if query != expect {
		t.Errorf("\n%s is incorrect \n%s", query, expect)
	}
}

func TestRelationshipWithProperties(t *testing.T) {
	statement, err := cypher.MatchElements(userNode.RelationshipTo(bikeNode, "OWNS").WithProperties(cypher.MapOf("boughtOn", cypher.LiteralOf("2019-04-16")))).
		ReturningByNamed(bikeNode, userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
	}
	query := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`)-[:`OWNS` {boughtOn: '2019-04-16'}]->(b:`Bike`) RETURN b, u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestRelationshipWithMinimumLength(t *testing.T) {
	statement, err := cypher.MatchElements(userNode.RelationshipTo(bikeNode, "OWNS").Min(3)).
		ReturningByNamed(bikeNode, userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
	}
	query := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`)-[:`OWNS`*3..]->(b:`Bike`) RETURN b, u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}

}

func TestRelationshipWithMaximumLength(t *testing.T) {
	statement, err := cypher.MatchElements(userNode.RelationshipTo(bikeNode, "OWNS").Max(5)).
		ReturningByNamed(bikeNode, userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
	}
	query := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`)-[:`OWNS`*..5]->(b:`Bike`) RETURN b, u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestRelationshipWithLength(t *testing.T) {
	statement, err := cypher.MatchElements(userNode.RelationshipTo(bikeNode, "OWNS").Length(3, 5)).
		ReturningByNamed(bikeNode, userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
	}
	query := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`)-[:`OWNS`*3..5]->(b:`Bike`) RETURN b, u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestRelationshipWithLengthAndProperties(t *testing.T) {
	statement, err := cypher.MatchElements(userNode.RelationshipTo(bikeNode, "OWNS").Length(3, 5).WithProperties(cypher.MapOf("boughtOn", cypher.LiteralOf("2019-04-16")))).
		ReturningByNamed(bikeNode, userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
	}
	query := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`)-[:`OWNS`*3..5 {boughtOn: '2019-04-16'}]->(b:`Bike`) RETURN b, u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestSimpleRelationshipWithReturn(t *testing.T) {
	owns := userNode.RelationshipTo(bikeNode, "OWNS").NamedByString("o")
	statement, err := cypher.MatchElements(owns).
		ReturningByNamed(bikeNode, userNode, owns).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
	}
	query := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`)-[o:`OWNS`]->(b:`Bike`) RETURN b, u, o"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestChainedRelations1(t *testing.T) {
	tripNode := cypher.NewNode("Trip").NamedByString("t")
	statementBuilder := cypher.MatchElements(userNode.
		RelationshipTo(bikeNode, "OWNS").NamedByString("r1").
		RelationshipTo(tripNode, "USED_ON").NamedC("r2"))
	expression := cypher.ExpressionChain(userNode.Property("name")).MatchesPattern(".*aName").Get()
	statement, err := statementBuilder.Where(expression).ReturningByNamed(bikeNode, userNode).Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`)-[r1:`OWNS`]->(b:`Bike`)-[r2:`USED_ON`]->(t:`Trip`) WHERE u.name =~ '.*aName' RETURN b, u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestChainedRelations2(t *testing.T) {
	tripNode := cypher.NewNode("Trip").NamedByString("t")
	statementBuilder := cypher.MatchElements(userNode.
		RelationshipTo(bikeNode, "OWNS").
		RelationshipTo(tripNode, "USED_ON").NamedC("r2"))
	expression := cypher.ExpressionChain(userNode.Property("name")).MatchesPattern(".*aName").Get()
	statement, err := statementBuilder.Where(expression).ReturningByNamed(bikeNode, userNode).Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`)-[:`OWNS`]->(b:`Bike`)-[r2:`USED_ON`]->(t:`Trip`) WHERE u.name =~ '.*aName' RETURN b, u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestChainedRelations3(t *testing.T) {
	tripNode := cypher.NewNode("Trip").NamedByString("t")
	statementBuilder := cypher.MatchElements(userNode.
		RelationshipTo(bikeNode, "OWNS").
		RelationshipTo(tripNode, "USED_ON").NamedC("r2").
		RelationshipFrom(userNode, "WAS_ON").NamedC("x").
		RelationshipBetween(cypher.NewNode("SOMETHING")).NamedC("y"))
	expression := cypher.ExpressionChain(userNode.Property("name")).MatchesPattern(".*aName").Get()
	statement, err := statementBuilder.Where(expression).ReturningByNamed(bikeNode, userNode).Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`)-[:`OWNS`]->(b:`Bike`)-[r2:`USED_ON`]->(t:`Trip`)<-[x:`WAS_ON`]-(u)-[y]-(:`SOMETHING`) WHERE u.name =~ '.*aName' RETURN b, u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestChainedRelationshipWithPropertiesAndLength(t *testing.T) {
	tripNode := cypher.NewNode("Trip").NamedByString("t")
	statementBuilder := cypher.MatchElements(userNode.
		RelationshipTo(bikeNode, "OWNS").
		RelationshipTo(tripNode, "USED_ON").NamedC("r2").Min(1).Properties(cypher.MapOf("when", cypher.LiteralOf("2019-04-16"))).
		RelationshipFrom(userNode, "WAS_ON").NamedC("x").Max(2).Properties(cypher.MapOf("whatever", cypher.LiteralOf("2020-04-16"))).
		RelationshipBetween(cypher.NewNode("SOMETHING")).NamedC("y").Length(2, 3).Properties(cypher.MapOf("idk", cypher.LiteralOf("2021-04-16"))))
	expression := cypher.ExpressionChain(userNode.Property("name")).MatchesPattern(".*aName").Get()
	statement, err := statementBuilder.Where(expression).ReturningByNamed(bikeNode, userNode).Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`)-[:`OWNS`]->(b:`Bike`)-[r2:`USED_ON`*1.. {when: '2019-04-16'}]->(t:`Trip`)<-[x:`WAS_ON`*..2 {whatever: '2020-04-16'}]-(u)-[y*2..3 {idk: '2021-04-16'}]-(:`SOMETHING`) WHERE u.name =~ '.*aName' RETURN b, u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestSizeOfRelationship(t *testing.T) {
	statementBuilder := cypher.MatchElements(cypher.AnyNodeNamed("a"))
	expression := cypher.ExpressionChain(cypher.CypherProperty("a", "name")).IsEqualTo(cypher.LiteralOf("Alice")).Get()
	statement, err := statementBuilder.Where(expression).
		Returning(cypher.FunctionSizeByPattern(cypher.AnyNodeNamed("a").RelationshipTo(cypher.CypherAnyNode())).As("fof").Get()).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query := cypher.NewRenderer().Render(statement)
	expect := "MATCH (a) WHERE a.name = 'Alice' RETURN size((a)-->()) AS fof"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}
