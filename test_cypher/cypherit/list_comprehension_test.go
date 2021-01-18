package cypherit

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestListComprehensionSimple(t *testing.T) {
	var builder cypher.BuildableStatement
	name := cypher.ASymbolic("a")
	//
	builder = cypher.
		CypherReturning(cypher.CypherListWith(name).
			In(cypher.ListOf(cypher.LiteralOf(1), cypher.LiteralOf(2), cypher.LiteralOf(3), cypher.LiteralOf(4))).
			ReturningDefault())
	Assert(t, builder, "RETURN [a IN [1, 2, 3, 4]]")
}

func TestWithReturning(t *testing.T) {
	var builder cypher.BuildableStatement
	name := cypher.ASymbolic("a")
	//
	builder = cypher.
		CypherReturning(cypher.CypherListWith(name).
			In(cypher.ListOf(cypher.LiteralOf(1), cypher.LiteralOf(2), cypher.LiteralOf(3), cypher.LiteralOf(4))).
			Returning(name.Remainder(cypher.LiteralOf(2)).Get()))
	Assert(t, builder, "RETURN [a IN [1, 2, 3, 4] | (a % 2)]")
}

func TestWithWhere(t *testing.T) {
	var builder cypher.BuildableStatement
	name := cypher.ASymbolic("a")
	//
	builder = cypher.
		CypherReturning(cypher.CypherListWith(name).
			In(cypher.ListOf(cypher.LiteralOf(1), cypher.LiteralOf(2), cypher.LiteralOf(3), cypher.LiteralOf(4))).
			Where(name.Gt(cypher.LiteralOf(2)).Get()).
			ReturningDefault())
	Assert(t, builder, "RETURN [a IN [1, 2, 3, 4] WHERE a > 2]")
}

func TestWithWhereAndReturning(t *testing.T) {
	var builder cypher.BuildableStatement
	name := cypher.ASymbolic("a")
	//
	builder = cypher.
		CypherReturning(cypher.CypherListWith(name).
			In(cypher.ListOf(cypher.LiteralOf(1), cypher.LiteralOf(2), cypher.LiteralOf(3), cypher.LiteralOf(4))).
			Where(name.Gt(cypher.LiteralOf(2)).Get()).
			Returning(name.Remainder(cypher.LiteralOf(2)).Get()))
	Assert(t, builder, "RETURN [a IN [1, 2, 3, 4] WHERE a > 2 | (a % 2)]")
}

func TestSomeMoreExample(t *testing.T) {
	var builder cypher.BuildableStatement
	name := cypher.ASymbolic("x")
	//
	builder = cypher.
		CypherReturning(cypher.CypherListWith(name).
			In(cypher.Range(cypher.LiteralOf(0), cypher.LiteralOf(10))).
			Where(name.Remainder(cypher.LiteralOf(2)).IsEqualTo(cypher.LiteralOf(0)).Get()).
			Returning(name.Pow(cypher.LiteralOf(3)).Get()).
			As("result").Get())
	Assert(t, builder, "RETURN [x IN range(0, 10) WHERE (x % 2) = 0 | x^3] AS result")
	//
	builder = cypher.
		CypherReturning(cypher.CypherListWith(name).
			In(cypher.Range(cypher.LiteralOf(0), cypher.LiteralOf(10))).
			Where(name.Remainder(cypher.LiteralOf(2)).IsEqualTo(cypher.LiteralOf(0)).Get()).
			ReturningDefault().
			As("result").Get())
	Assert(t, builder, "RETURN [x IN range(0, 10) WHERE (x % 2) = 0] AS result")
	//
	builder = cypher.
		CypherReturning(cypher.CypherListWith(name).
			In(cypher.Range(cypher.LiteralOf(0), cypher.LiteralOf(10))).
			Returning(name.Pow(cypher.LiteralOf(3)).Get()).
			As("result").Get())
	Assert(t, builder, "RETURN [x IN range(0, 10) | x^3] AS result")
}
