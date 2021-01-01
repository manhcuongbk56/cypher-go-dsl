package cypherit

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestOrderOnWithShouldWork(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.
		MatchElements(cypher.NewNode("Movie").
			NamedByString("m").
			RelationshipFrom(cypher.NewNode("Person").NamedByString("p"), "ACTED_IN").Named("r")).
		With(cypher.CypherName("m"), cypher.CypherName("p")).
		OrderBySortItem(cypher.CypherSort(cypher.CypherProperty("m", "title")),
			cypher.CypherSort(cypher.CypherProperty("p", "name"))).
		Returning(cypher.CypherProperty("m", "title").As("movie").Get(),
			cypher.FunctionCollect(cypher.CypherProperty("p", "name")).As("actors").Get())
	Assert(t, builder, "MATCH (m:`Movie`)<-[r:`ACTED_IN`]-(p:`Person`) WITH m, p ORDER BY m.title, p.name RETURN m.title AS movie, collect(p.name) AS actors")
}

func TestConcatenatedOrdering(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.
		MatchElements(cypher.NewNode("Movie").
			NamedByString("m").
			RelationshipFrom(cypher.NewNode("Person").NamedByString("p"), "ACTED_IN").Named("r")).
		With(cypher.CypherName("m"), cypher.CypherName("p")).
		OrderBySortItem(cypher.CypherSort(cypher.CypherProperty("m", "title")).Ascending(),
			cypher.CypherSort(cypher.CypherProperty("p", "name")).Ascending()).
		Returning(cypher.CypherProperty("m", "title").As("movie").Get(),
			cypher.FunctionCollect(cypher.CypherProperty("p", "name")).As("actors").Get())
	Assert(t, builder, "MATCH (m:`Movie`)<-[r:`ACTED_IN`]-(p:`Person`) WITH m, p ORDER BY m.title ASC, p.name ASC RETURN m.title AS movie, collect(p.name) AS actors")
}
