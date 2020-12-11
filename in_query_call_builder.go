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

func (i InQueryCallBuilder) Where(condition Condition) OngoingReadingWithWhere {
	i.conditionBuilder.Where(condition)
	i.defaultBuilder.currentOngoingCall = i
	return i.defaultBuilder
}

func (i InQueryCallBuilder) WherePattern(pattern RelationshipPattern) OngoingReadingWithWhere {
	return i.Where(RelationshipPatternConditionCreate(pattern))
}

func (i InQueryCallBuilder) ReturningByString(variables ...string) OngoingReadingAndReturn {
	return i.Returning(CreateSymbolicNameByString(variables...)...)
}

func (i InQueryCallBuilder) ReturningByNamed(variables ...Named) OngoingReadingAndReturn {
	return i.Returning(CreateSymbolicNameByNamed(variables...)...)
}

func (i InQueryCallBuilder) Returning(expression ...Expression) OngoingReadingAndReturn {
	i.defaultBuilder.currentOngoingCall = i
	return i.defaultBuilder.Returning(expression...)
}

func (i InQueryCallBuilder) ReturningDistinctByString(variables ...string) OngoingReadingAndReturn {
	return i.ReturningDistinct(CreateSymbolicNameByString(variables...)...)
}

func (i InQueryCallBuilder) ReturningDistinctByNamed(variables ...Named) OngoingReadingAndReturn {
	return i.ReturningDistinct(CreateSymbolicNameByNamed(variables...)...)
}

func (i InQueryCallBuilder) ReturningDistinct(expression ...Expression) OngoingReadingAndReturn {
	i.defaultBuilder.currentOngoingCall = i
	return i.defaultBuilder.ReturningDistinct(expression...)
}

func (i InQueryCallBuilder) Build() (Statement, error) {
	panic("implement me")
}

func (i InQueryCallBuilder) WithArgs(arguments ...Expression) OngoingInQueryCallWithArguments {
	i.arguments = arguments
	return i
}

func (i InQueryCallBuilder) YieldSymbolic(name ...SymbolicName) OngoingInQueryCallWithReturnFields {
	expressions := make([]Expression, len(name))
	for i := range name {
		expressions[i] = name[i]
	}
	i.yieldItems = yieldAllOf(expressions...)
	return i
}

func (i InQueryCallBuilder) YieldString(yieldedItems ...string) OngoingInQueryCallWithReturnFields {
	names := make([]SymbolicName, len(yieldedItems))
	for i := range yieldedItems {
		names[i] = SymbolicNameCreate(yieldedItems[i])
	}
	return i.YieldSymbolic(names...)
}

func (i InQueryCallBuilder) Yield(aliasedResultFields ...AliasedExpression) OngoingInQueryCallWithReturnFields {
	expressions := make([]Expression, len(aliasedResultFields))
	for i := range aliasedResultFields {
		expressions[i] = aliasedResultFields[i]
	}
	i.yieldItems = yieldAllOf(expressions...)
	return i
}

func (i InQueryCallBuilder) WithByString(variables ...string) OrderableOngoingReadingAndWithWithoutWhere {
	return i.With(CreateSymbolicNameByString(variables...)...)
}

func (i InQueryCallBuilder) WithByNamed(variables ...Named) OrderableOngoingReadingAndWithWithoutWhere {
	return i.With(CreateSymbolicNameByNamed(variables...)...)
}

func (i InQueryCallBuilder) With(expressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere {
	i.defaultBuilder.currentOngoingCall = i
	return i.defaultBuilder.With(expressions...)
}

func (i InQueryCallBuilder) WithDistinctByString(variables ...string) OrderableOngoingReadingAndWithWithoutWhere {
	return i.WithDistinct(CreateSymbolicNameByString(variables...)...)
}

func (i InQueryCallBuilder) WithDistinctByNamed(variables ...Named) OrderableOngoingReadingAndWithWithoutWhere {
	return i.WithDistinct(CreateSymbolicNameByNamed(variables...)...)
}

func (i InQueryCallBuilder) WithDistinct(expressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere {
	i.defaultBuilder.currentOngoingCall = i
	return i.defaultBuilder.WithDistinct(expressions...)
}

func (i InQueryCallBuilder) Call(statement Statement) OngoingReadingWithoutWhere {
	i.defaultBuilder.currentOngoingCall = i
	return i.defaultBuilder.Call(statement)
}

func (i InQueryCallBuilder) isNotNil() bool {
	return i.notNil
}
