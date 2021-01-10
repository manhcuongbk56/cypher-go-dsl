package cypherit

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

var bikeNode = cypher.ANode("Bike").NamedByString("b")
var userNode = cypher.ANode("User").NamedByString("u")

func TestSimpleWith(t *testing.T) {
	statement, err := cypher.
		Match(userNode.RelationshipTo(bikeNode, "OWNS")).
		Where(userNode.Property("a").IsNull().Get()).
		WithByNamed(bikeNode, userNode).
		ReturningByNamed(bikeNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`)-[:`OWNS`]->(b:`Bike`) WHERE u.a IS NULL WITH b, u RETURN b"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestShouldRendererLeadingWith(t *testing.T) {
	statement, err := cypher.
		With(cypher.AParam("listOfPropertyMaps").As("p").Get()).
		UnwindByString("p").As("item").
		ReturningByString("item").
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "WITH $listOfPropertyMaps AS p UNWIND p AS item RETURN item"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestSimpleWithChained(t *testing.T) {
	tripNode := cypher.ANode("Trip").NamedByString("t")
	statement, err := cypher.
		Match(userNode.RelationshipTo(bikeNode, "OWNS")).
		Where(userNode.Property("a").IsNull().Get()).
		WithByNamed(bikeNode, userNode).
		Match(tripNode).
		Where(tripNode.Property("name").IsEqualTo(cypher.LiteralOf("Festive500")).Get()).
		WithByNamed(bikeNode, userNode, tripNode).
		ReturningByNamed(bikeNode, userNode, tripNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`)-[:`OWNS`]->(b:`Bike`) WHERE u.a IS NULL WITH b, u MATCH (t:`Trip`) WHERE t.name = 'Festive500' WITH b, u, t RETURN b, u, t"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestDeletingSimpleWith(t *testing.T) {
	statement, err := cypher.
		Match(userNode.RelationshipTo(bikeNode, "OWNS")).
		Where(userNode.Property("a").IsNull().Get()).
		DeleteByNamed(userNode).
		WithByNamed(bikeNode, userNode).
		ReturningByNamed(bikeNode, userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`)-[:`OWNS`]->(b:`Bike`) WHERE u.a IS NULL DELETE u WITH b, u RETURN b, u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestDeletingSimpleWithReverse(t *testing.T) {
	statement, err := cypher.
		Match(userNode.RelationshipTo(bikeNode, "OWNS")).
		Where(userNode.Property("a").IsNull().Get()).
		WithByNamed(bikeNode, userNode).
		DeleteByNamed(userNode).
		ReturningByNamed(bikeNode, userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`)-[:`OWNS`]->(b:`Bike`) WHERE u.a IS NULL WITH b, u DELETE u RETURN b, u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestMixedClauseWithWith(t *testing.T) {
	tripNode := cypher.ANode("Trip").NamedByString("t")
	statement, err := cypher.
		Match(userNode.RelationshipTo(bikeNode, "OWNS")).
		Match(tripNode).
		DeleteByNamed(tripNode).
		WithByNamed(bikeNode, tripNode).
		Match(userNode).
		WithByNamed(bikeNode, userNode).
		ReturningByNamed(bikeNode, userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`)-[:`OWNS`]->(b:`Bike`) MATCH (t:`Trip`) DELETE t WITH b, t MATCH (u) WITH b, u RETURN b, u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}
