package cypherit

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestShouldRenderOperations(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	n := cypher.AnyNodeNamed("n")
	builder = cypher.
		Match(n).
		Returning(cypher.LiteralOf(1).Add(cypher.LiteralOf(2)).Get())
	Assert(t, builder, "MATCH (n) RETURN (1 + 2)")
}

func TestShouldRenderComparison(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	n := cypher.AnyNodeNamed("n")
	builder = cypher.
		Match(n).
		Returning(cypher.LiteralOf(1).Gt(cypher.LiteralOf(2)).Get())
	Assert(t, builder, "MATCH (n) RETURN 1 > 2")
	//
	builder = cypher.
		Match(n).
		Returning(cypher.LiteralOf(1).Gt(cypher.LiteralOf(2)).IsTrue().Get())
	Assert(t, builder, "MATCH (n) RETURN (1 > 2) = true")
	//
	builder = cypher.
		Match(n).
		Returning(cypher.LiteralOf(1).Gt(cypher.LiteralOf(2)).IsTrue().IsFalse().Get())
	Assert(t, builder, "MATCH (n) RETURN ((1 > 2) = true) = false")
}
