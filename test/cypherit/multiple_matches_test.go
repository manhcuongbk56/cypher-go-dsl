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
	query := cypher.NewRenderer().Render(statement)
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
	query := cypher.NewRenderer().Render(statement)
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
	query := cypher.NewRenderer().Render(statement)
	expect := "MATCH (b:`Bike`) WHERE b.a IS NOT NULL MATCH (u:`User`), (o:`U`) WHERE u.a IS NULL RETURN b"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestMultiWhereMultiCondition(t *testing.T) {
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
	query := cypher.NewRenderer().Render(statement)
	expect := "MATCH (b:`Bike`) WHERE (b.a IS NOT NULL AND b.b IS NULL) MATCH (u:`User`), (o:`U`) WHERE (u.a IS NULL OR id(u) = 4711) RETURN b"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}
