package single_query_single_part

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

var bikeNode = cypher.ANode("Bike").NamedByString("b")
var userNode = cypher.ANode("User").NamedByString("u")

func TestUnrelatedNode(t *testing.T) {
	statement, err := cypher.Match(bikeNode, userNode, cypher.ANode("U").NamedByString("o")).
		ReturningByNamed(bikeNode, userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (b:`Bike`), (u:`User`), (o:`U`) RETURN b, u"
	if query != expect {
		t.Errorf("%s is incorrect \n %s", query, expect)
	}
}

func TestAsteriskShouldWork(t *testing.T) {
	statement, err := cypher.Match(bikeNode, userNode, cypher.ANode("U").NamedByString("o")).
		Returning(cypher.AnAsterisk()).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (b:`Bike`), (u:`User`), (o:`U`) RETURN *"
	if query != expect {
		t.Errorf("%s is incorrect \n %s", query, expect)
	}
}

func TestAliasedExpressionInReturn(t *testing.T) {
	unnamedNode := cypher.ANode("ANode").NamedByString("anode")
	namedNode := cypher.ANode("AnotherNode").NamedByString("o")
	statement, err := cypher.Match(unnamedNode, namedNode).
		Returning(namedNode.As("theOtherNode")).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (anode:`ANode`), (o:`AnotherNode`) RETURN o AS theOtherNode"
	if query != expect {
		t.Errorf("\n%s is incorrect \n%s", query, expect)
	}
}

func TestSimpleRelationship(t *testing.T) {
	statement, err := cypher.Match(userNode.RelationshipTo(bikeNode, "OWNS")).
		ReturningByNamed(bikeNode, userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`)-[:`OWNS`]->(b:`Bike`) RETURN b, u"
	if query != expect {
		t.Errorf("\n%s is incorrect \n%s", query, expect)
	}
}

func TestMultipleRelationshipTypes(t *testing.T) {
	statement, err := cypher.Match(userNode.RelationshipTo(bikeNode, "OWNS", "RIDES")).
		ReturningByNamed(bikeNode, userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`)-[:`OWNS`|`RIDES`]->(b:`Bike`) RETURN b, u"
	if query != expect {
		t.Errorf("\n%s is incorrect \n%s", query, expect)
	}
}

func TestRelationshipWithProperties(t *testing.T) {
	statement, err := cypher.
		Match(userNode.RelationshipTo(bikeNode, "OWNS").
			WithProperties(cypher.MapOf("boughtOn", cypher.LiteralOf("2019-04-16")))).
		ReturningByNamed(bikeNode, userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`)-[:`OWNS` {boughtOn: '2019-04-16'}]->(b:`Bike`) RETURN b, u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestRelationshipWithMinimumLength(t *testing.T) {
	statement, err := cypher.Match(userNode.RelationshipTo(bikeNode, "OWNS").Min(3)).
		ReturningByNamed(bikeNode, userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`)-[:`OWNS`*3..]->(b:`Bike`) RETURN b, u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}

}

