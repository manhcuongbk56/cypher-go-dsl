package cypherit

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestValueAtShouldWork(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.CypherReturning(cypher.CypherValueAt(cypher.RangeRaw(0, 10), 3))

	Assert(t, builder, "RETURN range(0, 10)[3]")
}

func TestSubListUntilShouldWork(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.CypherReturning(cypher.ASubListUntil(cypher.RangeRaw(0, 10), 3))

	Assert(t, builder, "RETURN range(0, 10)[..3]")
}

func TestSubListFromShouldWork(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.CypherReturning(cypher.CypherSubListFrom(cypher.RangeRaw(0, 10), -3))

	Assert(t, builder, "RETURN range(0, 10)[-3..]")
}

func TestSubListShouldWork(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.CypherReturning(cypher.ASubList(cypher.RangeRaw(0, 10), 2, 4))

	Assert(t, builder, "RETURN range(0, 10)[2..4]")
}

func TestShouldWorkWithMapProjections(t *testing.T) {
	var builder cypher.BuildableStatement
	person := cypher.ANode("Person").NamedByString("person")
	location := cypher.ANode("Location").NamedByString("personLivesIn")
	//
	builder = cypher.
		Match(person).
		Returning(person.Project("livesIn", cypher.CypherValueAt(cypher.ListBasedOn(person.RelationshipTo(location, "LIVES_IN")).
			Returning(location.Project("name")), 0)))

	Assert(t, builder, "MATCH (person:`Person`) RETURN person{livesIn: [(person)-[:`LIVES_IN`]->(personLivesIn:`Location`) | personLivesIn{.name}][0]}")
}

func TestShouldSupportExpressions(t *testing.T) {
	var builder cypher.BuildableStatement
	person := cypher.ANode("Person").NamedByString("person")
	location := cypher.ANode("Location").NamedByString("personLivesIn")
	//
	builder = cypher.
		Match(person).
		Returning(person.Project("livesIn",
			cypher.SubList(cypher.ListBasedOn(person.RelationshipTo(location, "LIVES_IN")).
				Returning(location.Project("name")), cypher.AParam("personLivedInOffset"),
				cypher.AParam("personLivedInOffset").Add(cypher.AParam("personLivedInFirst")).Get())))

	Assert(t, builder, "MATCH (person:`Person`) RETURN person{livesIn: [(person)-[:`LIVES_IN`]->(personLivesIn:`Location`) | personLivesIn{.name}][$personLivedInOffset..($personLivedInOffset + $personLivedInFirst)]}")
}
