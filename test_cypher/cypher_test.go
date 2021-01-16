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

func TestEscapeIfNecessary(t *testing.T) {
	inputs := make([][]string, 0)
	inputs = append(inputs, []string{" ", " "})
	inputs = append(inputs, []string{"a", "a"})
	inputs = append(inputs, []string{"ALabel", "ALabel"})
	inputs = append(inputs, []string{"A Label", "`A Label`"})
	inputs = append(inputs, []string{"`A `Label", "```A ``Label`"})
	inputs = append(inputs, []string{"Spring Data Neo4j⚡️RX", "`Spring Data Neo4j⚡️RX`"})
	for _, input := range inputs {
		escaped := cypher.EscapeIfNecessary(input[0])
		if escaped != input[1] {
			t.Errorf("escaped is not match:\n %s, expect is:\n %s", escaped, input[1])
		}
	}
}
