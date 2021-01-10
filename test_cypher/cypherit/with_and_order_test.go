package cypherit

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestOrderOnWithShouldWork(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.
		Match(cypher.ANode("Movie").
			NamedByString("m").
			RelationshipFrom(cypher.ANode("Person").NamedByString("p"), "ACTED_IN").Named("r")).
		With(cypher.ASymbolic("m"), cypher.ASymbolic("p")).
		OrderBySortItem(cypher.CypherSort(cypher.AProperty("m", "title")),
			cypher.CypherSort(cypher.AProperty("p", "name"))).
		Returning(cypher.AProperty("m", "title").As("movie").Get(),
			cypher.FunctionCollect(cypher.AProperty("p", "name")).As("actors").Get())
	Assert(t, builder, "MATCH (m:`Movie`)<-[r:`ACTED_IN`]-(p:`Person`) WITH m, p ORDER BY m.title, p.name RETURN m.title AS movie, collect(p.name) AS actors")
}

func TestConcatenatedOrdering(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.
		Match(cypher.ANode("Movie").
			NamedByString("m").
			RelationshipFrom(cypher.ANode("Person").NamedByString("p"), "ACTED_IN").Named("r")).
		With(cypher.ASymbolic("m"), cypher.ASymbolic("p")).
		OrderBySortItem(cypher.CypherSort(cypher.AProperty("m", "title")).Ascending(),
			cypher.CypherSort(cypher.AProperty("p", "name")).Ascending()).
		Returning(cypher.AProperty("m", "title").As("movie").Get(),
			cypher.FunctionCollect(cypher.AProperty("p", "name")).As("actors").Get())
	Assert(t, builder, "MATCH (m:`Movie`)<-[r:`ACTED_IN`]-(p:`Person`) WITH m, p ORDER BY m.title ASC, p.name ASC RETURN m.title AS movie, collect(p.name) AS actors")
}
