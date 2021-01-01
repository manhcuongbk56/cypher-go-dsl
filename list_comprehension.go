package cypher

import "errors"

type ListComprehension struct {
	ExpressionContainer
	variable       SymbolicName
	listExpression Expression
	where          Where
	listDefinition Expression
	key            string
	err            error
	notNil         bool
}

func ListComprehensionCreate(variable SymbolicName, listExpression Expression, where Where, listDefinition Expression) ListComprehension {
	if variable.GetError() != nil {
		return ListComprehensionError(variable.GetError())
	}
	list := ListComprehension{
		variable:       variable,
		listExpression: listExpression,
		where:          where,
		listDefinition: listDefinition,
		notNil:         true,
	}
	list.key = getAddress(&list)
	list.ExpressionContainer = ExpressionWrap(list)
	return list
}

func ListComprehensionError(err error) ListComprehension {
	return ListComprehension{
		err: err,
	}
}

func ListComprehensionWith(variable SymbolicName) OngoingDefinitionWithVariable {
	if !variable.isNotNil() {
		return ListComprehensionBuilder{err: errors.New("list comprehension: a variable is required")}
	}
	return ListComprehensionBuilderCreate(variable)
}

func (l ListComprehension) GetError() error {
	return l.err
}

func (l ListComprehension) accept(visitor *CypherRenderer) {
	visitor.enter(l)
	l.variable.accept(visitor)
	IN.accept(visitor)
	l.listExpression.accept(visitor)
	VisitIfNotNull(l.where, visitor)
	if l.listDefinition != nil {
		PIPE.accept(visitor)
		l.listDefinition.accept(visitor)
	}
	visitor.leave(l)
}

func (l ListComprehension) enter(renderer *CypherRenderer) {
	renderer.append("[")
}

func (l ListComprehension) leave(renderer *CypherRenderer) {
	renderer.append("]")
}

func (l ListComprehension) getKey() string {
	return l.key
}

func (l ListComprehension) isNotNil() bool {
	return l.notNil
}

func (l ListComprehension) GetExpressionType() ExpressionType {
	return "ListComprehension"
}

type OngoingDefinitionWithVariable interface {
	In(list Expression) OngoingDefinitionWithList
}

type OngoingDefinitionWithList interface {
	OngoingDefinitionWithoutReturn
	Where(condition Condition) OngoingDefinitionWithoutReturn
}

type OngoingDefinitionWithoutReturn interface {
	ReturningByNamed(variables ...Named) ListComprehension
	Returning(listDefinition ...Expression) ListComprehension
	ReturningDefault() ListComprehension
}

type ListComprehensionBuilder struct {
	variable       SymbolicName
	listExpression Expression
	where          Where
	err            error
}

func ListComprehensionBuilderCreate(variable SymbolicName) ListComprehensionBuilder {
	return ListComprehensionBuilder{
		variable: variable,
	}
}

func (l ListComprehensionBuilder) In(list Expression) OngoingDefinitionWithList {
	l.listExpression = list
	return l
}

func (l ListComprehensionBuilder) ReturningByNamed(variables ...Named) ListComprehension {
	return l.Returning(CreateSymbolicNameByNamed(variables...)...)
}

func (l ListComprehensionBuilder) Returning(expressions ...Expression) ListComprehension {
	return ListComprehensionCreate(l.variable, l.listExpression, l.where,
		ListOrSingleExpression(expressions...))
}

func (l ListComprehensionBuilder) ReturningDefault() ListComprehension {
	return ListComprehensionCreate(l.variable, l.listExpression, l.where, nil)
}

func (l ListComprehensionBuilder) Where(condition Condition) OngoingDefinitionWithoutReturn {
	l.where = WhereCreate(condition)
	return l
}
