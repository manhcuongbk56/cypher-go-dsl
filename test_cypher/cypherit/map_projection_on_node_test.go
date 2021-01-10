package cypherit

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestProjectionSimple(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	n := cypher.AnyNodeNamed("n")
	builder = cypher.
		Match(n).
		Returning(n.Project("__internalNeo4jId__", cypher.FunctionIdByNode(n), "name"))
	Assert(t, builder, "MATCH (n) RETURN n{__internalNeo4jId__: id(n), .name}")
	//
	builder = cypher.
		Match(n).
		Returning(n.Project("name", "__internalNeo4jId__", cypher.FunctionIdByNode(n)))
	Assert(t, builder, "MATCH (n) RETURN n{.name, __internalNeo4jId__: id(n)}")
}

func TestDoc21221(t *testing.T) {
	var builder cypher.BuildableStatement
	expected := "MATCH (actor:`Person` {name: 'Tom Hanks'})-[:`ACTED_IN`]->(movie:`Movie`) RETURN " +
		"actor{.name, .realName, movies: collect(movie{.title, .released})}"
	actor := cypher.ANode("Person").NamedByString("actor")
	movie := cypher.ANode("Movie").NamedByString("movie")
	//
	builder = cypher.
		Match(actor.WithRawProperties("name", cypher.LiteralOf("Tom Hanks")).
			RelationshipTo(movie, "ACTED_IN")).
		Returning(actor.Project("name", "realName", "movies", cypher.FunctionCollect(movie.Project("title", "released"))))
	Assert(t, builder, expected)
	//
	builder = cypher.
		Match(actor.WithRawProperties("name", cypher.LiteralOf("Tom Hanks")).
			RelationshipTo(movie, "ACTED_IN")).
		Returning(actor.Project("name", "realName", "movies", cypher.FunctionCollect(movie.Project(movie.Property("title"), movie.Property("released")))))
	Assert(t, builder, expected)
	//
	builder = cypher.
		Match(actor.WithRawProperties("name", cypher.LiteralOf("Tom Hanks")).
			RelationshipTo(movie, "ACTED_IN")).
		Returning(actor.Project("name", "realName", "movies", cypher.FunctionCollect(movie.Project("title", "year", movie.Property("released")))))
	Assert(t, builder, "MATCH (actor:`Person` {name: 'Tom Hanks'})-[:`ACTED_IN`]->(movie:`Movie`) RETURN actor{"+
		".name, .realName, movies: collect(movie{.title, year: movie.released})}")
}

func TestNested(t *testing.T) {
	var builder cypher.BuildableStatement
	n := cypher.ANode("Person").NamedByString("p")
	m := cypher.ANode("Movie").NamedByString("m")
	//
	builder = cypher.
		Match(n.RelationshipTo(m, "ACTED_IN")).
		Returning(n.Project("__internalNeo4jId__", cypher.FunctionIdByNode(n), "name", "nested",
			m.Project("title", "__internalNeo4jId__", cypher.FunctionIdByNode(m))))
	Assert(t, builder, "MATCH (p:`Person`)-[:`ACTED_IN`]->(m:`Movie`) RETURN p{__internalNeo4jId__: id(p), "+
		".name, nested: m{.title, __internalNeo4jId__: id(m)}}")
}

func TestAddedProjections(t *testing.T) {
	var builder cypher.BuildableStatement
	p := cypher.ANode("Person").NamedByString("p")
	m := cypher.ANode("Movie").NamedByString("m")
	rel := p.RelationshipTo(m, "ACTED_IN").NamedByString("r")
	//
	builder = cypher.
		Match(rel).
		Returning(p.Project("__internalNeo4jId__", cypher.FunctionIdByNode(p), "name").
			And(rel).
			And(m).
			And(p.Property("foo")).
			And("a", p.Property("x")))
	Assert(t, builder, "MATCH (p:`Person`)-[r:`ACTED_IN`]->(m:`Movie`) RETURN p{__internalNeo4jId__: id(p), .name, r, m, .foo, a: p.x}")
}