func TestRelationshipWithMaximumLength(t *testing.T) {
	statement, err := cypher.Match(userNode.RelationshipTo(bikeNode, "OWNS").Max(5)).
		ReturningByNamed(bikeNode, userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`)-[:`OWNS`*..5]->(b:`Bike`) RETURN b, u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestRelationshipWithLength(t *testing.T) {
	statement, err := cypher.Match(userNode.RelationshipTo(bikeNode, "OWNS").Length(3, 5)).
		ReturningByNamed(bikeNode, userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`)-[:`OWNS`*3..5]->(b:`Bike`) RETURN b, u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestRelationshipWithLengthAndProperties(t *testing.T) {
	statement, err := cypher.Match(userNode.RelationshipTo(bikeNode, "OWNS").Length(3, 5).WithProperties(cypher.MapOf("boughtOn", cypher.LiteralOf("2019-04-16")))).
		ReturningByNamed(bikeNode, userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`)-[:`OWNS`*3..5 {boughtOn: '2019-04-16'}]->(b:`Bike`) RETURN b, u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestSimpleRelationshipWithReturn(t *testing.T) {
	owns := userNode.RelationshipTo(bikeNode, "OWNS").NamedByString("o")
	statement, err := cypher.Match(owns).
		ReturningByNamed(bikeNode, userNode, owns).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`)-[o:`OWNS`]->(b:`Bike`) RETURN b, u, o"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestChainedRelations1(t *testing.T) {
	tripNode := cypher.ANode("Trip").NamedByString("t")
	statementBuilder := cypher.Match(userNode.
		RelationshipTo(bikeNode, "OWNS").NamedByString("r1").
		RelationshipTo(tripNode, "USED_ON").NamedC("r2"))
	expression := userNode.Property("name").MatchesPattern(".*aName").Get()
	statement, err := statementBuilder.Where(expression).ReturningByNamed(bikeNode, userNode).Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`)-[r1:`OWNS`]->(b:`Bike`)-[r2:`USED_ON`]->(t:`Trip`) WHERE u.name =~ '.*aName' RETURN b, u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestChainedRelations2(t *testing.T) {
	tripNode := cypher.ANode("Trip").NamedByString("t")
	statementBuilder := cypher.Match(userNode.
		RelationshipTo(bikeNode, "OWNS").
		RelationshipTo(tripNode, "USED_ON").NamedC("r2"))
	expression := cypher.ExpressionWrap(userNode.Property("name")).MatchesPattern(".*aName").Get()
	statement, err := statementBuilder.Where(expression).ReturningByNamed(bikeNode, userNode).Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`)-[:`OWNS`]->(b:`Bike`)-[r2:`USED_ON`]->(t:`Trip`) WHERE u.name =~ '.*aName' RETURN b, u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestChainedRelations3(t *testing.T) {
	tripNode := cypher.ANode("Trip").NamedByString("t")
	statementBuilder := cypher.Match(userNode.
		RelationshipTo(bikeNode, "OWNS").
		RelationshipTo(tripNode, "USED_ON").NamedC("r2").
		RelationshipFrom(userNode, "WAS_ON").NamedC("x").
		RelationshipBetween(cypher.ANode("SOMETHING").NamedByString("something")).NamedC("y"))
	expression := cypher.ExpressionWrap(userNode.Property("name")).MatchesPattern(".*aName").Get()
	statement, err := statementBuilder.Where(expression).ReturningByNamed(bikeNode, userNode).Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`)-[:`OWNS`]->(b:`Bike`)-[r2:`USED_ON`]->(t:`Trip`)<-[x:`WAS_ON`]-(u)-[y]-(something:`SOMETHING`) WHERE u.name =~ '.*aName' RETURN b, u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestChainedRelationshipWithPropertiesAndLength(t *testing.T) {
	tripNode := cypher.ANode("Trip").NamedByString("t")
	statementBuilder := cypher.Match(userNode.
		RelationshipTo(bikeNode, "OWNS").
		RelationshipTo(tripNode, "USED_ON").NamedC("r2").Min(1).Properties(cypher.MapOf("when", cypher.LiteralOf("2019-04-16"))).
		RelationshipFrom(userNode, "WAS_ON").NamedC("x").Max(2).Properties(cypher.MapOf("whatever", cypher.LiteralOf("2020-04-16"))).
		RelationshipBetween(cypher.ANode("SOMETHING").NamedByString("something")).NamedC("y").Length(2, 3).Properties(cypher.MapOf("idk", cypher.LiteralOf("2021-04-16"))))
	expression := cypher.ExpressionWrap(userNode.Property("name")).MatchesPattern(".*aName").Get()
	statement, err := statementBuilder.Where(expression).ReturningByNamed(bikeNode, userNode).Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`)-[:`OWNS`]->(b:`Bike`)-[r2:`USED_ON`*1.. {when: '2019-04-16'}]->(t:`Trip`)<-[x:`WAS_ON`*..2 {whatever: '2020-04-16'}]-(u)-[y*2..3 {idk: '2021-04-16'}]-(something:`SOMETHING`) WHERE u.name =~ '.*aName' RETURN b, u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestSizeOfRelationship(t *testing.T) {
	statementBuilder := cypher.Match(cypher.AnyNodeNamed("a"))
	expression := cypher.ExpressionWrap(cypher.AProperty("a", "name")).IsEqualTo(cypher.LiteralOf("Alice")).Get()
	statement, err := statementBuilder.Where(expression).
		Returning(cypher.SizeByPattern(cypher.AnyNodeNamed("a").RelationshipTo(cypher.AnyNode())).As("fof").Get()).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (a) WHERE a.name = 'Alice' RETURN size((a)-->()) AS fof"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestSizeOfRelationshipChain(t *testing.T) {
	statementBuilder := cypher.Match(cypher.AnyNodeNamed("a"))
	expression := cypher.ExpressionWrap(cypher.AProperty("a", "name")).IsEqualTo(cypher.LiteralOf("Alice")).Get()
	statement, err := statementBuilder.Where(expression).
		Returning(cypher.SizeByPattern(cypher.AnyNodeNamed("a").RelationshipTo(cypher.AnyNode()).RelationshipTo(cypher.AnyNode())).As("fof").Get()).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (a) WHERE a.name = 'Alice' RETURN size((a)-->()-->()) AS fof"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestSortOrderDefault(t *testing.T) {
	statement, err := cypher.Match(userNode).ReturningByNamed(userNode).
		OrderBySortItem(cypher.Sort(userNode.Property("name"))).Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) RETURN u ORDER BY u.name"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestSortOrderAscending(t *testing.T) {
	statement, err := cypher.Match(userNode).ReturningByNamed(userNode).
		OrderBySortItem(cypher.Sort(userNode.Property("name")).Ascending()).Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) RETURN u ORDER BY u.name ASC"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestSortOrderDescending(t *testing.T) {
	statement, err := cypher.Match(userNode).ReturningByNamed(userNode).
		OrderBySortItem(cypher.Sort(userNode.Property("name")).Descending()).Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) RETURN u ORDER BY u.name DESC"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestSortOrderConcatenation(t *testing.T) {
	statement, err := cypher.Match(userNode).ReturningByNamed(userNode).
		OrderBySortItem(cypher.Sort(userNode.Property("name")).Descending(),
			cypher.Sort(userNode.Property("age")).Ascending()).Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) RETURN u ORDER BY u.name DESC, u.age ASC"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestSortOrderDefaultExpression(t *testing.T) {
	statement, err := cypher.Match(userNode).ReturningByNamed(userNode).
		OrderBy(userNode.Property("name")).Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) RETURN u ORDER BY u.name"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestSortOrderAscendingExpression(t *testing.T) {
	statement, err := cypher.Match(userNode).ReturningByNamed(userNode).
		OrderBySortItem(userNode.Property("name").Ascending()).Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) RETURN u ORDER BY u.name ASC"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestSortOrderDescendingExpression(t *testing.T) {
	statement, err := cypher.Match(userNode).ReturningByNamed(userNode).
		OrderBySortItem(userNode.Property("name").Descending()).Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) RETURN u ORDER BY u.name DESC"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestSortOrderConcatenationExpression(t *testing.T) {
	statement, err := cypher.Match(userNode).ReturningByNamed(userNode).
		OrderBy(userNode.Property("name")).Descending().
		And(userNode.Property("age")).Ascending().
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) RETURN u ORDER BY u.name DESC, u.age ASC"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestSkip(t *testing.T) {
	statement, err := cypher.Match(userNode).ReturningByNamed(userNode).
		Skip(1).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) RETURN u SKIP 1"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestLimit(t *testing.T) {
	statement, err := cypher.Match(userNode).ReturningByNamed(userNode).
		Limit(1).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) RETURN u LIMIT 1"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestSkipAndLimit(t *testing.T) {
	statement, err := cypher.Match(userNode).ReturningByNamed(userNode).
		Skip(1).
		Limit(1).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) RETURN u SKIP 1 LIMIT 1"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestDistinct(t *testing.T) {
	statement, err := cypher.Match(userNode).ReturningDistinctByNamed(userNode).
		Skip(1).
		Limit(1).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) RETURN DISTINCT u SKIP 1 LIMIT 1"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}
