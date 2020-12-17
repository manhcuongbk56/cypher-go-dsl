package cypher_go_dsl

import (
	"fmt"
	"testing"
)

func TestRenderSimpleQuery(t *testing.T) {
	device := NewNode("Device").Named("d").WithRawProperties("entity.id", StringLiteral{content: "7d729555-0d61-46ae-ab79-ce43e72f751b", notNil: true})
	customer := NewNode("Customer").Named("c")
	relation := device.RelationshipTo(customer, "HAS")
	statement, _ := Matchs(relation).
		Returning(customer.symbolicName).
		Build()
	query := NewRenderer().Render(statement)
	fmt.Println(query)
}

func TestRenderComplexQuery(t *testing.T) {
	device := NewNode("Farm").Named("b")
	customer := NewNode("Customer").Named("c")
	relation := device.RelationshipTo(customer, "HAS")
	statement, _ := Matchs(relation).
		Returning(customer.symbolicName).
		Build()
	query := NewRenderer().Render(statement)
	fmt.Println(query)
}

func TestGh48(t *testing.T) {
	n := NewNode("Label").Named("n")
	statement, err := Matchs(n).
		SetWithNamed(n, MapOf("a", StringLiteralCreate("bar"), "b", StringLiteralCreate("baz"))).
		ReturningByNamed(n).
		Build()
	if err != nil {
		fmt.Print(err)
	}
	query := NewRenderer().Render(statement)
	fmt.Println(query)
}
