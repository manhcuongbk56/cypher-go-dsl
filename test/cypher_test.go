package test

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestRenderSimpleQuery(t *testing.T) {
	device := cypher.NewNode("Device").NamedByString("d").WithRawProperties("entity.id", cypher.StringLiteralCreate("7d729555-0d61-46ae-ab79-ce43e72f751b"))
	customer := cypher.NewNode("Customer").NamedByString("c")
	relation := device.RelationshipTo(customer, "HAS")
	statement, _ := cypher.MatchElements(relation).
		Returning(customer.GetSymbolicName()).
		Build()
	query := cypher.NewRenderer().Render(statement)
	if query != "MATCH (d:`Device` {`entity.id`: '7d729555-0d61-46ae-ab79-ce43e72f751b'})-[:`HAS`]->(c:`Customer`) RETURN c" {
		t.Errorf("query is not match:\n %s", query)
	}
}

func TestRenderComplexQuery(t *testing.T) {
	device := cypher.NewNode("Farm").NamedByString("b")
	customer := cypher.NewNode("Customer").NamedByString("c")
	relation := device.RelationshipTo(customer, "HAS")
	statement, _ := cypher.MatchElements(relation).
		Returning(customer.GetSymbolicName()).
		Build()
	query := cypher.NewRenderer().Render(statement)
	if query != "MATCH (b:`Farm`)-[:`HAS`]->(c:`Customer`) RETURN c" {
		t.Errorf("query is not match:\n %s", query)
	}
}

func TestWithCountFunction(t *testing.T) {

}
