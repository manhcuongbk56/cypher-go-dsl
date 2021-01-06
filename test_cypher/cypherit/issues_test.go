package cypherit

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestGh70(t *testing.T) {
	var builder cypher.BuildableStatement
	strawberry := cypher.NewNodeWithProperties("Fruit", cypher.MapOf("kind", cypher.LiteralOf("strawberry"))).
		NamedByString("s")
	//
	builder = cypher.
		MatchElements(strawberry).
		Set(strawberry.Property("color").To(cypher.LiteralOf("red")))
	Assert(t, builder, "MATCH (s:`Fruit` {kind: 'strawberry'}) SET s.color = 'red'")
}

func TestGh167(t *testing.T) {
	var builder cypher.BuildableStatement
	app := cypher.NewNode("Location").NamedByString("app").WithRawProperties("uuid", cypher.Param("app_uuid"))
	locStart := cypher.NewNode("Location").NamedByString("loc_start")
	resume := cypher.NewNode("Resume").NamedByString("r")
	offer := cypher.NewNode("Offer").NamedByString("o")
	startN := cypher.NewNode("ResumeNode").NamedByString("start_n")
	aFl := app.RelationshipFrom(locStart, "PART_OF").Length(0, 3)
	lFr := locStart.RelationshipFrom(resume, "IN", "IN_ANALYTICS")
	builder = cypher.
		MatchElements(aFl, lFr).
		WithDistinctByNamed(resume, locStart, app).
		Match(resume.RelationshipTo(offer.WithRawProperties("is_valid", cypher.CypherLiteralTrue()), "IN_COHORT_OF").
			RelationshipTo(cypher.AnyNodeNamed("app"), "IN")).
		WithDistinctByNamed(resume, locStart, app, offer).
		Match(offer.RelationshipTo(startN, "FOR")).
		Where(cypher.FunctionIdByNode(startN).In(cypher.Param("start_ids")).Get()).
		ReturningDistinctByNamed(resume, locStart, app, offer, startN)
	Assert(t, builder, "MATCH (app:`Location` {uuid: $app_uuid})<-[:`PART_OF`*0..3]-(loc_start:`Location`), (loc_start)<-[:`IN`|`IN_ANALYTICS`]-(r:`Resume`) WITH DISTINCT r, "+
		"loc_start, app MATCH (r)-[:`IN_COHORT_OF`]->(o:`Offer` {is_valid: true})-[:`IN`]->(app) WITH DISTINCT r, "+
		"loc_start, app, o MATCH (o:`Offer`)-[:`FOR`]->(start_n:`ResumeNode`) WHERE id(start_n) IN $start_ids RETURN DISTINCT r, loc_start, app, o, start_n")
}

func TestGh174(t *testing.T) {
	var builder cypher.BuildableStatement
	r := cypher.NewNode("Resume").NamedByString("r")
	o := cypher.NewNode("Offer").NamedByString("o")
	//
	builder = cypher.
		MatchElements(r.RelationshipTo(o, "FOR")).
		Where(r.HasLabels("LastResume").Not().Get()).
		And(cypher.FunctionCoalesce(o.Property("valid_only"), cypher.CypherLiteralFalse()).IsEqualTo(cypher.CypherLiteralFalse()).
			And(r.HasLabels("InvalidStatus").Not().Get()).
			Or(o.Property("valid_only").IsTrue().And(r.HasLabels("InvalidStatus")).Get()).Get()).
		ReturningDistinctByNamed(r, o)
	Assert(t, builder, "MATCH (r:`Resume`)-[:`FOR`]->(o:`Offer`) WHERE (NOT (r:`LastResume`) AND (coalesce(o.valid_only, false) = false AND NOT (r:`InvalidStatus`) OR (o.valid_only = true AND r:`InvalidStatus`))) RETURN DISTINCT r, o")
}

func TestGh184(t *testing.T) {
	var builder cypher.BuildableStatement
	r := cypher.NewNode("Resume").NamedByString("r")
	u := cypher.NewNode("UserSearchable").NamedByString("u")
	o := cypher.NewNode("Offer").NamedByString("o")
	//
	builder = cypher.
		MatchElements(r.RelationshipFrom(u, "HAS")).
		Where(r.HasLabels("LastResume").Not().Get()).
		And(cypher.FunctionCoalesce(o.Property("valid_only"), cypher.CypherLiteralFalse()).IsEqualTo(cypher.CypherLiteralFalse()).
			And(r.HasLabels("InvalidStatus").Not().Get()).
			Or(o.Property("valid_only").IsTrue().And(r.HasLabels("ValidStatus")).Get()).Get()).
		And(r.Property("is_internship").IsTrue().
			And(cypher.FunctionSizeByPattern(r.RelationshipTo(cypher.AnyNode(), "PART_OF")).IsEmpty().Get()).
			Not().Get()).
		And(r.Property("is_sandwich_training").IsTrue().
			And(cypher.FunctionSizeByPattern(r.RelationshipTo(cypher.AnyNode(), "PART_OF")).IsEmpty().Get()).
			Not().Get()).
		ReturningDistinctByNamed(r, o)
	Assert(t, builder, "MATCH (r:`Resume`)<-[:`HAS`]-(u:`UserSearchable`) "+
		"WHERE (NOT (r:`LastResume`) "+
		"AND (coalesce(o.valid_only, false) = false "+
		"AND NOT (r:`InvalidStatus`) "+
		"OR (o.valid_only = true "+
		"AND r:`ValidStatus`)) "+
		"AND NOT ("+
		"(r.is_internship = true AND size(size((r)-[:`PART_OF`]->())) = 0)"+
		") "+
		"AND NOT ("+
		"(r.is_sandwich_training = true AND size(size((r)-[:`PART_OF`]->())) = 0)"+
		")"+
		") RETURN DISTINCT r, o")
}

