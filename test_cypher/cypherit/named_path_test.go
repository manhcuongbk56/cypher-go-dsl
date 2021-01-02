package cypherit

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestDoc3148(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	namePath := cypher.
		CypherPathByString("p").
		DefinedBy(cypher.AnyNodeNamed("michael").
			WithRawProperties("name", cypher.LiteralOf("Michael Douglas")).
			RelationshipTo(cypher.AnyNode()))
	builder = cypher.MatchElements(namePath).
		ReturningByNamed(namePath)

	Assert(t, builder, "MATCH p = (michael {name: 'Michael Douglas'})-->() RETURN p")
}

func TestShouldWorkInListComprehensions(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	namePath := cypher.
		CypherPathByString("p").
		DefinedBy(cypher.AnyNodeNamed("n").
			RelationshipTo(cypher.AnyNode(), "LIKES", "OWNS").
			Unbounded())
	builder = cypher.CypherReturning(cypher.ListBasedOnNamed(namePath).ReturningByNamed(namePath))

	Assert(t, builder, "RETURN [p = (n)-[:`LIKES`|`OWNS`*]->() | p]")
}
