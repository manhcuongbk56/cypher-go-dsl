package test

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestFindAllMovies(t *testing.T) {
	movie := cypher.CypherNewNode("Movie").NamedByString("m")
	statement, _ := cypher.CypherMatch(movie).
		ReturningByNamed(movie).
		Build()
	query := cypher.NewRenderer().Render(statement)
	expect := "MATCH (m:`Movie`) RETURN m"
	if query != expect {
		t.Errorf("%s is incorrect \n %s", query, expect)
	}
}

func TestDefaultStatementBuilder_OptionalMatch(t *testing.T) {
	farm := cypher.CypherNewNode("Farm").NamedByString("b")
	statement, _ := cypher.CypherMatch(farm).
		Where(cypher.ConditionsNot(farm.RelationshipFrom(cypher.CypherAnyNode(), "HAS"))).
		WithByString("b").
		OptionalMatch(farm.RelationshipTo(cypher.CypherAnyNode1("p"), "HAS")).
		ReturningByString("b", "p").
		Build()
	query := cypher.NewRenderer().Render(statement)
	expect := "MATCH (b:`Farm`) WHERE NOT (b)<-[:`HAS`]-() WITH b OPTIONAL MATCH (b)-[:`HAS`]->(p) RETURN b, p"
	if query != expect {
		t.Errorf("%s is incorrect \n %s", query, expect)
	}
}
