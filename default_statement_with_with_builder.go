package cypher_go_dsl

type DefaultStatementWithWithBuilder struct {
	defaultBuilder   *DefaultStatementBuilder
	conditionBuilder ConditionBuilder
	returnList       []Expression
	orderBuilder     OrderBuilder
	distinct         bool
	err              error
}

func DefaultStatementWithWithBuilderCreate(defaultBuilder *DefaultStatementBuilder, distinct bool) DefaultStatementWithWithBuilder {
	return DefaultStatementWithWithBuilder{
		defaultBuilder:   defaultBuilder,
		distinct:         distinct,
		conditionBuilder: ConditionBuilderCreate(),
		returnList:       make([]Expression, 0),
		orderBuilder:     OrderBuilderCreate(),
	}
}

func DefaultStatementWithWithBuilderError(err error) DefaultStatementWithWithBuilder {
	return DefaultStatementWithWithBuilder{err: err}
}

func (d *DefaultStatementWithWithBuilder) buildWith() With {
	if len(d.returnList) == 0 {
		return With{}
	}
	returnItems := ExpressionListCreate(d.returnList)
	condition := d.conditionBuilder.buildCondition()
	var where Where
	if condition == nil || !condition.isNotNil() {
		where = Where{}

	} else {
		where = WhereCreate(condition)
	}
	returnedWith := WithCreate(d.distinct, returnItems, d.orderBuilder.BuildOrder(), d.orderBuilder.skip,
		d.orderBuilder.limit, where)
	d.returnList = nil
	d.orderBuilder.reset()
	return returnedWith
}

func (d *DefaultStatementWithWithBuilder) addExpression(expressions ...Expression) {
	d.returnList = append(d.returnList, expressions...)
}

func (d DefaultStatementWithWithBuilder) Skip(number int) ExposesLimitAndOngoingReadingAndWith {
	d.orderBuilder.Skip(number)
	return d
}

func (d DefaultStatementWithWithBuilder) Limit(number int) OngoingReadingAndWith {
	d.orderBuilder.Limit(number)
	return d
}

func (d DefaultStatementWithWithBuilder) Descending() OngoingReadingAndWithWithWhereAndOrder {
	d.orderBuilder.Descending()
	return d
}

func (d DefaultStatementWithWithBuilder) Ascending() OngoingReadingAndWithWithWhereAndOrder {
	d.orderBuilder.Ascending()
	return d
}

func (d DefaultStatementWithWithBuilder) OrderBySortItem(sortItem ...SortItem) OrderableOngoingReadingAndWithWithWhere {
	d.orderBuilder.OrderBySortItem(sortItem...)
	return d
}

func (d DefaultStatementWithWithBuilder) OrderByExpression(expression Expression) OngoingOrderDefinition {
	d.orderBuilder.OrderByExpression(expression)
	return d
}

