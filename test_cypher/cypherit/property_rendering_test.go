package cypherit

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestShouldRenderNodeProperties(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	nodes := make([]cypher.Node, 0)
	nodes = append(nodes, cypher.NewNodeWithProperties("Test", cypher.MapOf("a", cypher.LiteralOf("b"))).NamedByString("test"))
	nodes = append(nodes, cypher.ANode("Test").NamedByString("test").WithProperties(cypher.MapOf("a", cypher.LiteralOf("b"))))
	nodes = append(nodes, cypher.ANode("Test").NamedByString("test").WithRawProperties("a", cypher.LiteralOf("b")))
	for _, nodeWithProperties := range nodes {
		builder = cypher.Match(nodeWithProperties).
			Returning(cypher.AnAsterisk())
		Assert(t, builder, "MATCH (test:`Test` {a: 'b'}) RETURN *")
		builder = cypher.Merge(nodeWithProperties).
			Returning(cypher.AnAsterisk())
		Assert(t, builder, "MERGE (test:`Test` {a: 'b'}) RETURN *")
	}
}

func TestNestedProperties(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	node := cypher.ANode("Test").NamedByString("test").WithRawProperties("outer", cypher.MapOf("a", cypher.LiteralOf("b")))
	builder = cypher.Match(node).
		Returning(cypher.AnAsterisk())
	Assert(t, builder, "MATCH (test:`Test` {outer: {a: 'b'}}) RETURN *")
}

func TestShouldNotRenderPropertiesInReturn(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	nodeWithProperties := bikeNode.WithRawProperties("a", cypher.LiteralOf("b"))
	builder = cypher.Match(nodeWithProperties, nodeWithProperties.RelationshipFrom(userNode, "OWNS")).
		ReturningByNamed(nodeWithProperties)
	Assert(t, builder, "MATCH (b:`Bike` {a: 'b'}), (b)<-[:`OWNS`]-(u:`User`) RETURN b")
}
