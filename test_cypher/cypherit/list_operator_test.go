package cypherit

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestValueAtShouldWork(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.CypherReturning(cypher.CypherValueAt(cypher.FunctionRangeRaw(0, 10), 3))

	Assert(t, builder, "RETURN range(0, 10)[3]")
}

func TestSubListUntilShouldWork(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.CypherReturning(cypher.CypherSubListUntil(cypher.FunctionRangeRaw(0, 10), 3))

	Assert(t, builder, "RETURN range(0, 10)[..3]")
}

func TestSubListFromShouldWork(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.CypherReturning(cypher.CypherSubListFrom(cypher.FunctionRangeRaw(0, 10), -3))

	Assert(t, builder, "RETURN range(0, 10)[-3..]")
}

func TestSubListShouldWork(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.CypherReturning(cypher.CypherSubList(cypher.FunctionRangeRaw(0, 10), 2, 4))

	Assert(t, builder, "RETURN range(0, 10)[2..4]")
}

func TestShouldWorkWithMapProjections(t *testing.T) {
	var builder cypher.BuildableStatement
	person := cypher.NewNode("Person").NamedByString("person")
	location := cypher.NewNode("Location").NamedByString("personLivesIn")
	//
	builder = cypher.
		MatchElements(person).
		Returning(person.Project("livesIn", cypher.CypherValueAt(cypher.ListBasedOn(person.RelationshipTo(location, "LIVES_IN")).
			Returning(location.Project("name")), 0)))

	Assert(t, builder, "MATCH (person:`Person`) RETURN person{livesIn: [(person)-[:`LIVES_IN`]->(personLivesIn:`Location`) | personLivesIn{.name}][0]}")
}

func TestShouldSupportExpressions(t *testing.T) {
	var builder cypher.BuildableStatement
	person := cypher.NewNode("Person").NamedByString("person")
	location := cypher.NewNode("Location").NamedByString("personLivesIn")
	//
	builder = cypher.
		MatchElements(person).
		Returning(person.Project("livesIn",
			cypher.SubList(cypher.ListBasedOn(person.RelationshipTo(location, "LIVES_IN")).
				Returning(location.Project("name")), cypher.CypherParameter("personLivedInOffset"),
				cypher.CypherParameter("personLivedInOffset").Add(cypher.CypherParameter("personLivedInFirst")).Get())))

	Assert(t, builder, "MATCH (person:`Person`) RETURN person{livesIn: [(person)-[:`LIVES_IN`]->(personLivesIn:`Location`) | personLivesIn{.name}][$personLivedInOffset..($personLivedInOffset + $personLivedInFirst)]}")
}
