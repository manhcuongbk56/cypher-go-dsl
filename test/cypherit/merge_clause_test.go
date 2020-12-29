package cypherit

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestShouldRenderMergeWithoutReturn(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.CypherMerge(userNode)
	Assert(t, builder, "MERGE (u:`User`)")
	//
	builder = cypher.CypherMerge(userNode.RelationshipTo(bikeNode, "OWNS").NamedByString("o"))
	Assert(t, builder, "MERGE (u:`User`)-[o:`OWNS`]->(b:`Bike`)")
}

func TestShouldRenderMultipleMergeWithoutReturn(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.CypherMerge(userNode).
		Merge(bikeNode)
	Assert(t, builder, "MERGE (u:`User`) MERGE (b:`Bike`)")
	//
	builder = cypher.
		CypherMerge(userNode.RelationshipTo(bikeNode, "OWNS").NamedByString("o")).
		Merge(cypher.NewNode("Other"))
	Assert(t, builder, "MERGE (u:`User`)-[o:`OWNS`]->(b:`Bike`) MERGE (:`Other`)")
}

func TestShouldRenderMergeReturn(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.CypherMerge(userNode).
		ReturningByNamed(userNode)
	Assert(t, builder, "MERGE (u:`User`) RETURN u")
	//
	r := userNode.RelationshipTo(bikeNode, "OWNS").NamedByString("o")
	builder = cypher.CypherMerge(r).
		ReturningByNamed(userNode, r)
	Assert(t, builder, "MERGE (u:`User`)-[o:`OWNS`]->(b:`Bike`) RETURN u, o")
	//
	builder = cypher.CypherMerge(userNode).
		ReturningByNamed(userNode).
		OrderBy(userNode.Property("name")).
		Skip(23).
		Limit(42)
	Assert(t, builder, "MERGE (u:`User`) RETURN u ORDER BY u.name SKIP 23 LIMIT 42")
}

func TestShouldRenderMultipleMergesReturn(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.CypherMerge(userNode).
		Merge(bikeNode).
		ReturningByNamed(userNode)
	Assert(t, builder, "MERGE (u:`User`) MERGE (b:`Bike`) RETURN u")
	//
	r := userNode.RelationshipTo(bikeNode, "OWNS").NamedByString("o")
	builder = cypher.
		CypherMerge(r).
		Merge(cypher.NewNode("Other")).
		ReturningByNamed(userNode, r)
	Assert(t, builder, "MERGE (u:`User`)-[o:`OWNS`]->(b:`Bike`) MERGE (:`Other`) RETURN u, o")
}

func TestShouldRenderMergeWithWith(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.CypherMerge(userNode).
		WithByNamed(userNode).
		ReturningByNamed(userNode)
	Assert(t, builder, "MERGE (u:`User`) WITH u RETURN u")
	//
	builder = cypher.
		CypherMerge(userNode).
		WithByNamed(userNode).
		Set(userNode.Property("x").To(cypher.LiteralOf("y")))
	Assert(t, builder, "MERGE (u:`User`) WITH u SET u.x = 'y'")
}

func TestMatchShouldExposeMerge(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.CypherMerge(userNode).
		Merge(userNode.RelationshipTo(bikeNode, "OWNS").NamedByString("o"))
	Assert(t, builder, "MATCH (u:`User`) MERGE (u)-[o:`OWNS`]->(b:`Bike`)")
}

func TestWithShouldExposeMerge(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.MatchElements(userNode).
		WithDistinctByNamed(userNode).
		Merge(userNode.RelationshipTo(bikeNode, "OWNS").NamedByString("o"))
	Assert(t, builder, "MATCH (u:`User`) WITH DISTINCT u MERGE (u)-[o:`OWNS`]->(b:`Bike`)")
}

func TestMixedCreateAndMerge(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	tripNode := cypher.NewNode("Trip").NamedByString("t")
	builder = cypher.CypherCreate(userNode).
		Merge(userNode.RelationshipTo(bikeNode, "OWNS").NamedByString("o")).
		WithDistinctByNamed(bikeNode).
		Merge(tripNode.RelationshipFrom(bikeNode, "USED_ON")).
		Returning(cypher.CypherAsterisk())
	Assert(t, builder, "CREATE (u:`User`) MERGE (u)-[o:`OWNS`]->(b:`Bike`) WITH DISTINCT b MERGE (t:`Trip`)<-[:`USED_ON`]-(b) RETURN *")
}

func TestSingleCreateAction(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	halloWeltString := cypher.LiteralOf("Hallo, Welt")
	builder = cypher.CypherMerge(userNode).
		OnCreate().Set(userNode.Property("p").To(halloWeltString))
	Assert(t, builder, "MERGE (u:`User`) ON CREATE SET u.p = 'Hallo, Welt'")
	//
	builder = cypher.CypherMerge(userNode).
		OnCreate().Set(userNode.Property("p"), halloWeltString)
	Assert(t, builder, "MERGE (u:`User`) ON CREATE SET u.p = 'Hallo, Welt'")
}

