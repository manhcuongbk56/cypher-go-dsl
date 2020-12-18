package cypher_go_dsl

import (
	"testing"
)

func TestRenderSimpleQuery(t *testing.T) {
	device := CypherNewNode("Device").NamedByString("d").WithRawProperties("entity.id", StringLiteral{content: "7d729555-0d61-46ae-ab79-ce43e72f751b", notNil: true})
	customer := CypherNewNode("Customer").NamedByString("c")
	relation := device.RelationshipTo(customer, "HAS")
	statement, _ := Matchs(relation).
		Returning(customer.symbolicName).
		Build()
	query := NewRenderer().Render(statement)
	if query != "MATCH (d:`Device` {`entity.id`: '7d729555-0d61-46ae-ab79-ce43e72f751b'})-[:`HAS`]->(c:`Customer`) RETURN c" {
		t.Errorf("query is not match:\n %s", query)
	}
}

func TestRenderComplexQuery(t *testing.T) {
	device := CypherNewNode("Farm").NamedByString("b")
	customer := CypherNewNode("Customer").NamedByString("c")
	relation := device.RelationshipTo(customer, "HAS")
	statement, _ := Matchs(relation).
		Returning(customer.symbolicName).
		Build()
	query := NewRenderer().Render(statement)
	if query != "MATCH (b:`Farm`)-[:`HAS`]->(c:`Customer`) RETURN c" {
		t.Errorf("query is not match:\n %s", query)
	}
}
