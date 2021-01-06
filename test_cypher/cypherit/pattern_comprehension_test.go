package cypherit

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestPatternComprehensionsSimple(t *testing.T) {
	var builder cypher.BuildableStatement
	a := cypher.NewNode("Person").WithRawProperties("name", cypher.LiteralOf("Keanu Reeves")).NamedByString("a")
	b := cypher.AnyNodeNamed("b")
	//
	builder = cypher.
		MatchElements(a).
		Returning(cypher.ListBasedOn(a.RelationshipBetween(b)).Returning(b.Property("released")).As("years").Get())
	Assert(t, builder, "MATCH (a:`Person` {name: 'Keanu Reeves'}) RETURN [(a)--(b) | b.released] AS years")
}

func TestPatternComprehensionsSimpleWithWhere(t *testing.T) {
	var builder cypher.BuildableStatement
	a := cypher.NewNode("Person").WithRawProperties("name", cypher.LiteralOf("Keanu Reeves")).NamedByString("a")
	b := cypher.AnyNodeNamed("b")
	//
	builder = cypher.
		MatchElements(a).
		Returning(cypher.ListBasedOn(a.RelationshipBetween(b)).Where(b.HasLabels("Movie")).Returning(b.Property("released")).As("years").Get())
	Assert(t, builder, "MATCH (a:`Person` {name: 'Keanu Reeves'}) RETURN [(a)--(b) WHERE b:`Movie` | b.released] AS years")
}

func TestPatternComprehensionsNested(t *testing.T) {
	var builder cypher.BuildableStatement
	n := cypher.NewNode("Person").NamedByString("n")
	o1 := cypher.NewNode("Organisation").NamedByString("o1")
	l1 := cypher.NewNode("Location").NamedByString("l1")
	p2 := cypher.NewNode("Person").NamedByString("p2")

	r_f1 := n.RelationshipTo(o1, "FOUNDED").NamedByString("r_f1")
	r_e1 := n.RelationshipTo(o1, "EMPLOYED_BY").NamedByString("r_e1")
	r_l1 := n.RelationshipTo(l1, "LIVES_AT").NamedByString("r_l1")
	r_l2 := l1.RelationshipFrom(p2, "LIVES_AT").NamedByString("r_l2")
	//
	builder = cypher.
		MatchElements(n).
		Returning(n.GetRequiredSymbolicName(),
			cypher.ListOf(cypher.ListBasedOn(r_f1).ReturningByNamed(r_f1, o1),
				cypher.ListBasedOn(r_e1).ReturningByNamed(r_e1, o1),
				cypher.ListBasedOn(r_l1).Returning(r_l1.GetRequiredSymbolicName(), l1.GetRequiredSymbolicName(), cypher.ListOf(cypher.ListBasedOn(r_l2).ReturningByNamed(r_l2, p2)))))
	Assert(t, builder, "MATCH (n:`Person`) RETURN n, [[(n)-[r_f1:`FOUNDED`]->(o1:`Organisation`) | [r_f1, o1]], [(n)-[r_e1:`EMPLOYED_BY`]->(o1) | [r_e1, o1]], [(n)-[r_l1:`LIVES_AT`]->(l1:`Location`) | [r_l1, l1, [[(l1)<-[r_l2:`LIVES_AT`]-(p2:`Person`) | [r_l2, p2]]]]]]")
}
