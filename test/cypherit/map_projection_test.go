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
		MatchElements(n).
		Returning(n.Project("__internalNeo4jId__", cypher.FunctionIdByNode(n), "name"))
	Assert(t, builder, "MATCH (n) RETURN n{__internalNeo4jId__: id(n), .name}")
	//
	builder = cypher.
		MatchElements(n).
		Returning(n.Project("name", "__internalNeo4jId__", cypher.FunctionIdByNode(n)))
	Assert(t, builder, "MATCH (n) RETURN n{.name, __internalNeo4jId__: id(n)}")
}

func TestDoc21221(t *testing.T) {
	var builder cypher.BuildableStatement
	expected := "MATCH (actor:`Person` {name: 'Tom Hanks'})-[:`ACTED_IN`]->(movie:`Movie`) RETURN " +
		"actor{.name, .realName, movies: collect(movie{.title, .released})}"
	actor := cypher.NewNode("Person").NamedByString("actor")
	movie := cypher.NewNode("Movie").NamedByString("movie")
	//
	builder = cypher.
		MatchElements(actor.WithRawProperties("name", cypher.LiteralOf("Tom Hanks")).
			RelationshipTo(movie, "ACTED_IN")).
		Returning(actor.Project("name", "realName", "movies", cypher.FunctionCollect(movie.Project("title", "released"))))
	Assert(t, builder, expected)
	//
	builder = cypher.
		MatchElements(actor.WithRawProperties("name", cypher.LiteralOf("Tom Hanks")).
			RelationshipTo(movie, "ACTED_IN")).
		Returning(actor.Project("name", "realName", "movies", cypher.FunctionCollect(movie.Project(movie.Property("title"), movie.Property("released")))))
	Assert(t, builder, expected)
	//
	builder = cypher.
		MatchElements(actor.WithRawProperties("name", cypher.LiteralOf("Tom Hanks")).
			RelationshipTo(movie, "ACTED_IN")).
		Returning(actor.Project("name", "realName", "movies", cypher.FunctionCollect(movie.Project("title", "year", movie.Property("released")))))
	Assert(t, builder, "MATCH (actor:`Person` {name: 'Tom Hanks'})-[:`ACTED_IN`]->(movie:`Movie`) RETURN actor{"+
		".name, .realName, movies: collect(movie{.title, year: movie.released})}")
}

func TestNested(t *testing.T) {
	var builder cypher.BuildableStatement
	n := cypher.NewNode("Person").NamedByString("p")
	m := cypher.NewNode("Movie").NamedByString("m")
	//
	builder = cypher.
		MatchElements(n.RelationshipTo(m, "ACTED_IN")).
		Returning(n.Project("__internalNeo4jId__", cypher.FunctionIdByNode(n), "name", "nested",
			m.Project("title", "__internalNeo4jId__", cypher.FunctionIdByNode(m))))
	Assert(t, builder, "MATCH (p:`Person`)-[:`ACTED_IN`]->(m:`Movie`) RETURN p{__internalNeo4jId__: id(p), "+
		".name, nested: m{.title, __internalNeo4jId__: id(m)}}")
}

func TestRequiredSymbolicNameShouldBeGenerated(t *testing.T) {
	var builder cypher.BuildableStatement
	person := cypher.NewNode("Person")
	//
	builder = cypher.
		MatchElements(person).
		Returning(person.Project("something"))
	Assert(t, builder, "")
}
