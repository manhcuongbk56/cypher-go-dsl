package cypher_go_dsl

import (
	"testing"
)

func TestFindAllMovies(t *testing.T) {
	movie := CypherNewNode("Movie").NamedByString("m")
	statement, _ := Matchs(movie).
		ReturningByNamed(movie).
		Build()
	query := NewRenderer().Render(statement)
	expect := "MATCH (m:`Movie`) RETURN m"
	if query != expect {
		t.Errorf("%s is incorrect \n %s", query, expect)
	}
}

func TestDefaultStatementBuilder_OptionalMatch(t *testing.T) {
	farm := CypherNewNode("Farm").NamedByString("b")
	statement, _ := Matchs(farm).
		Where(ConditionsNot(farm.RelationshipFrom(CypherAnyNode(), "HAS"))).
		WithByString("b").
		OptionalMatch(farm.RelationshipTo(CypherAnyNode1("p"), "HAS")).
		ReturningByString("b", "p").
		Build()
	query := NewRenderer().Render(statement)
	expect := "MATCH (b:`Farm`) WHERE NOT (b)<-[:`HAS`]-() WITH b OPTIONAL MATCH (b)-[:`HAS`]->(p) RETURN b, p"
	if query != expect {
		t.Errorf("%s is incorrect \n %s", query, expect)
	}
}
