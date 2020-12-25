package cypher

import (
	"math/big"
)

type NumberLiteral struct {
	ExpressionContainer
	content *big.Float
	key     string
	notNil  bool
	err     error
}

func NumberLiteralCreate(content int64) NumberLiteral {
	n := NumberLiteral{
		content: new(big.Float).SetInt(new(big.Int).SetInt64(content)),
		notNil:  true,
	}
	n.key = getAddress(&n)
	n.ExpressionContainer = ExpressionWrap(n)
	return n
}

func NumberLiteralCreate2(content float64) NumberLiteral {
	n := NumberLiteral{
		content: big.NewFloat(content),
		notNil:  true,
	}
	n.key = getAddress(&n)
	n.ExpressionContainer = ExpressionWrap(n)
	return n
}

func NumberLiteralCreate1(content int) NumberLiteral {
	n := NumberLiteral{
		content: new(big.Float).SetInt(new(big.Int).SetInt64(int64(content))),
		notNil:  true,
	}
	n.key = getAddress(&n)
	n.ExpressionContainer = ExpressionWrap(n)
	return n
}

func (n NumberLiteral) getError() error {
	return n.err
}

func (n NumberLiteral) isNotNil() bool {
	return n.notNil
}

func (n NumberLiteral) getKey() string {
	return n.key
}

func (n NumberLiteral) GetExpressionType() ExpressionType {
	return "NumberLiteral"
}

func (n NumberLiteral) GetContent() interface{} {
	return n.content
}

func (n NumberLiteral) AsString() string {
	return n.content.String()
}

func (n NumberLiteral) accept(visitor *CypherRenderer) {
	visitor.enter(n)
	visitor.leave(n)
}

func (n NumberLiteral) enter(renderer *CypherRenderer) {
	renderer.append(n.AsString())
}

func (n NumberLiteral) leave(renderer *CypherRenderer) {
}
