package cypherit

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestShouldRenderRemoveOnNodes(t *testing.T) {
	builder := cypher.Match(userNode).
		RemoveByNode(userNode, "A", "B").
		ReturningByNamed(userNode)
	Assert(t, builder, "MATCH (u:`User`) REMOVE u:`A`:`B` RETURN u")
	//
	builder = cypher.Match(userNode).
		WithByNamed(userNode).
		SetByNode(userNode, "A", "B").
		RemoveByNode(userNode, "C", "D").
		ReturningByNamed(userNode)
	Assert(t, builder, "MATCH (u:`User`) WITH u SET u:`A`:`B` REMOVE u:`C`:`D` RETURN u")
}

func TestShouldRenderRemoveOfProperties(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.Match(userNode).
		Remove(userNode.Property("a"), userNode.Property("b")).
		ReturningByNamed(userNode)
	Assert(t, builder, "MATCH (u:`User`) REMOVE u.a, u.b RETURN u")
	//
	builder = cypher.Match(userNode).
		WithByNamed(userNode).
		Remove(userNode.Property("a")).
		Remove(userNode.Property("b")).
		ReturningByNamed(userNode)
	Assert(t, builder, "MATCH (u:`User`) WITH u REMOVE u.a REMOVE u.b RETURN u")
}
