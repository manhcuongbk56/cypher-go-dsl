package cypherit

import (
	"errors"
	"testing"

	"github.com/manhcuongbk56/cypher-go-dsl"
)

func TestListLiteralAcceptShouldWork(t *testing.T) {
	// Create content for the ListLiteral
	content := []cypher.Literal{
		cypher.LiteralOf(1),
		cypher.LiteralOf("test"),
		cypher.LiteralOf(true),
	}

	// Create a ListLiteral
	listLiteral := cypher.ListLiteralCreate(content)

	// Create a builder that uses the ListLiteral
	builder := cypher.CypherReturning(listLiteral)

	// Assert the rendered Cypher query
	expected := "RETURN [1, 'test', true]"
	Assert(t, builder, expected)
}

func TestListLiteralWithNestedContent(t *testing.T) {
	// Create nested ListLiteral content
	innerList := cypher.ListLiteralCreate([]cypher.Literal{
		cypher.LiteralOf(2),
		cypher.LiteralOf(3),
	})
	outerList := cypher.ListLiteralCreate([]cypher.Literal{
		cypher.LiteralOf(1),
		innerList,
		cypher.LiteralOf("nested"),
	})

	// Create a builder that uses the outer ListLiteral
	builder := cypher.CypherReturning(outerList)

	// Assert the rendered Cypher query
	expected := "RETURN [1, [2, 3], 'nested']"
	Assert(t, builder, expected)
}

func TestListLiteralErrorInBuilderShouldWork(t *testing.T) {
	// Create a ListLiteral with an error
	listLiteral := cypher.ListLiteralError(errors.New("test error"))

	// Ensure it doesn't render successfully
	builder := cypher.CypherReturning(listLiteral)
	_ ,err := builder.Build() // This should trigger an error
	if err == nil {
		t.Fatalf("Expected an error due to error in ListLiteral, but got nil")
	}
}
