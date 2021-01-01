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
	query, _ := cypher.NewRenderer().Render(statement)
	if query != expect {
		t.Errorf("\n%s with length %d is incorrect, expect is \n%s with length %d", query, len(query), expect, len(expect))
	}
}

func AssertStatement(t *testing.T, statement cypher.Statement, expect string) {
	query, err := cypher.NewRenderer().Render(statement)
	if err != nil {
		t.Errorf("statement have error:\n %s", err)
		return
	}
	if query != expect {
		t.Errorf("\n%s with length %d is incorrect, expect is \n%s with length %d", query, len(query), expect, len(expect))
	}
}

func AssertStatementError(t *testing.T, statement cypher.Statement, expectError string) {
	_, err := cypher.NewRenderer().Render(statement)
	if err == nil {
		t.Errorf("expect statement have error but not")
		return
	}
	if err.Error() != expectError {
		t.Errorf("\n%s with length %d is incorrect, expect is \n%s with length %d", err.Error(), len(err.Error()), expectError, len(expectError))
	}
}
