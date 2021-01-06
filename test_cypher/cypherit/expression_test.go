package cypherit

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestShouldRenderParameters(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.
		MatchElements(userNode).
		WhereConditionContainer(userNode.Property("a").IsEqualTo(cypher.Param("aParameter"))).
		DetachDeleteByNamed(userNode).
		ReturningByNamed(userNode)
	Assert(t, builder, "MATCH (u:`User`) WHERE u.a = $aParameter DETACH DELETE u RETURN u")
}

func TestShouldRenderMap(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.
		MatchElements(cypher.AnyNodeNamed("n")).
		Returning(cypher.FunctionPoint(cypher.MapOf(
			"latitude", cypher.Param("latitude"),
			"longitude", cypher.Param("longitude"),
			"crs", cypher.LiteralOf(4326))))
	Assert(t, builder, "MATCH (n) RETURN point({latitude: $latitude, longitude: $longitude, crs: 4326})")
}

func TestShouldRenderPointFunction(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	n := cypher.AnyNodeNamed("n")
	builder = cypher.
		MatchElements(n).
		WhereConditionContainer(cypher.FunctionDistance(n.Property("location"),
			cypher.FunctionPointByParameter(cypher.Param("point.point"))).
			Gt(cypher.Param("point.distance"))).
		ReturningByNamed(n)
	Assert(t, builder, "MATCH (n) WHERE distance(n.location, point($point.point)) > $point.distance RETURN n")
}
