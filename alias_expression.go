package cypher_go_dsl

import (
	"errors"
)

type AliasedExpression struct {
	delegate Expression
	alias    string
	key      string
	notNil   bool
	err      error
}

func AliasedExpressionCreate(delegate Expression, alias string) AliasedExpression {
	if delegate == nil {
		return AliasedExpressionError(errors.New("expression to alias can't be nil"))
	}
	if delegate.getError() != nil {
		return AliasedExpressionError(delegate.getError())
	}
	if alias == "" {
		return AliasedExpressionError(errors.New("the alias may not be empty"))
	}
	a := AliasedExpression{
		delegate: delegate,
		alias:    alias,
		notNil:   true,
	}
	a.key = getAddress(&a)
	return a
}

func AliasedExpressionError(err error) AliasedExpression {
	return AliasedExpression{
		err: err,
	}
}

func (aliased AliasedExpression) getError() error {
	return aliased.err
}

func (aliased AliasedExpression) isNotNil() bool {
	return aliased.notNil
}

func (aliased AliasedExpression) GetExpressionType() ExpressionType {
	return EXPRESSION
}

func (aliased AliasedExpression) As(newAlias string) AliasedExpression {
	if newAlias == "" {
		return AliasedExpressionError(errors.New("the alias may not be empty"))
	}
	return AliasedExpressionCreate(aliased.delegate, newAlias)
}

func (aliased AliasedExpression) accept(visitor *CypherRenderer) {
	(*visitor).enter(aliased)
	NameOrExpression(aliased.delegate).accept(visitor)
	(*visitor).leave(aliased)
}

func (aliased AliasedExpression) getKey() string {
	return aliased.key
}

func (aliased AliasedExpression) enter(renderer *CypherRenderer) {
	if _, visited := renderer.visitableToAliased[aliased.key]; visited {
		renderer.append(escapeName(aliased.alias))
	}
}

func (aliased AliasedExpression) leave(renderer *CypherRenderer) {
	if _, visited := renderer.visitableToAliased[aliased.key]; !visited {
		renderer.append(" AS ")
		renderer.append(escapeName(aliased.alias))
	}
}
