package cypher_go_dsl

import "testing"

func TestGh48(t *testing.T) {
	n := CypherNewNode("Label").NamedByString("n")
	statement, err := CypherMatch(n).
		SetWithNamed(n, CypherMapOf("a", StringLiteralCreate("bar"), "b", StringLiteralCreate("baz"))).
		ReturningByNamed(n).
		Build()
	if err != nil {
		t.Errorf("error When build query: %s", err)
	}
	query := NewRenderer().Render(statement)
	if query != "MATCH (n:`Label`) SET n = {`a`: 'bar', `b`: 'baz'} RETURN n" {
		t.Errorf("Query is not match: %s", query)
	}
}

//func TestGh51(t *testing.T) {
//	n := CypherAnyNode1("n")
//	foobarProp := proper
//	statement, err := CypherMatch(n).
//		SetWithNamed(n, CypherMapOf("a", StringLiteralCreate("bar"), "b", StringLiteralCreate("baz"))).
//		ReturningByNamed(n).
//		Build()
//	if err != nil {
//		t.Errorf("error When build query: %s", err)
//	}
//	query := NewRenderer().Render(statement)
//	if query != "MATCH (n:`Label`) SET n = {`a`: 'bar', `b`: 'baz'} RETURN n" {
//		t.Errorf("Query is not match: %s", query)
//	}
//}
