package cypher_go_dsl

type InQueryCallBuilder struct {
	defaultBuilder   *DefaultStatementBuilder
	procedureName    ProcedureName
	arguments        []Expression
	yieldItems       YieldItems
	conditionBuilder ConditionBuilder
	notNil           bool
}

func InQueryCallBuilderCreate(defaultBuilder *DefaultStatementBuilder, procedureName ProcedureName) InQueryCallBuilder {
	return InQueryCallBuilder{
		defaultBuilder:   defaultBuilder,
		procedureName:    procedureName,
		conditionBuilder: ConditionBuilderCreate(),
		notNil:           true,
	}
}

func (i InQueryCallBuilder) where(condition Condition) OngoingReadingWithWhere {
	i.conditionBuilder.Where(condition)
	i.defaultBuilder.currentOngoingCall = i
	return i.defaultBuilder
}

func (i InQueryCallBuilder) wherePattern(pattern RelationshipPattern) OngoingReadingWithWhere {
	return i.where(RelationshipPatternConditionCreate(pattern))
}

func (i InQueryCallBuilder) returningByString(variables ...string) OngoingReadingAndReturn {
	return i.returning(CreateSymbolicNameByString(variables...)...)
}

func (i InQueryCallBuilder) returningByNamed(variables ...Named) OngoingReadingAndReturn {
	return i.returning(CreateSymbolicNameByNamed(variables...)...)
}

func (i InQueryCallBuilder) returning(expression ...Expression) OngoingReadingAndReturn {
	i.defaultBuilder.currentOngoingCall = i
	return i.defaultBuilder.returning(expression...)
}

func (i InQueryCallBuilder) returningDistinctByString(variables ...string) OngoingReadingAndReturn {
	return i.returningDistinct(CreateSymbolicNameByString(variables...)...)
}

func (i InQueryCallBuilder) returningDistinctByNamed(variables ...Named) OngoingReadingAndReturn {
	return i.returningDistinct(CreateSymbolicNameByNamed(variables...)...)
}

func (i InQueryCallBuilder) returningDistinct(expression ...Expression) OngoingReadingAndReturn {
	i.defaultBuilder.currentOngoingCall = i
	return i.defaultBuilder.returningDistinct(expression...)
}

func (i InQueryCallBuilder) build() (Statement, error) {
	panic("implement me")
}

func (i InQueryCallBuilder) withArgs(arguments ...Expression) OngoingInQueryCallWithArguments {
	i.arguments = arguments
	return i
}

func (i InQueryCallBuilder) yieldSymbolic(name ...SymbolicName) OngoingInQueryCallWithReturnFields {
	expressions := make([]Expression, len(name))
	for i := range name {
		expressions[i] = name[i]
	}
	i.yieldItems = yieldAllOf(expressions...)
	return i
}

func (i InQueryCallBuilder) yieldString(yieldedItems ...string) OngoingInQueryCallWithReturnFields {
	names := make([]SymbolicName, len(yieldedItems))
	for i := range yieldedItems {
		names[i] = SymbolicNameCreate(yieldedItems[i])
	}
	return i.yieldSymbolic(names...)
}

func (i InQueryCallBuilder) yield(aliasedResultFields ...AliasedExpression) OngoingInQueryCallWithReturnFields {
	expressions := make([]Expression, len(aliasedResultFields))
	for i := range aliasedResultFields {
		expressions[i] = aliasedResultFields[i]
	}
	i.yieldItems = yieldAllOf(expressions...)
	return i
}

func (i InQueryCallBuilder) withByString(variables ...string) OrderableOngoingReadingAndWithWithoutWhere {
	return i.with(CreateSymbolicNameByString(variables...)...)
}

func (i InQueryCallBuilder) withByNamed(variables ...Named) OrderableOngoingReadingAndWithWithoutWhere {
	return i.with(CreateSymbolicNameByNamed(variables...)...)
}

func (i InQueryCallBuilder) with(expressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere {
	i.defaultBuilder.currentOngoingCall = i
	return i.defaultBuilder.with(expressions...)
}

func (i InQueryCallBuilder) withDistinctByString(variables ...string) OrderableOngoingReadingAndWithWithoutWhere {
	return i.withDistinct(CreateSymbolicNameByString(variables...)...)
}

func (i InQueryCallBuilder) withDistinctByNamed(variables ...Named) OrderableOngoingReadingAndWithWithoutWhere {
	return i.withDistinct(CreateSymbolicNameByNamed(variables...)...)
}

func (i InQueryCallBuilder) withDistinct(expressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere {
	i.defaultBuilder.currentOngoingCall = i
	return i.defaultBuilder.withDistinct(expressions...)
}

func (i InQueryCallBuilder) call(statement Statement) OngoingReadingWithoutWhere {
	i.defaultBuilder.currentOngoingCall = i
	return i.defaultBuilder.call(statement)
}

func (i InQueryCallBuilder) isNotNil() bool {
	return i.notNil
}
