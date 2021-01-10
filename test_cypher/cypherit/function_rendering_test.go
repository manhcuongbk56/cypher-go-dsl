package cypherit

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestInWhereClause(t *testing.T) {
	statement, err := cypher.
		Match(userNode).
		Where(userNode.InternalId().IsEqualTo(cypher.LiteralOf(1)).Get()).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) WHERE id(u) = 1 RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestInReturnClause(t *testing.T) {
	statement, err := cypher.
		Match(userNode).
		Returning(cypher.FunctionCount(userNode)).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) RETURN count(u)"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestInReturnClauseWithDistinct(t *testing.T) {
	statement, err := cypher.
		Match(userNode).
		Returning(cypher.FunctionCountDistinct(userNode)).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) RETURN count(DISTINCT u)"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestAliasedInReturnClause(t *testing.T) {
	statement, err := cypher.
		Match(userNode).
		Returning(cypher.FunctionCount(userNode).As("cnt").Get()).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) RETURN count(u) AS cnt"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestShouldSupportMoreThanOneArgument(t *testing.T) {
	statement, err := cypher.
		Match(userNode).
		Returning(cypher.FunctionCoalesce(userNode.Property("a"), userNode.Property("b"), cypher.LiteralOf("¯\\_(ツ)_/¯"))).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	//TODO: some thing went wrong with escape letter
	expect := "MATCH (u:`User`) RETURN coalesce(u.a, u.b, '¯\\\\_(ツ)_/¯')"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestLiteralShouldDealWithNil(t *testing.T) {
	statement, err := cypher.
		Match(userNode).
		Returning(cypher.FunctionCoalesce(cypher.LiteralOf(nil), userNode.Property("field")).As("p").Get()).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) RETURN coalesce(NULL, u.field) AS p"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestFunctionBaseOnRelationship(t *testing.T) {
	relationShip := cypher.ANode("Person").NamedByString("bacon").
		WithRawProperties("name", cypher.LiteralOf("Kevin Bacon")).
		RelationshipBetween(cypher.ANode("Person").NamedByString("meg").WithRawProperties("name", cypher.LiteralOf("Meg Ryan"))).
		Unbounded()
	statement, err := cypher.
		Match(cypher.AShortestPath("p").DefinedBy(relationShip)).
		ReturningByString("p").
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH p = shortestPath((bacon:`Person` {name: 'Kevin Bacon'})-[*]-(meg:`Person` {name: 'Meg Ryan'})) RETURN p"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}