func TestGh185(t *testing.T) {
	var builder cypher.BuildableStatement
	r := cypher.NewNode("Resume").NamedByString("r")
	u := cypher.NewNode("UserSearchable").NamedByString("u")
	//
	builder = cypher.
		MatchElements(r.RelationshipFrom(u, "HAS")).
		Where(cypher.ConditionsNot(cypher.PredicateExistsByPattern(r.RelationshipTo(u, "EXCLUDES")))).
		ReturningDistinctByNamed(r)
	Assert(t, builder, "MATCH (r:`Resume`)<-[:`HAS`]-(u:`UserSearchable`) WHERE NOT (exists((r)-[:`EXCLUDES`]->(u))) RETURN DISTINCT r")
}

func TestGh187(t *testing.T) {
	var builder cypher.BuildableStatement
	r := cypher.NewNode("Resume").NamedByString("r")
	u := cypher.NewNode("User").NamedByString("u")
	//
	builder = cypher.
		MatchElements(r.RelationshipFrom(u, "HAS")).
		With(cypher.FunctionHead(cypher.FunctionCollect(r.GetRequiredSymbolicName())).As("r").Get()).
		ReturningByNamed(r)
	Assert(t, builder, "MATCH (r:`Resume`)<-[:`HAS`]-(u:`User`) WITH head(collect(r)) AS r RETURN r")
}

func TestGh188(t *testing.T) {
	var builder cypher.BuildableStatement
	r := cypher.NewNode("Resume").NamedByString("r")
	u := cypher.NewNode("User").NamedByString("u")
	//
	builder = cypher.
		MatchElements(r.RelationshipFrom(u, "HAS")).
		Returning(cypher.FunctionCountDistinctByExpression(r.GetRequiredSymbolicName()).As("r").Get())
	Assert(t, builder, "MATCH (r:`Resume`)<-[:`HAS`]-(u:`User`) RETURN count(DISTINCT r) AS r")
}

func TestGh197(t *testing.T) {
	var builder cypher.BuildableStatement
	n := cypher.NewNode("Person").NamedByString("n")
	//AVG
	builder = cypher.
		MatchElements(n).
		Returning(cypher.FunctionAvg(n.Property("age")))
	Assert(t, builder, "MATCH (n:`Person`) RETURN avg(n.age)")
	//MAX/MIN
	list := cypher.ListOf(cypher.LiteralOf(1),
		cypher.LiteralOf("a"),
		cypher.LiteralOf(nil),
		cypher.LiteralOf(0.2),
		cypher.LiteralOf("b"),
		cypher.LiteralOf("1"),
		cypher.LiteralOf("99"))
	builder = cypher.CypherUnwind(list).
		As("val").
		Returning(cypher.FunctionMax(cypher.CypherName("val")))
	Assert(t, builder, "UNWIND [1, 'a', NULL, 0.2, 'b', '1', '99'] AS val RETURN max(val)")
	builder = cypher.CypherUnwind(list).
		As("val").
		Returning(cypher.FunctionMin(cypher.CypherName("val")))
	Assert(t, builder, "UNWIND [1, 'a', NULL, 0.2, 'b', '1', '99'] AS val RETURN min(val)")

	//percentileCont/percentileDisc
	builder = cypher.MatchElements(n).
		Returning(cypher.FunctionPercentileCont(n.Property("age"), 0.4))
	Assert(t, builder, "MATCH (n:`Person`) RETURN percentileCont(n.age, 0.4)")
	builder = cypher.MatchElements(n).
		Returning(cypher.FunctionPercentileDisc(n.Property("age"), 0.5))
	Assert(t, builder, "MATCH (n:`Person`) RETURN percentileDisc(n.age, 0.5)")

	//stDev/stDevP
	builder = cypher.MatchElements(n).
		Where(n.Property("name").In(cypher.ListOf(cypher.LiteralOf("A"), cypher.LiteralOf("B"), cypher.LiteralOf("C"))).Get()).
		Returning(cypher.FunctionStDev(n.Property("age")))
	Assert(t, builder, "MATCH (n:`Person`) WHERE n.name IN ['A', 'B', 'C'] RETURN stDev(n.age)")
	builder = cypher.MatchElements(n).
		Where(n.Property("name").In(cypher.ListOf(cypher.LiteralOf("A"), cypher.LiteralOf("B"), cypher.LiteralOf("C"))).Get()).
		Returning(cypher.FunctionStDevP(n.Property("age")))
	Assert(t, builder, "MATCH (n:`Person`) WHERE n.name IN ['A', 'B', 'C'] RETURN stDevP(n.age)")
	// sum
	builder = cypher.MatchElements(n).
		With(cypher.ListOf(cypher.MapOf("type", n.GetRequiredSymbolicName(), "nb", cypher.FunctionSum(n.GetRequiredSymbolicName()))).As("counts").Get()).
		Returning(cypher.FunctionSum(n.Property("age")))
	Assert(t, builder, "MATCH (n:`Person`) WITH [{type: n, nb: sum(n)}] AS counts RETURN sum(n.age)")
}

