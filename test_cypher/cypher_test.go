package test_cypher

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestRenderSimpleQuery(t *testing.T) {
	device := cypher.ANode("Device").NamedByString("d").WithRawProperties("entity.id", cypher.StringLiteralCreate("7d729555-0d61-46ae-ab79-ce43e72f751b"))
	customer := cypher.ANode("Customer").NamedByString("c")
	relation := device.RelationshipTo(customer, "HAS")
	statement, _ := cypher.Match(relation).
		Returning(customer.GetSymbolicName()).
		Build()
	query, _ := cypher.NewRenderer().Render(statement)
	if query != "MATCH (d:`Device` {entity.id: '7d729555-0d61-46ae-ab79-ce43e72f751b'})-[:`HAS`]->(c:`Customer`) RETURN c" {
		t.Errorf("query is not match:\n %s", query)
	}
}

func TestRenderComplexQuery(t *testing.T) {
	device := cypher.ANode("Farm").NamedByString("b")
	customer := cypher.ANode("Customer").NamedByString("c")
	relation := device.RelationshipTo(customer, "HAS")
	statement, _ := cypher.Match(relation).
		Returning(customer.GetSymbolicName()).
		Build()
	query, _ := cypher.NewRenderer().Render(statement)
	if query != "MATCH (b:`Farm`)-[:`HAS`]->(c:`Customer`) RETURN c" {
		t.Errorf("query is not match:\n %s", query)
	}
}

func TestWithCountFunction(t *testing.T) {

}
