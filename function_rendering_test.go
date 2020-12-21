package cypher

import (
	"testing"
)

func TestCountInReturnClause(t *testing.T) {
	userNode := CypherNewNode("User").NamedByString("u")
	statement, err := CypherMatch(userNode).Returning(FunctionCount(userNode)).Build()
	if err != nil {
		t.Errorf("error when rendering statement: %s", err)
	}
	query := NewRenderer().Render(statement)
	if query != "MATCH (u:`User`) RETURN FunctionCount(u)" {
		t.Errorf("query is not match:\n %s", query)
	}
}
