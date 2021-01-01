package cypherit

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestSimple(t *testing.T) {
	statement, err := cypher.
		MatchElements(bikeNode).
		Match(userNode, cypher.NewNode("U").NamedByString("o")).
		ReturningByNamed(bikeNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (b:`Bike`) MATCH (u:`User`), (o:`U`) RETURN b"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestSimpleWhere(t *testing.T) {
	statement, err := cypher.
		MatchElements(bikeNode).
		Match(userNode, cypher.NewNode("U").NamedByString("o")).
		Where(userNode.Property("a").IsNull().Get()).
		ReturningByNamed(bikeNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (b:`Bike`) MATCH (u:`User`), (o:`U`) WHERE u.a IS NULL RETURN b"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestMultiWhere(t *testing.T) {
	statement, err := cypher.
		MatchElements(bikeNode).
		Where(bikeNode.Property("a").IsNotNull().Get()).
		Match(userNode, cypher.NewNode("U").NamedByString("o")).
		Where(userNode.Property("a").IsNull().Get()).
		ReturningByNamed(bikeNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (b:`Bike`) WHERE b.a IS NOT NULL MATCH (u:`User`), (o:`U`) WHERE u.a IS NULL RETURN b"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestMultiWhereMultiConditions(t *testing.T) {
	statement, err := cypher.
		MatchElements(bikeNode).
		Where(bikeNode.Property("a").IsNotNull().Get()).
		And(bikeNode.Property("b").IsNull().Get()).
		Match(userNode, cypher.NewNode("U").NamedByString("o")).
		Where(userNode.Property("a").IsNull().Or(userNode.InternalId().IsEqualTo(cypher.LiteralOf(4711)).Get()).Get()).
		ReturningByNamed(bikeNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (b:`Bike`) WHERE (b.a IS NOT NULL AND b.b IS NULL) MATCH (u:`User`), (o:`U`) WHERE (u.a IS NULL OR id(u) = 4711) RETURN b"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestOptionalMatch(t *testing.T) {
	statement, err := cypher.
		OptionalMatch(bikeNode).
		Match(userNode, cypher.NewNode("U").NamedByString("o")).
		Where(userNode.Property("a").IsNull().Get()).
		ReturningByNamed(bikeNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "OPTIONAL MATCH (b:`Bike`) MATCH (u:`User`), (o:`U`) WHERE u.a IS NULL RETURN b"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestUsingSameWithStepWithoutReassign(t *testing.T) {
	statement, err := cypher.
		MatchElements(bikeNode).
		WithByNamed(bikeNode).
		OptionalMatch(userNode).
		OptionalMatch(cypher.NewNode("Trip")).
		Returning(cypher.CypherAsterisk()).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (b:`Bike`) WITH b OPTIONAL MATCH (u:`User`) OPTIONAL MATCH (:`Trip`) RETURN *"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestUsingSameWithStepWithoutReassignThenUpdate(t *testing.T) {
	statement, err := cypher.
		MatchElements(bikeNode).
		WithByNamed(bikeNode).
		OptionalMatch(userNode).
		OptionalMatch(cypher.NewNode("Trip")).
		DeleteByString("u").
		Returning(cypher.CypherAsterisk()).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (b:`Bike`) WITH b OPTIONAL MATCH (u:`User`) OPTIONAL MATCH (:`Trip`) DELETE u RETURN *"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestQueryPartsShouldBeExtractableInQueries(t *testing.T) {
	statement, err := cypher.
		MatchElements(cypher.NewNode("S1").NamedByString("n")).
		Where(cypher.CypherProperty("n", "a").IsEqualTo(cypher.LiteralOf("A")).Get()).
		WithByString("n").
		Match(cypher.AnyNodeNamed("n").RelationshipTo(cypher.NewNode("S2").NamedByString("m"), "SOMEHOW_RELATED")).
		WithByString("n", "m").
		ReturningByString("n", "m").
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (n:`S1`) WHERE n.a = 'A' WITH n MATCH (n)-[:`SOMEHOW_RELATED`]->(m:`S2`) WITH n, m RETURN n, m"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestOptionalNext(t *testing.T) {
	statement, err := cypher.
		MatchElements(bikeNode).
		OptionalMatch(userNode, cypher.NewNode("U").NamedByString("o")).
		Where(userNode.Property("a").IsNull().Get()).
		ReturningByNamed(bikeNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (b:`Bike`) OPTIONAL MATCH (u:`User`), (o:`U`) WHERE u.a IS NULL RETURN b"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestOptionalMatchThenDelete(t *testing.T) {
	buildableStatement := cypher.
		MatchElements(bikeNode).
		OptionalMatch(userNode, cypher.NewNode("U").NamedByString("o")).
		DeleteByNamed(userNode, bikeNode)
	Assert(t, buildableStatement, "MATCH (b:`Bike`) OPTIONAL MATCH (u:`User`), (o:`U`) DELETE u, b")
}
