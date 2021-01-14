package cypher

import (
	"fmt"
	"regexp"
)

type StringLiteral struct {
	ExpressionContainer
	content string
	key     string
	notNil  bool
	err     error
}

var RESERVED_CHARS, _ = regexp.Compile("([" + regexp.QuoteMeta("\\'\"") + "])")
var CHARACTER_MAP = map[string]string{
	"\\": "\\\\",
	"\"": "\\\"",
	"'":  "\\'",
}
var QUOTED_LITERAL_FORMAT = "'%s'"

func StringLiteralCreate(content string) StringLiteral {
	stringLiteral := StringLiteral{
		content: content,
		notNil:  true,
	}
	stringLiteral.key = getAddress(&stringLiteral)
	stringLiteral.ExpressionContainer = ExpressionWrap(stringLiteral)
	return stringLiteral
}

func StringLiteralError(err error) StringLiteral {
	return StringLiteralError(err)
}

func (s StringLiteral) GetError() error {
	return s.err
}

func (s StringLiteral) isNotNil() bool {
	return s.notNil
}

func escapeStringLiteral(value string) string {
	if value == "" {
		return value
	}
	return RESERVED_CHARS.ReplaceAllStringFunc(value, func(s string) string {
		return CHARACTER_MAP[s]
	})
}

func (s StringLiteral) getKey() string {
	return s.key
}

func (s StringLiteral) GetExpressionType() ExpressionType {
	return "StringLiteral"
}

func (s StringLiteral) GetContent() interface{} {
	return s.content
}

func (s StringLiteral) AsString() string {
	return fmt.Sprintf(QUOTED_LITERAL_FORMAT, escapeStringLiteral(s.content))
}

func (s StringLiteral) accept(visitor *CypherRenderer) {
	visitor.enter(s)
	visitor.leave(s)
}

func (s StringLiteral) enter(renderer *CypherRenderer) {
	renderer.append(s.AsString())
}

func (s StringLiteral) leave(renderer *CypherRenderer) {
}
