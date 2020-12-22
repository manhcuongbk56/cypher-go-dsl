package cypher

import (
	"testing"
)

func TestCypherLiteralOf(t *testing.T) {
	literalSlice := make([]Literal, 0)
	literalSlice = append(literalSlice, StringLiteralCreate("test1"))
	literalSlice = append(literalSlice, NumberLiteralCreate1(1))
	literalSlice = append(literalSlice, BooleanLiteralCreate(true))
	type args struct {
		object interface{}
	}
	tests := []struct {
		name string
		args args
		want Literal
	}{
		{name: "testString", args: struct{ object interface{} }{object: "stringLiteral"}, want: StringLiteralCreate("stringLiteral")},
		{name: "testBoolean", args: struct{ object interface{} }{object: true}, want: BooleanLiteralCreate(true)},
		{name: "testNil", args: struct{ object interface{} }{object: nil}, want: NIL_INSTANCE},
		{name: "testNumber", args: struct{ object interface{} }{object: 12}, want: NumberLiteralCreate1(12)},
		{name: "testLiteralSlice", args: struct{ object interface{} }{object: literalSlice}, want: ListLiteralCreate(literalSlice)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CypherLiteralOf(tt.args.object); got.AsString() != tt.want.AsString() {
				t.Errorf("CypherLiteralOf() = %v, want %v", got.AsString(), tt.want.AsString())
			}
		})
	}
}
