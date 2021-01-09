package cypherit

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestAllShouldWork(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	p := cypher.
		CypherPathByString("p").
		DefinedBy(cypher.AnyNodeNamed("a").
			RelationshipTo(cypher.AnyNodeNamed("b")).Min(1).Max(3))
	builder = cypher.MatchElements(p).
		Where(cypher.CypherProperty("a", "name").IsEqualTo(cypher.LiteralOf("Alice")).Get()).
		And(cypher.CypherProperty("b", "name").IsEqualTo(cypher.LiteralOf("Daniel")).Get()).
		And(cypher.PredicateAll("x").In(cypher.FunctionNodes(p)).
			Where(cypher.CypherProperty("x", "age").Gt(cypher.LiteralOf(30)).Get())).
		ReturningByNamed(p)

	Assert(t, builder, "MATCH p = (a)-[*1..3]->(b) WHERE (a.name = 'Alice' AND b.name = 'Daniel' AND all(x IN nodes(p) WHERE x.age > 30)) RETURN p")
}

func TestAnyShouldWork1(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	a := cypher.AnyNodeNamed("a")
	builder = cypher.
		MatchElements(a).
		Where(nil).
		Returning(a.Property("name"), a.Property("array"))

	Assert(t, builder, "MATCH (a) WHERE (a.name = 'Eskil' AND any(x IN a.array WHERE x = 'one')) RETURN a.name, a.array")
}

func TestAnyShouldWork(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	a := cypher.AnyNodeNamed("a")
	builder = cypher.
		MatchElements(a).
		Where(cypher.CypherProperty("a", "name").IsEqualTo(cypher.LiteralOf("Eskil")).Get()).
		And(cypher.PredicateAny("x").In(a.Property("array")).
			Where(cypher.CypherName("x").IsEqualTo(cypher.LiteralOf("one")).Get())).
		Returning(a.Property("name"), a.Property("array"))

	Assert(t, builder, "MATCH (a) WHERE (a.name = 'Eskil' AND any(x IN a.array WHERE x = 'one')) RETURN a.name, a.array")
}

func TestNoneShouldWork(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	p := cypher.
		CypherPathByString("p").
		DefinedBy(cypher.AnyNodeNamed("a").
			RelationshipTo(cypher.AnyNodeNamed("b")).Min(1).Max(3))
	builder = cypher.MatchElements(p).
		Where(cypher.CypherProperty("a", "name").IsEqualTo(cypher.LiteralOf("Alice")).Get()).
		And(cypher.
			PredicateNone("x").
			In(cypher.FunctionNodes(p)).
			Where(cypher.CypherProperty("x", "age").IsEqualTo(cypher.LiteralOf(25)).Get())).
		ReturningByNamed(p)

	Assert(t, builder, "MATCH p = (a)-[*1..3]->(b) WHERE (a.name = 'Alice' AND none(x IN nodes(p) WHERE x.age = 25)) RETURN p")
}

func TestSingleShouldWork(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	p := cypher.
		CypherPathByString("p").
		DefinedBy(cypher.AnyNodeNamed("n").
			RelationshipTo(cypher.AnyNodeNamed("b")))
	builder = cypher.MatchElements(p).
		Where(cypher.CypherProperty("n", "name").IsEqualTo(cypher.LiteralOf("Alice")).Get()).
		And(cypher.
			PredicateSingle("var").
			In(cypher.FunctionNodes(p)).
			Where(cypher.CypherProperty("var", "eyes").IsEqualTo(cypher.LiteralOf("blue")).Get())).
		ReturningByNamed(p)

	Assert(t, builder, "MATCH p = (n)-->(b) WHERE (n.name = 'Alice' AND single(var IN nodes(p) WHERE var.eyes = 'blue')) RETURN p")
}
