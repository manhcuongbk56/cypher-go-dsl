package cypherit

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestEqualWithStringLiteral(t *testing.T) {
	statement, err := cypher.
		MatchElements(userNode).
		WhereConditionContainer(userNode.Property("name").IsEqualTo(cypher.LiteralOf("Test"))).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) WHERE u.name = 'Test' RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestEqualWithNumberLiteral(t *testing.T) {
	statement, err := cypher.
		MatchElements(userNode).
		WhereConditionContainer(userNode.Property("age").IsEqualTo(cypher.LiteralOf(21))).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) WHERE u.age = 21 RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}
