package cypher_go_dsl

import (
	"fmt"
	"testing"
)

func testCreateTree(t *testing.T) {
	device, _ := NewNode("Devices").Named("d").WithRawProperties("entity.id", "7d729555-0d61-46ae-ab79-ce43e72f751b")
	customer := NewNode("Customer").Named("c")
	relation := device.RelationshipFrom(customer, "HAS")
	statement := Matchs(relation).
		returning(customer.symbolicName).Build()
}

func TestWhereType(t *testing.T) {
	var where interface{} = Where{}
	visitable, ok := where.(Visitable)
	fmt.Print(visitable)
	fmt.Print(ok)
}
