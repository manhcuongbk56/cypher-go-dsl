package cypher_go_dsl

import (
	"fmt"
	"testing"
)

func TestCreateTree(t *testing.T) {
	device, _ := NewNode("Devices").Named("d").WithRawProperties("entity.id", StringLiteral{content: "7d729555-0d61-46ae-ab79-ce43e72f751b"})
	customer := NewNode("Customer").Named("c")
	relation := device.RelationshipFrom(customer, "HAS")
	statement := Matchs(relation).
					returning(customer.symbolicName).
					Build()
	query := NewRenderer().Render(statement)
	fmt.Println(query)
}

func TestWhereType(t *testing.T) {
	var where interface{} = Where{}
	visitable, ok := where.(Visitable)
	fmt.Print(visitable)
	fmt.Print(ok)
}