func TestSingleMatchAction(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	halloWeltString := cypher.LiteralOf("Hallo, Welt")
	builder = cypher.CypherMerge(userNode).
		OnMatch().Set(userNode.Property("p").To(halloWeltString))
	Assert(t, builder, "MERGE (u:`User`) ON MATCH SET u.p = 'Hallo, Welt'")
	//
	builder = cypher.CypherMerge(userNode).
		OnMatch().Set(userNode.Property("p"), halloWeltString)
	Assert(t, builder, "MERGE (u:`User`) ON MATCH SET u.p = 'Hallo, Welt'")
}

func TestStuffThanSingleMatchAction(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.CypherCreate(bikeNode).
		Set(bikeNode.Property("nice").To(cypher.CypherLiteralTrue())).
		Merge(userNode).OnMatch().Set(userNode.Property("happy").To(cypher.CypherLiteralTrue())).
		Create(userNode.RelationshipTo(bikeNode, "OWNS"))
	Assert(t, builder, "CREATE (b:`Bike`) SET b.nice = true MERGE (u:`User`) ON MATCH SET u.happy = true CREATE (u)-[:`OWNS`]->(b)")
}

func TestSingleActionMultipleProperties(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.CypherMerge(userNode).
		OnMatch().Set(userNode.Property("p1").To(cypher.LiteralOf("v1")), userNode.Property("p2").To(cypher.LiteralOf("v2")))
	Assert(t, builder, "MERGE (u:`User`) ON MATCH SET u.p1 = 'v1', u.p2 = 'v2'")
	//
	builder = cypher.CypherMerge(userNode).
		OnCreate().Set(userNode.Property("p1").To(cypher.LiteralOf("v1")), userNode.Property("p2").To(cypher.LiteralOf("v2")))
	Assert(t, builder, "MERGE (u:`User`) ON CREATE SET u.p1 = 'v1', u.p2 = 'v2'")
}

func TestMultipleActionsMultipleProperties(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.CypherMerge(userNode).
		OnMatch().Set(userNode.Property("p1").To(cypher.LiteralOf("v1")), userNode.Property("p2").To(cypher.LiteralOf("v2"))).
		OnCreate().Set(userNode.Property("p3").To(cypher.LiteralOf("v3")), userNode.Property("p4").To(cypher.LiteralOf("v4")))
	Assert(t, builder, "MERGE (u:`User`) ON MATCH SET u.p1 = 'v1', u.p2 = 'v2' ON CREATE SET u.p3 = 'v3', u.p4 = 'v4'")
}

func TestSingleCreateThanMatchAction(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	helloWorldString := cypher.LiteralOf("Hello, World")
	halloWeltString := cypher.LiteralOf("Hallo, Welt")
	builder = cypher.CypherMerge(userNode).
		OnCreate().Set(userNode.Property("p").To(helloWorldString)).
		OnMatch().Set(userNode.Property("p").To(halloWeltString))
	Assert(t, builder, "MERGE (u:`User`) ON CREATE SET u.p = 'Hello, World' ON MATCH SET u.p = 'Hallo, Welt'")
}

func TestSingleMatchThanCreateAction(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	helloWorldString := cypher.LiteralOf("Hello, World")
	halloWeltString := cypher.LiteralOf("Hallo, Welt")
	builder = cypher.CypherMerge(userNode).
		OnMatch().Set(userNode.Property("p").To(halloWeltString)).
		OnCreate().Set(userNode.Property("p").To(helloWorldString))
	Assert(t, builder, "MERGE (u:`User`) ON MATCH SET u.p = 'Hallo, Welt' ON CREATE SET u.p = 'Hello, World'")
}

func TestMultipleMatchesAndCreates(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.CypherMerge(userNode).
		OnMatch().Set(userNode.Property("p1").To(cypher.LiteralOf("v1"))).
		OnCreate().Set(userNode.Property("p2").To(cypher.LiteralOf("v2"))).
		OnCreate().Set(userNode.Property("p3").To(cypher.LiteralOf("v3"))).
		OnMatch().Set(userNode.Property("p4").To(cypher.LiteralOf("v4")))
	Assert(t, builder, "MERGE (u:`User`) ON MATCH SET u.p1 = 'v1' ON CREATE SET u.p2 = 'v2' ON CREATE SET u.p3 = 'v3' ON MATCH SET u.p4 = 'v4'")
}

func TestActionThanSet(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.CypherMerge(userNode).
		OnMatch().Set(userNode.Property("p1").To(cypher.LiteralOf("v1"))).
		Set(userNode.Property("p2").To(cypher.LiteralOf("v2"))).
		ReturningByNamed(userNode)
	Assert(t, builder, "MERGE (u:`User`) ON MATCH SET u.p1 = 'v1' SET u.p2 = 'v2' RETURN u")
}

func TestActionThanContinue(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.CypherMerge(userNode).
		OnMatch().Set(userNode.Property("p1").To(cypher.LiteralOf("v1"))).
		WithByNamed(userNode).
		ReturningByNamed(userNode)
	Assert(t, builder, "MERGE (u:`User`) ON MATCH SET u.p1 = 'v1' WITH u RETURN u")
}