func TestGh200(t *testing.T) {
	var builder cypher.BuildableStatement
	r := cypher.NewNode("Resume").NamedByString("r")
	//
	builder = cypher.
		MatchElements(r).
		With(r.GetRequiredSymbolicName()).
		ReturningDistinct(r.GetRequiredSymbolicName())
	Assert(t, builder, "MATCH (r:`Resume`) WITH r RETURN DISTINCT r")
}

func TestGh204(t *testing.T) {
	var builder cypher.BuildableStatement
	a := cypher.NewNode("A").NamedByString("a")
	b := cypher.NewNode("B").NamedByString("b")
	c := cypher.NewNode("C").NamedByString("c")
	//
	builder = cypher.
		MatchElements(a.RelationshipTo(b).RelationshipTo(c).Max(2)).
		ReturningByNamed(a)
	Assert(t, builder, "MATCH (a:`A`)-->(b:`B`)-[*..2]->(c:`C`) RETURN a")
}

func TestGh245(t *testing.T) {
	var builder cypher.BuildableStatement
	n := cypher.NewNode("Person").NamedByString("p")
	expected := "MATCH (p:`Person`) RETURN p{alias: p.name}"
	//
	builder = cypher.
		MatchElements(n).
		Returning(n.Project("alias", n.Property("name")))
	Assert(t, builder, expected)
	//
	builder = cypher.
		MatchElements(n).
		Returning(n.Project(n.Property("name").As("alias").Get()))
	Assert(t, builder, expected)
}

func TestGh44(t *testing.T) {
	var builder cypher.BuildableStatement
	n := cypher.AnyNodeNamed("n")
	//
	builder = cypher.
		MatchElements(n).
		Returning(cypher.FunctionCollectDistinctByNamed(n).As("distinctNodes").Get())
	Assert(t, builder, "MATCH (n) RETURN collect(DISTINCT n) AS distinctNodes")
}

func TestGh84(t *testing.T) {
	var builder cypher.BuildableStatement
	parent := cypher.NewNode("Parent").NamedByString("parent")
	child := cypher.NewNode("Child").NamedByString("child")
	//
	builder = cypher.
		CypherCall("apoc.create.relationship").
		WithArgs(parent.GetRequiredSymbolicName(), cypher.LiteralOf("ChildEdge"), cypher.MapOf("score", cypher.LiteralOf(0.33), "weight", cypher.LiteralOf(1.7)), child.GetRequiredSymbolicName()).
		YieldString("rel")
	Assert(t, builder, "CALL apoc.create.relationship(parent, 'ChildEdge', {score: 0.33, weight: 1.7}, child) YIELD rel")
}

func TestAliasesShouldBeEscapedIfNecessary(t *testing.T) {
	var builder cypher.BuildableStatement
	alias := cypher.CypherName("n").As("das ist ein Alias").Get()
	//
	builder = cypher.
		MatchElements(cypher.AnyNode().NamedByString("n")).
		With(alias).
		Returning(alias)
	Assert(t, builder, "MATCH (n) WITH n AS `das ist ein Alias` RETURN `das ist ein Alias`")
}

func TestProjectedPropertiesShouldBeEscapedIfNecessary(t *testing.T) {
	var builder cypher.BuildableStatement
	node := cypher.AnyNode().NamedByString("n")
	//
	builder = cypher.
		MatchElements(node).
		Returning(node.Project("property 1", "property 2"))
	Assert(t, builder, "MATCH (n) RETURN n{.`property 1`, .`property 2`}")
}

func TestMapKeysShouldBeEscapedIfNecessary(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.
		CypherReturning(cypher.MapOf("key 1", cypher.CypherLiteralTrue(), "key 2", cypher.CypherLiteralFalse()))
	Assert(t, builder, "RETURN {`key 1`: true, `key 2`: false}")
}
