package cypher_go_dsl

type ListComprehension struct {
	variable       SymbolicName
	listExpression Expression
	where          Where
	listDefinition Expression
	key            string
	err            error
	notNil         bool
}

func ListComprehensionCreate(variable SymbolicName, listExpression Expression, where Where, listDefinition Expression) ListComprehension {
	if variable.getError() != nil {
		return ListComprehensionError(variable.getError())
	}
	list := ListComprehension{
		variable:       variable,
		listExpression: listExpression,
		where:          where,
		listDefinition: listExpression,
		notNil:         true,
	}
	list.key = getAddress(&list)
	return list
}

func ListComprehensionError(err error) ListComprehension {
	return ListComprehension{
		err: err,
	}
}

func (l ListComprehension) getError() error {
	return l.err
}

func (l ListComprehension) accept(visitor *CypherRenderer) {
	visitor.enter(l)
	l.variable.accept(visitor)
	IN.accept(visitor)
	l.listExpression.accept(visitor)
	VisitIfNotNull(l.where, visitor)
	if l.listExpression != nil {
		PIPE.accept(visitor)
		l.listExpression.accept(visitor)
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
	in(list Expression) OngoingDefinitionWithList
}

type OngoingDefinitionWithList interface {
	OngoingDefinitionWithoutReturn
	whereByCondition(condition Condition) OngoingDefinitionWithoutReturn
}

type OngoingDefinitionWithoutReturn interface {
	returningByNamed(variables ...Named) ListComprehension
	returning(listDefinition ...Expression) ListComprehension
	returningDefault() ListComprehension
}

type ListComprehensionBuilder struct {
	variable       SymbolicName
	listExpression Expression
	where          Where
}

func (l ListComprehensionBuilder) in(list Expression) OngoingDefinitionWithList {
	l.listExpression = list
	return l
}

func (l ListComprehensionBuilder) returningByNamed(variables ...Named) ListComprehension {
	return l.returning(CreateSymbolicNameByNamed(variables...)...)
}

func (l ListComprehensionBuilder) returning(expressions ...Expression) ListComprehension {
	return ListComprehensionCreate(l.variable, l.listExpression, l.where,
		ListOrSingleExpression(expressions...))
}

func (l ListComprehensionBuilder) returningDefault() ListComprehension {
	return ListComprehensionCreate(l.variable, l.listExpression, l.where, nil)
}

func (l ListComprehensionBuilder) whereByCondition(condition Condition) OngoingDefinitionWithoutReturn {
	l.where = WhereCreate(condition)
	return l
}