func (d DefaultStatementWithWithBuilder) ReturningByString(variables ...string) OngoingReadingAndReturn {
	return d.Returning(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementWithWithBuilder) ReturningByNamed(variables ...Named) OngoingReadingAndReturn {
	return d.Returning(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementWithWithBuilder) Returning(expression ...Expression) OngoingReadingAndReturn {
	return d.defaultBuilder.AddWith(d.buildWith()).
		Returning(expression...)
}

func (d DefaultStatementWithWithBuilder) ReturningDistinctByString(variables ...string) OngoingReadingAndReturn {
	return d.ReturningDistinct(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementWithWithBuilder) ReturningDistinctByNamed(variables ...Named) OngoingReadingAndReturn {
	return d.ReturningDistinct(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementWithWithBuilder) ReturningDistinct(expression ...Expression) OngoingReadingAndReturn {
	return d.defaultBuilder.AddWith(d.buildWith()).
		ReturningDistinct(expression...)
}

func (d DefaultStatementWithWithBuilder) WithByString(variables ...string) OrderableOngoingReadingAndWithWithoutWhere {
	return d.With(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementWithWithBuilder) WithByNamed(variables ...Named) OrderableOngoingReadingAndWithWithoutWhere {
	return d.With(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementWithWithBuilder) With(expressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere {
	return d.defaultBuilder.AddWith(d.buildWith()).
		With(expressions...)
}

func (d DefaultStatementWithWithBuilder) WithDistinctByString(variables ...string) OrderableOngoingReadingAndWithWithoutWhere {
	return d.WithDistinct(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementWithWithBuilder) WithDistinctByNamed(variables ...Named) OrderableOngoingReadingAndWithWithoutWhere {
	return d.WithDistinct(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementWithWithBuilder) WithDistinct(expressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere {
	return d.defaultBuilder.AddWith(d.buildWith()).
		WithDistinct(expressions...)
}

func (d DefaultStatementWithWithBuilder) DeleteByString(variables ...string) OngoingUpdate {
	return d.Delete(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementWithWithBuilder) DeleteByNamed(variables ...Named) OngoingUpdate {
	return d.Delete(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementWithWithBuilder) Delete(expressions ...Expression) OngoingUpdate {
	return d.defaultBuilder.AddWith(d.buildWith()).
		Delete(expressions...)
}

func (d DefaultStatementWithWithBuilder) DetachDeleteByString(variables ...string) OngoingUpdate {
	return d.DetachDelete(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementWithWithBuilder) DetachDeleteByNamed(variables ...Named) OngoingUpdate {
	return d.DetachDelete(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementWithWithBuilder) DetachDelete(expressions ...Expression) OngoingUpdate {
	return d.defaultBuilder.AddWith(d.buildWith()).
		Delete(expressions...)
}

func (d DefaultStatementWithWithBuilder) Merge(pattern ...PatternElement) OngoingUpdate {
	return d.defaultBuilder.AddWith(d.buildWith()).
		Merge(pattern...)
}

func (d DefaultStatementWithWithBuilder) Set(expressions ...Expression) BuildableStatementAndOngoingMatchAndUpdate {
	return d.defaultBuilder.AddWith(d.buildWith()).
		Set(expressions...)
}

func (d DefaultStatementWithWithBuilder) SetWithNamed(variable Named, expression Expression) BuildableStatementAndOngoingMatchAndUpdate {
	return d.Set(variable.getRequiredSymbolicName(), expression)
}

func (d DefaultStatementWithWithBuilder) SetByNode(node Node, labels ...string) BuildableStatementAndOngoingMatchAndUpdate {
	return d.defaultBuilder.AddWith(d.buildWith()).
		SetByNode(node, labels...)
}

func (d DefaultStatementWithWithBuilder) RemoveByNode(node Node, labels ...string) BuildableStatementAndOngoingMatchAndUpdate {
	return d.defaultBuilder.AddWith(d.buildWith()).
		RemoveByNode(node, labels...)
}

func (d DefaultStatementWithWithBuilder) Remove(properties ...Property) BuildableStatementAndOngoingMatchAndUpdate {
	return d.defaultBuilder.AddWith(d.buildWith()).
		Remove(properties...)
}

func (d DefaultStatementWithWithBuilder) Unwinds(expression ...Expression) OngoingUnwind {
	return d.Unwind(ListOf(expression...))
}

func (d DefaultStatementWithWithBuilder) UnwindByString(variable string) OngoingUnwind {
	return d.Unwind(Name(variable))
}

func (d DefaultStatementWithWithBuilder) Unwind(expression Expression) OngoingUnwind {
	return d.defaultBuilder.AddWith(d.buildWith()).
		Unwind(expression)
}

func (d DefaultStatementWithWithBuilder) Create(element ...PatternElement) OngoingUpdate {
	return d.defaultBuilder.AddWith(d.buildWith()).
		Create(element...)
}

func (d DefaultStatementWithWithBuilder) Call(statement Statement) OngoingReadingWithoutWhere {
	return d.defaultBuilder.AddWith(d.buildWith()).
		Call(statement)
}

func (d DefaultStatementWithWithBuilder) Call1(namespaceAndProcedure ...string) OngoingInQueryCallWithoutArguments {
	return d.defaultBuilder.AddWith(d.buildWith()).
		Call1(namespaceAndProcedure...)
}

func (d DefaultStatementWithWithBuilder) Match(pattern ...PatternElement) OngoingReadingWithoutWhere {
	return d.MatchDefault(false, pattern...)
}

func (d DefaultStatementWithWithBuilder) OptionalMatch(pattern ...PatternElement) OngoingReadingWithoutWhere {
	return d.MatchDefault(true, pattern...)
}

func (d DefaultStatementWithWithBuilder) MatchDefault(optional bool, pattern ...PatternElement) OngoingReadingWithoutWhere {
	return d.defaultBuilder.AddWith(d.buildWith()).
		MatchDefault(optional, pattern...)
}

func (d DefaultStatementWithWithBuilder) Where(condition Condition) OrderableOngoingReadingAndWithWithWhere {
	d.conditionBuilder.Where(condition)
	return d
}

func (d DefaultStatementWithWithBuilder) WherePattern(pattern RelationshipPattern) OrderableOngoingReadingAndWithWithWhere {
	return d.Where(RelationshipPatternConditionCreate(pattern))
}

func (d DefaultStatementWithWithBuilder) And(condition Condition) OrderableOngoingReadingAndWithWithWhere {
	d.conditionBuilder.And(condition)
	return d
}

func (d DefaultStatementWithWithBuilder) AndPattern(pattern RelationshipPattern) OrderableOngoingReadingAndWithWithWhere {
	return d.And(RelationshipPatternConditionCreate(pattern))
}

func (d DefaultStatementWithWithBuilder) Or(condition Condition) OrderableOngoingReadingAndWithWithWhere {
	d.conditionBuilder.Or(condition)
	return d
}

func (d DefaultStatementWithWithBuilder) OrPattern(pattern RelationshipPattern) OrderableOngoingReadingAndWithWithWhere {
	return d.Or(RelationshipPatternConditionCreate(pattern))
}

func (d DefaultStatementWithWithBuilder) And1(expression Expression) OngoingOrderDefinition {
	d.orderBuilder.And(expression)
	return d
}
