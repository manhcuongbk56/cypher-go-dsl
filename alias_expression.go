package cypher_go_dsl

import (
	"errors"
	"fmt"
)

type AliasedExpression struct {
	delegate Expression
	alias    string
	key      string
	notNil   bool
	err      error
}

func (aliased AliasedExpression) getError() error {
	return aliased.err
}

func AliasedExpressionCreate(delegate Expression, alias string) AliasedExpression {
	return AliasedExpression{
		delegate: delegate,
		alias:    alias,
		notNil:   true,
	}
}

func (aliased AliasedExpression) isNotNil() bool {
	return aliased.notNil
}

func (aliased AliasedExpression) GetExpressionType() ExpressionType {
	return EXPRESSION
}

func (aliased AliasedExpression) As(newAlias string) AliasedExpression {
	if newAlias == "" {
		return AliasedExpression{err: errors.New("the alias may not be empty")}
	}
	return AliasedExpressionCreate(aliased.delegate, newAlias)
}

func (aliased AliasedExpression) accept(visitor *CypherRenderer) {
	aliased.key = fmt.Sprint(&aliased)
	(*visitor).enter(aliased)
	NameOrExpression(aliased.delegate).accept(visitor)
	(*visitor).leave(aliased)
}

func (aliased AliasedExpression) getKey() string {
	return aliased.key
}

func (aliased AliasedExpression) enter(renderer *CypherRenderer) {
	panic("implement me")
}

func (aliased AliasedExpression) leave(renderer *CypherRenderer) {
	panic("implement me")
}
