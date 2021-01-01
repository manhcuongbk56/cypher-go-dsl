package cypherit

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestUnwindWithoutWith(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	rootNode := cypher.AnyNodeNamed("n")
	label := cypher.CypherName("label")
	builder = cypher.MatchElements(rootNode).
		WhereConditionContainer(rootNode.InternalId().IsEqualTo(cypher.LiteralOf(1))).
		Unwind(rootNode.Labels()).
		As("label").
		With(label).
		Where(label.In(cypher.CypherParameter("fixedLabels")).Not().Get()).
		Returning(cypher.FunctionCollect(label).As("labels").Get())
	Assert(t, builder, "MATCH (n) WHERE id(n) = 1 UNWIND labels(n) AS label WITH label WHERE NOT (label IN "+
		"$fixedLabels) RETURN collect(label) AS labels")
}

func TestShouldRenderLeadingUnwind(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.CypherUnwindMulti(cypher.LiteralOf(1), cypher.CypherLiteralTrue(), cypher.CypherLiteralFalse()).
		As("n").Returning(cypher.CypherName("n"))
	Assert(t, builder, "UNWIND [1, true, false] AS n RETURN n")
}

func TestShouldRenderLeadingUnwindWithUpdate(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.CypherUnwindMulti(cypher.LiteralOf(1), cypher.CypherLiteralTrue(), cypher.CypherLiteralFalse()).
		As("n").
		Merge(bikeNode.WithRawProperties("b", cypher.CypherName("n"))).
		ReturningByNamed(bikeNode)
	Assert(t, builder, "UNWIND [1, true, false] AS n MERGE (b:`Bike` {b: n}) RETURN b")
}

func TestShouldRenderLeadingUnwindWithCreate(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.CypherUnwindMulti(cypher.LiteralOf(1), cypher.CypherLiteralTrue(), cypher.CypherLiteralFalse()).
		As("n").
		Create(bikeNode.WithRawProperties("b", cypher.CypherName("n"))).
		ReturningByNamed(bikeNode)
	Assert(t, builder, "UNWIND [1, true, false] AS n CREATE (b:`Bike` {b: n}) RETURN b")
}

func TestShouldRenderUnwind(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	collected := cypher.FunctionCollectByNamed(bikeNode).As("collected").Get()
	builder = cypher.MatchElements(bikeNode).
		With(collected).
		Unwind(collected).As("x").
		WithByString("x").
		Delete(cypher.CypherName("x")).
		ReturningByString("x")
	Assert(t, builder, "MATCH (b:`Bike`) WITH collect(b) AS collected UNWIND collected AS x WITH x DELETE x "+
		"RETURN x")
}
