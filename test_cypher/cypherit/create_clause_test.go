package cypherit

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestShouldRenderCreateWithoutReturn(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.CypherCreate(userNode)
	Assert(t, builder, "CREATE (u:`User`)")
	//
	builder = cypher.CypherCreate(userNode.RelationshipTo(bikeNode, "OWNS").NamedByString("o"))
	Assert(t, builder, "CREATE (u:`User`)-[o:`OWNS`]->(b:`Bike`)")
}

func TestShouldRenderMultipleCreatesWithoutReturn(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.CypherCreate(userNode).
		Create(bikeNode)
	Assert(t, builder, "CREATE (u:`User`) CREATE (b:`Bike`)")
	//
	builder = cypher.CypherCreate(userNode.RelationshipTo(bikeNode, "OWNS").NamedByString("o")).
		Create(OtherNode())
	Assert(t, builder, "CREATE (u:`User`)-[o:`OWNS`]->(b:`Bike`) CREATE (other:`Other`)")
}

func TestShouldRenderCreateReturn(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.CypherCreate(userNode).
		ReturningByNamed(userNode)
	Assert(t, builder, "CREATE (u:`User`) RETURN u")
	//
	r := userNode.RelationshipTo(bikeNode, "OWNS").NamedByString("o")
	builder = cypher.CypherCreate(r).
		ReturningByNamed(userNode, r)
	Assert(t, builder, "CREATE (u:`User`)-[o:`OWNS`]->(b:`Bike`) RETURN u, o")
	//
	builder = cypher.CypherCreate(userNode).
		ReturningByNamed(userNode).
		OrderBy(userNode.Property("name")).
		Skip(23).
		Limit(42)
	Assert(t, builder, "CREATE (u:`User`) RETURN u ORDER BY u.name SKIP 23 LIMIT 42")
}

func TestShouldRenderMultipleCreatesReturn(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.CypherCreate(userNode).
		Create(bikeNode).
		ReturningByNamed(userNode)
	Assert(t, builder, "CREATE (u:`User`) CREATE (b:`Bike`) RETURN u")
	//
	r := userNode.RelationshipTo(bikeNode, "OWNS").NamedByString("o")
	builder = cypher.
		CypherCreate(r).
		Create(OtherNode()).
		ReturningByNamed(userNode, r)
	Assert(t, builder, "CREATE (u:`User`)-[o:`OWNS`]->(b:`Bike`) CREATE (other:`Other`) RETURN u, o")
}

func TestShouldRenderCreateWithWith(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.CypherCreate(userNode).
		WithByNamed(userNode).
		ReturningByNamed(userNode)
	Assert(t, builder, "CREATE (u:`User`) WITH u RETURN u")
	//
	builder = cypher.
		CypherCreate(userNode).
		WithByNamed(userNode).
		Set(userNode.Property("x").To(cypher.LiteralOf("y")))
	Assert(t, builder, "CREATE (u:`User`) WITH u SET u.x = 'y'")
}

func TestMatchShouldExposeCreate(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.Match(userNode).
		Create(userNode.RelationshipTo(bikeNode, "OWNS").NamedByString("o"))
	Assert(t, builder, "MATCH (u:`User`) CREATE (u)-[o:`OWNS`]->(b:`Bike`)")
}

func TestWithShouldExposeCreate(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.Match(userNode).
		WithDistinctByNamed(userNode).
		Create(userNode.RelationshipTo(bikeNode, "OWNS").NamedByString("o"))
	Assert(t, builder, "MATCH (u:`User`) WITH DISTINCT u CREATE (u)-[o:`OWNS`]->(b:`Bike`)")
}
