package cypherit

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestFragmentOfStatementShouldBeReusable(t *testing.T) {
	personNode := cypher.NewNode("Person").NamedByString("p")
	ageProperty := personNode.Property("age")
	returning := cypher.MatchElements(personNode).
		ReturningByString("p")
	builder1 := returning.OrderBySortItem(ageProperty.Ascending()).Limit(1)
	builder2 := returning.OrderBySortItem(ageProperty.Descending()).Limit(1)
	Assert(t, builder1, "MATCH (p:`Person`) RETURN p ORDER BY p.age ASC LIMIT 1")
	Assert(t, builder2, "MATCH (p:`Person`) RETURN p ORDER BY p.age DESC LIMIT 1")
}

func TestAliasedFunctionsShouldNotBeRenderedTwiceInProjection(t *testing.T) {
	o := cypher.NewNode("Order").NamedByString("o")
	li := cypher.NewNode("LineItem").NamedByString("li")
	hasLineItems := o.RelationshipTo(li).NamedByString("h")
	netAmount := cypher.FunctionSum(li.Property("price").Multiply(li.Property("quantity")).Get()).As("netAmount")
	totalAmount := netAmount.Multiply(cypher.LiteralOf(1).Add(cypher.CypherParameter("taxRate")).Get()).As("totalAmount")
	returning := cypher.MatchElements(hasLineItems).
		WhereConditionContainer(o.Property("id").IsEqualTo(cypher.CypherParameter("id"))).
		With(o.GetRequiredSymbolicName(), netAmount.Get(), totalAmount.Get()).
		Returning(o.Project(o.Property("x"),
			netAmount.Get(),
			totalAmount.Get(),
			netAmount.Multiply(cypher.CypherParameter("taxRate")).As("taxAmount").Get()))
	Assert(t, returning, "MATCH (o:`Order`)-[h]->(li:`LineItem`) WHERE o.id = $id WITH o, sum((li.price * li.quantity)) AS netAmount, (netAmount * (1 + $taxRate)) AS totalAmount RETURN o{.x, netAmount: netAmount, totalAmount: totalAmount, taxAmount: (netAmount * $taxRate)}")
}
