package cypherit

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func Assert(t *testing.T, buildableStatement cypher.BuildableStatement, expect string) {
	statement, err := buildableStatement.Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query := cypher.NewRenderer().Render(statement)
	if query != expect {
		t.Errorf("\n%s with length %d is incorrect, expect is \n%s with length %d", query, len(query), expect, len(expect))
	}
}
