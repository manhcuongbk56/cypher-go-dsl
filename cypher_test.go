package cypher_go_dsl

import (
	"fmt"
	"testing"
)

func TestRenderSimpleQuery(t *testing.T) {
	device, _ := NewNode("Device").Named("d").WithRawProperties("entity.id", StringLiteral{content: "7d729555-0d61-46ae-ab79-ce43e72f751b", notNil: true})
	customer := NewNode("Customer").Named("c")
	relation := device.RelationshipTo(customer, "HAS")
	statement, _ := Matchs(relation).
		returning(customer.symbolicName).
		build()
	query := NewRenderer().Render(statement)
	fmt.Println(query)
}

func TestRenderComplexQuery(t *testing.T) {
	device := NewNode("Farm").Named("b")
	customer := NewNode("Customer").Named("c")
	relation := device.RelationshipTo(customer, "HAS")
	statement, _ := Matchs(relation).
		returning(customer.symbolicName).
		build()
	query := NewRenderer().Render(statement)
	fmt.Println(query)
}

func TestExpressionContainer_Add(t *testing.T) {
	var property = PropertyLookup{
		propertyKeyName: "test",
	}
	a := property.EndsWith(property)
	fmt.Print(a)
}

func TestNodeAsNamedType(t *testing.T) {
	var node interface{} = Node{}
	a, b := node.(Named)
	fmt.Print(a)
	fmt.Print(b)
}
