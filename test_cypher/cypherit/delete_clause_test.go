package cypherit

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestShouldRenderDeleteWithoutReturn(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.Match(userNode).
		DetachDeleteByNamed(userNode)
	Assert(t, builder, "MATCH (u:`User`) DETACH DELETE u")

	builder = cypher.Match(userNode).
		WithByNamed(userNode).
		DetachDeleteByNamed(userNode)
	Assert(t, builder, "MATCH (u:`User`) WITH u DETACH DELETE u")
	//
	builder = cypher.Match(userNode).
		WhereConditionContainer(userNode.Property("a").IsNotNull()).
		And(userNode.Property("b").IsNull().Get()).
		DeleteByNamed(userNode)
	Assert(t, builder, "MATCH (u:`User`) WHERE (u.a IS NOT NULL AND u.b IS NULL) DELETE u")
	//
	builder = cypher.Match(userNode, bikeNode).
		DeleteByNamed(userNode, bikeNode)
	Assert(t, builder, "MATCH (u:`User`), (b:`Bike`) DELETE u, b")
}

func TestShouldRenderDeleteWithReturn(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.Match(userNode).
		DetachDeleteByNamed(userNode).
		ReturningByNamed(userNode)
	Assert(t, builder, "MATCH (u:`User`) DETACH DELETE u RETURN u")

	builder = cypher.Match(userNode).
		WhereConditionContainer(userNode.Property("a").IsNotNull()).
		And(userNode.Property("b").IsNull().Get()).
		DetachDeleteByNamed(userNode).
		ReturningByNamed(userNode).
		OrderBy(userNode.Property("a")).Ascending().
		Skip(2).Limit(1)
	Assert(t, builder, "MATCH (u:`User`) WHERE (u.a IS NOT NULL AND u.b IS NULL) DETACH DELETE u RETURN u ORDER "+
		"BY u.a ASC SKIP 2 LIMIT 1")
	//
	builder = cypher.Match(userNode).
		WhereConditionContainer(userNode.Property("a").IsNotNull()).
		And(userNode.Property("b").IsNull().Get()).
		DetachDeleteByNamed(userNode).
		ReturningDistinctByNamed(userNode).
		OrderBy(userNode.Property("a")).Ascending().
		Skip(2).Limit(1)
	Assert(t, builder, "MATCH (u:`User`) WHERE (u.a IS NOT NULL AND u.b IS NULL) DETACH DELETE u RETURN DISTINCT"+
		" u ORDER BY u.a ASC SKIP 2 LIMIT 1")
	//
	builder = cypher.Match(userNode, bikeNode).
		DeleteByNamed(userNode, bikeNode)
	Assert(t, builder, "MATCH (u:`User`), (b:`Bike`) DELETE u, b")
}

func TestShouldRenderNodeDelete(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	n := cypher.AnyNodeNamed("n")
	r := n.RelationshipBetween(cypher.AnyNode()).NamedByString("r0")
	builder = cypher.
		Match(n).
		WhereConditionContainer(n.InternalId().IsEqualTo(cypher.LiteralOf(4711))).
		OptionalMatch(r).
		DeleteByNamed(r, n)
	Assert(t, builder, "MATCH (n) WHERE id(n) = 4711 OPTIONAL MATCH (n)-[r0]-() DELETE r0, n")
}

func TestShouldRenderChainedDeletes(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	n := cypher.AnyNodeNamed("n")
	r := n.RelationshipBetween(cypher.AnyNode()).NamedByString("r0")
	builder = cypher.
		Match(n).
		WhereConditionContainer(n.InternalId().IsEqualTo(cypher.LiteralOf(4711))).
		OptionalMatch(r).
		DeleteByNamed(r, n).
		DeleteByNamed(bikeNode).
		DetachDeleteByNamed(userNode)
	Assert(t, builder, "MATCH (n) WHERE id(n) = 4711 OPTIONAL MATCH (n)-[r0]-() DELETE r0, n DELETE b DETACH "+
		"DELETE u")
}
