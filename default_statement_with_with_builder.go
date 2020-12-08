package cypher_go_dsl

type DefaultStatementWithWithBuilder struct {
	defaultBuilder   *DefaultStatementBuilder
	conditionBuilder ConditionBuilder
	returnList       []Expression
	orderBuilder     OrderBuilder
	distinct         bool
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

func (d *DefaultStatementWithWithBuilder) buildWith() With {
	if len(d.returnList) == 0 {
		return With{}
	}
	returnItems := ExpressionListCreate(d.returnList)
	condition := d.conditionBuilder.buildCondition()
	var where Where
	if condition.isNotNil() {
		where = WhereCreate(condition)
	} else {
		where = Where{}
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

func (d DefaultStatementWithWithBuilder) skip(number int) ExposesLimitAndOngoingReadingAndWith {
	d.orderBuilder.Skip(number)
	return d
}

func (d DefaultStatementWithWithBuilder) limit(number int) OngoingReadingAndWith {
	d.orderBuilder.Limit(number)
	return d
}

func (d DefaultStatementWithWithBuilder) descending() OngoingReadingAndWithWithWhereAndOrder {
	d.orderBuilder.Descending()
	return d
}

func (d DefaultStatementWithWithBuilder) ascending() OngoingReadingAndWithWithWhereAndOrder {
	d.orderBuilder.Ascending()
	return d
}

func (d DefaultStatementWithWithBuilder) orderBySortItem(sortItem ...SortItem) OrderableOngoingReadingAndWithWithWhere {
	d.orderBuilder.OrderBySortItem(sortItem...)
	return d
}

func (d DefaultStatementWithWithBuilder) orderByExpression(expression Expression) OngoingOrderDefinition {
	d.orderBuilder.OrderByExpression(expression)
	return d
}

func (d DefaultStatementWithWithBuilder) returningByString(variables ...string) OngoingReadingAndReturn {
	return d.returning(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementWithWithBuilder) returningByNamed(variables ...Named) OngoingReadingAndReturn {
	return d.returning(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementWithWithBuilder) returning(expression ...Expression) OngoingReadingAndReturn {
	return d.defaultBuilder.addWith(d.buildWith()).
		returning(expression...)
}

func (d DefaultStatementWithWithBuilder) returningDistinctByString(variables ...string) OngoingReadingAndReturn {
	return d.returningDistinct(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementWithWithBuilder) returningDistinctByNamed(variables ...Named) OngoingReadingAndReturn {
	return d.returningDistinct(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementWithWithBuilder) returningDistinct(expression ...Expression) OngoingReadingAndReturn {
	return d.defaultBuilder.addWith(d.buildWith()).
		returningDistinct(expression...)
}

func (d DefaultStatementWithWithBuilder) withByString(variables ...string) OrderableOngoingReadingAndWithWithoutWhere {
	return d.with(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementWithWithBuilder) withByNamed(variables ...Named) OrderableOngoingReadingAndWithWithoutWhere {
	return d.with(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementWithWithBuilder) with(expressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere {
	return d.defaultBuilder.addWith(d.buildWith()).
		with(expressions...)
}

func (d DefaultStatementWithWithBuilder) withDistinctByString(variables ...string) OrderableOngoingReadingAndWithWithoutWhere {
	return d.withDistinct(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementWithWithBuilder) withDistinctByNamed(variables ...Named) OrderableOngoingReadingAndWithWithoutWhere {
	return d.withDistinct(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementWithWithBuilder) withDistinct(expressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere {
	return d.defaultBuilder.addWith(d.buildWith()).
		withDistinct(expressions...)
}

func (d DefaultStatementWithWithBuilder) deleteByString(variables ...string) OngoingUpdate {
	return d.delete(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementWithWithBuilder) deleteByNamed(variables ...Named) OngoingUpdate {
	return d.delete(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementWithWithBuilder) delete(expressions ...Expression) OngoingUpdate {
	return d.defaultBuilder.addWith(d.buildWith()).
		delete(expressions...)
}

func (d DefaultStatementWithWithBuilder) detachDeleteByString(variables ...string) OngoingUpdate {
	return d.detachDelete(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementWithWithBuilder) detachDeleteByNamed(variables ...Named) OngoingUpdate {
	return d.detachDelete(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementWithWithBuilder) detachDelete(expressions ...Expression) OngoingUpdate {
	return d.defaultBuilder.addWith(d.buildWith()).
		delete(expressions...)
}

func (d DefaultStatementWithWithBuilder) merge(pattern ...PatternElement) OngoingUpdate {
	return d.defaultBuilder.addWith(d.buildWith()).
		merge(pattern...)
}

func (d DefaultStatementWithWithBuilder) set(expressions ...Expression) BuildableStatementAndOngoingMatchAndUpdate {
	return d.defaultBuilder.addWith(d.buildWith()).
		set(expressions...)
}

func (d DefaultStatementWithWithBuilder) setWithNamed(variable Named, expression Expression) BuildableStatementAndOngoingMatchAndUpdate {
	return d.set(variable.getSymbolicName(), expression)
}

func (d DefaultStatementWithWithBuilder) setByNode(node Node, labels ...string) BuildableStatementAndOngoingMatchAndUpdate {
	return d.defaultBuilder.addWith(d.buildWith()).
		setByNode(node, labels...)
}

func (d DefaultStatementWithWithBuilder) removeByNode(node Node, labels ...string) BuildableStatementAndOngoingMatchAndUpdate {
	return d.defaultBuilder.addWith(d.buildWith()).
		removeByNode(node, labels...)
}

func (d DefaultStatementWithWithBuilder) remove(properties ...Property) BuildableStatementAndOngoingMatchAndUpdate {
	return d.defaultBuilder.addWith(d.buildWith()).
		remove(properties...)
}

func (d DefaultStatementWithWithBuilder) unwinds(expression ...Expression) OngoingUnwind {
	return d.unwind(ListOf(expression...))
}

func (d DefaultStatementWithWithBuilder) unwindByString(variable string) OngoingUnwind {
	return d.unwind(Name(variable))
}

func (d DefaultStatementWithWithBuilder) unwind(expression Expression) OngoingUnwind {
	return d.defaultBuilder.addWith(d.buildWith()).
		unwind(expression)
}

func (d DefaultStatementWithWithBuilder) create(element ...PatternElement) OngoingUpdate {
	return d.defaultBuilder.addWith(d.buildWith()).
		create(element...)
}

func (d DefaultStatementWithWithBuilder) call(statement Statement) OngoingReadingWithoutWhere {
	return d.defaultBuilder.addWith(d.buildWith()).
		call(statement)
}

func (d DefaultStatementWithWithBuilder) call1(namespaceAndProcedure ...string) OngoingInQueryCallWithoutArguments {
	return d.defaultBuilder.addWith(d.buildWith()).
		call1(namespaceAndProcedure...)
}

func (d DefaultStatementWithWithBuilder) match(pattern ...PatternElement) OngoingReadingWithoutWhere {
	return d.MatchDefault(false, pattern...)
}

func (d DefaultStatementWithWithBuilder) optionalMatch(pattern ...PatternElement) OngoingReadingWithoutWhere {
	return d.MatchDefault(true, pattern...)
}

func (d DefaultStatementWithWithBuilder) MatchDefault(optional bool, pattern ...PatternElement) OngoingReadingWithoutWhere {
	return d.defaultBuilder.addWith(d.buildWith()).
		MatchDefault(optional, pattern...)
}

func (d DefaultStatementWithWithBuilder) where(condition Condition) OrderableOngoingReadingAndWithWithWhere {
	d.conditionBuilder.Where(condition)
	return d
}

func (d DefaultStatementWithWithBuilder) wherePattern(pattern RelationshipPattern) OrderableOngoingReadingAndWithWithWhere {
	return d.where(RelationshipPatternConditionCreate(pattern))
}

func (d DefaultStatementWithWithBuilder) and(condition Condition) OrderableOngoingReadingAndWithWithWhere {
	d.conditionBuilder.And(condition)
	return d
}

func (d DefaultStatementWithWithBuilder) andPattern(pattern RelationshipPattern) OrderableOngoingReadingAndWithWithWhere {
	return d.and(RelationshipPatternConditionCreate(pattern))
}

func (d DefaultStatementWithWithBuilder) or(condition Condition) OrderableOngoingReadingAndWithWithWhere {
	d.conditionBuilder.Or(condition)
	return d
}

func (d DefaultStatementWithWithBuilder) orPattern(pattern RelationshipPattern) OrderableOngoingReadingAndWithWithWhere {
	return d.or(RelationshipPatternConditionCreate(pattern))
}

func (d DefaultStatementWithWithBuilder) and1(expression Expression) OngoingOrderDefinition {
	d.orderBuilder.And(expression)
	return d
}
