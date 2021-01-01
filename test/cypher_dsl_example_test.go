package test

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestFindAllMovies(t *testing.T) {
	movie := cypher.NewNode("Movie").NamedByString("m")
	statement, _ := cypher.MatchElements(movie).
		ReturningByNamed(movie).
		Build()
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (m:`Movie`) RETURN m"
	if query != expect {
		t.Errorf("%s is incorrect \n %s", query, expect)
	}
}

func TestDefaultStatementBuilder_OptionalMatch(t *testing.T) {
	farm := cypher.NewNode("Farm").NamedByString("b")
	statement, _ := cypher.MatchElements(farm).
		Where(cypher.ConditionsNot(farm.RelationshipFrom(cypher.AnyNode(), "HAS"))).
		WithByString("b").
		OptionalMatch(farm.RelationshipTo(cypher.CypherAnyNode1("p"), "HAS")).
		ReturningByString("b", "p").
		Build()
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (b:`Farm`) WHERE NOT (b)<-[:`HAS`]-() WITH b OPTIONAL MATCH (b)-[:`HAS`]->(p) RETURN b, p"
	if query != expect {
		t.Errorf("%s is incorrect \n %s", query, expect)
	}
}
