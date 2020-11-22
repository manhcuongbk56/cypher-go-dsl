package cypher_go_dsl

import (
 "fmt"
 "testing"
)

func testCreateTree(t *testing.T) {
 device, err := NewNode("Devices").Named("d").WithRawProperties("entity.id", "7d729555-0d61-46ae-ab79-ce43e72f751b")
 if err != nil {
  fmt.Print(err)
  return
 }
 customer := NewNode("Customer").Named("c")
 element := device.RelationshipFrom(customer, "HAS")
 MatchCondition(element).returning(customer.symbolicName).build90
}

func TestWhereType(t *testing.T) {
 var where interface{} = Where{}
 visitable, ok := where.(Visitable)
 fmt.Print(visitable)
 fmt.Print(ok)
}

