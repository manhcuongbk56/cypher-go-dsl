package cypherit

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestShouldRenderSetAfterCreate(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.Create(userNode).
		Set(userNode.Property("p").To(cypher.LiteralOf("Hallo, Welt")))
	Assert(t, builder, "CREATE (u:`User`) SET u.p = 'Hallo, Welt'")
}

func TestShouldRenderSetAfterMerge(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.Merge(userNode).
		Set(userNode.Property("p").To(cypher.LiteralOf("Hallo, Welt")))
	Assert(t, builder, "MERGE (u:`User`) SET u.p = 'Hallo, Welt'")
}

func TestShouldRenderSetAfterCreateAndWith(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.Create(userNode).
		WithByNamed(userNode).
		Set(userNode.Property("p").To(cypher.LiteralOf("Hallo, Welt")))
	Assert(t, builder, "CREATE (u:`User`) WITH u SET u.p = 'Hallo, Welt'")
}

func TestShouldRenderSetAfterMergeAndWith(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.Merge(userNode).
		WithByNamed(userNode).
		Set(userNode.Property("p").To(cypher.LiteralOf("Hallo, Welt")))
	Assert(t, builder, "MERGE (u:`User`) WITH u SET u.p = 'Hallo, Welt'")
}

func TestShouldRenderSet(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.Match(userNode).
		Set(userNode.Property("p").To(cypher.LiteralOf("Hallo, Welt")))
	Assert(t, builder, "MATCH (u:`User`) SET u.p = 'Hallo, Welt'")
	//
	builder = cypher.Match(userNode).
		Set(userNode.Property("p").To(cypher.LiteralOf("Hallo, Welt"))).
		Set(userNode.Property("a").To(cypher.LiteralOf("Selber hallo.")))
	Assert(t, builder, "MATCH (u:`User`) SET u.p = 'Hallo, Welt' SET u.a = 'Selber hallo.'")
	//
	builder = cypher.Match(userNode).
		Set(userNode.Property("p").To(cypher.LiteralOf("Hallo")), userNode.Property("g").To(cypher.LiteralOf("Welt")))
	Assert(t, builder, "MATCH (u:`User`) SET u.p = 'Hallo', u.g = 'Welt'")
}

func TestShouldRenderSetOnNodes(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.Match(userNode).
		SetByNode(userNode, "A", "B").
		ReturningByNamed(userNode)
	Assert(t, builder, "MATCH (u:`User`) SET u:`A`:`B` RETURN u")
	//
	builder = cypher.Match(userNode).
		WithByNamed(userNode).
		SetByNode(userNode, "A", "B").
		SetByNode(userNode, "C", "D").
		ReturningByNamed(userNode)
	Assert(t, builder, "MATCH (u:`User`) WITH u SET u:`A`:`B` SET u:`C`:`D` RETURN u")
}

func TestShouldRenderSetFromAListOfExpression(t *testing.T) {
	var builder cypher.BuildableStatement

	builder = cypher.Match(userNode).
		Set(userNode.Property("p"), cypher.LiteralOf("Hallo, Welt"))
	Assert(t, builder, "MATCH (u:`User`) SET u.p = 'Hallo, Welt'")
	//
	builder = cypher.Match(userNode).
		Set(userNode.Property("p"), cypher.LiteralOf("Hallo"), userNode.Property("g"), cypher.LiteralOf("Welt"))
	Assert(t, builder, "MATCH (u:`User`) SET u.p = 'Hallo', u.g = 'Welt'")
	//
	builder = cypher.Match(userNode).
		Set(userNode.Property("p"), cypher.LiteralOf("Hallo, Welt")).
		Set(userNode.Property("p"), cypher.LiteralOf("Hallo"), userNode.Property("g"), cypher.LiteralOf("Welt"))
	Assert(t, builder, "MATCH (u:`User`) SET u.p = 'Hallo, Welt' SET u.p = 'Hallo', u.g = 'Welt'")
	//
	builder = cypher.Match(userNode).
		Set(userNode.Property("p"))
	_, err := builder.Build()
	if err == nil {
		t.Errorf("expect error but success")
	}
	if err.Error() != "the list of expression to OperationSet must be even" {
		t.Errorf("wrong error")
	}
}

func TestShouldRenderMixedSet(t *testing.T) {
	var builder cypher.BuildableStatement
	builder = cypher.Match(userNode).
		Set(userNode.Property("p1"), cypher.LiteralOf("Two expressions")).
		Set(userNode.Property("p2"), cypher.LiteralOf("A set expression")).
		Set(userNode.Property("p3"), cypher.LiteralOf("One of two set expression"),
			userNode.Property("p4"), cypher.LiteralOf("Two of two set expression")).
		Set(userNode.Property("p5"), cypher.LiteralOf("Pair one of 2 expressions"),
			userNode.Property("p6"), cypher.LiteralOf("Pair two of 4 expressions")).
		Returning(cypher.AnAsterisk())
	Assert(t, builder, "MATCH (u:`User`) SET u.p1 = 'Two expressions' SET u.p2 = 'A set expression' SET u.p3 = 'One of two set expression', u.p4 = 'Two of two set expression' SET u.p5 = 'Pair one of 2 expressions', u.p6 = 'Pair two of 4 expressions' RETURN *")
}
