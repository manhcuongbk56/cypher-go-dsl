package cypherit

import (
	"fmt"
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestEqualWithStringLiteralAAA(t *testing.T) {
	customer := cypher.ANode("Customer").NamedByString("c")
	farm := cypher.ANode("Farm").NamedByString("f")
	barn := cypher.ANode("Barn").NamedByString("b")
	silo := cypher.ANode("Silo").NamedByString("s")
	silo2 := cypher.ANode("Silo").NamedByString("s")
	siloDimension := cypher.ANode("SiloDimension").NamedByString("sd")
	supplier := cypher.ANode("Supplier").NamedByString("su")
	device := cypher.ANode("Device").NamedByString("dv")
	r1 := customer.RelationshipTo(farm, "HAS").RelationshipTo(barn, "HAS").RelationshipTo(silo, "HAS").
		RelationshipFrom(siloDimension, "SIZES")
	r2 := supplier.RelationshipTo(silo, "SUPPLIES")
	r3 := silo2.RelationshipTo(device, "HAS")
	statement, err := cypher.
		Match(r1).
		Match(r2).
		Match(r3).
		WhereConditionContainer(silo.Property("entity.id").IsEqualTo(cypher.LiteralOf("b65f60bf-d05d-4ca4-9146-fe4bc58bbc3c"))).
		ReturningByNamed(customer, farm, barn, silo, siloDimension, supplier, device).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	fmt.Println(query)
}

func TestEqualWithStringLiteral(t *testing.T) {
	statement, err := cypher.
		Match(userNode).
		WhereConditionContainer(userNode.Property("name").IsEqualTo(cypher.LiteralOf("Test"))).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) WHERE u.name = 'Test' RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestEqualWithNumberLiteral(t *testing.T) {
	statement, err := cypher.
		Match(userNode).
		WhereConditionContainer(userNode.Property("age").IsEqualTo(cypher.LiteralOf(21))).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) WHERE u.age = 21 RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}
