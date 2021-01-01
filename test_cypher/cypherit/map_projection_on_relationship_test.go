package cypherit

import (
	"fmt"
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestProjectionSimpleOnRelationship(t *testing.T) {
	var builder cypher.BuildableStatement
	p := cypher.NewNode("Person").NamedByString("p")
	m := cypher.NewNode("Movie").NamedByString("m")
	rel := p.RelationshipTo(m, "ACTED_IN").NamedByString("r")
	//
	builder = cypher.
		MatchElements(rel).
		Returning(rel.Project("__internalNeo4jId__", cypher.FunctionIdByRelationship(rel), "roles"))
	Assert(t, builder, "MATCH (p:`Person`)-[r:`ACTED_IN`]->(m:`Movie`) RETURN r{__internalNeo4jId__: id(r), .roles}")
}

func TestNestedOnRelationship(t *testing.T) {
	var builder cypher.BuildableStatement
	p := cypher.NewNode("Person").NamedByString("p")
	m := cypher.NewNode("Movie").NamedByString("m")
	rel := p.RelationshipTo(m, "ACTED_IN").NamedByString("r")
	//
	builder = cypher.
		MatchElements(rel).
		Returning(m.Project("title", "roles", rel.Project("__internalNeo4jId__", cypher.FunctionIdByRelationship(rel), "roles")))
	Assert(t, builder, "MATCH (p:`Person`)-[r:`ACTED_IN`]->(m:`Movie`) RETURN m{.title, roles: r{__internalNeo4jId__: id(r), .roles}}")
}

func TestAsterisk(t *testing.T) {
	var builder cypher.BuildableStatement
	n := cypher.AnyNodeNamed("n")
	//
	builder = cypher.
		MatchElements(n).
		Returning(n.Project(cypher.CypherAsterisk()))
	Assert(t, builder, "MATCH (n) RETURN n{.*}")
}


func TestProjectInvalid(t *testing.T) {
	var expect = "map projection create new content: unknown type cypher.FunctionInvocation cannot be used with an implicit name as map entry"
	n := cypher.AnyNodeNamed("n")
	mapProjection := n.Project(cypher.FunctionIdByNode(n))
	if mapProjection.GetError() == nil {
		t.Error("expect error but got nil")
		return
	}
	if mapProjection.GetError().Error() != expect {
		fmt.Printf("get error : %s but expect %s", mapProjection.GetError().Error(), expect)
		return
	}
	//
	mapProjection = n.Project("a", cypher.MapOf("a", cypher.LiteralOf("b")), cypher.FunctionIdByNode(n))
	if mapProjection.GetError() == nil {
		t.Error("expect error but got nil")
		return
	}
	if mapProjection.GetError().Error() != expect {
		fmt.Printf("get error : %s but expect %s", mapProjection.GetError().Error(), expect)
		return
	}
}

