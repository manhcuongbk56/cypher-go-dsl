package test_cypher

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestFindAllMovies(t *testing.T) {
	movie := cypher.ANode("Movie").NamedByString("m")
	statement, _ := cypher.Match(movie).
		ReturningByNamed(movie).
		Build()
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (m:`Movie`) RETURN m"
	if query != expect {
		t.Errorf("%s is incorrect \n %s", query, expect)
	}
}

func TestDefaultStatementBuilder_OptionalMatch(t *testing.T) {
	farm := cypher.ANode("Farm").NamedByString("b")
	statement, _ := cypher.Match(farm).
		Where(cypher.ConditionsNotByPattern(farm.RelationshipFrom(cypher.AnyNode(), "HAS"))).
		WithByString("b").
		OptionalMatch(farm.RelationshipTo(cypher.AnyNodeNamed("p"), "HAS")).
		ReturningByString("b", "p").
		Build()
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (b:`Farm`) WHERE NOT (b)<-[:`HAS`]-() WITH b OPTIONAL MATCH (b)-[:`HAS`]->(p) RETURN b, p"
	if query != expect {
		t.Errorf("%s is incorrect \n %s", query, expect)
	}
}
