package cypher_go_dsl

type DefaultStatementWithUpdateBuilder struct {
	defaultBuilder *DefaultStatementBuilder
	returnList     []Expression
	orderBuilder   OrderBuilder
	distinct       bool
	builder        UpdatingClauseBuilder
	notNil         bool
	err            error
}

func DefaultStatementWithUpdateBuilderCreate(defaultBuilder *DefaultStatementBuilder) DefaultStatementWithUpdateBuilder {
	return DefaultStatementWithUpdateBuilder{
		defaultBuilder: defaultBuilder,
		distinct:       false,
	}
}

func DefaultStatementWithUpdateBuilderCreate1(defaultBuilder *DefaultStatementBuilder, updateType UpdateType, patternOrExpression []Visitable) DefaultStatementWithUpdateBuilder {
	builder, err := getUpdatingClauseBuilder(updateType, patternOrExpression...)
	if err != nil {
		return DefaultStatementWithUpdateBuilder{
			err: err,
		}
	}
	return DefaultStatementWithUpdateBuilder{
		defaultBuilder: defaultBuilder,
		distinct:       false,
		builder:        builder,
	}
}

func DefaultStatementWithUpdateBuilderCreate2(defaultBuilder *DefaultStatementBuilder, updateType UpdateType, expressions ...Expression) DefaultStatementWithUpdateBuilder {
	visitables := make([]Visitable, len(expressions))
	for i := range expressions {
		visitables[i] = expressions[i]
	}
	builder, err := getUpdatingClauseBuilder(updateType, visitables...)
	if err != nil {
		return DefaultStatementWithUpdateBuilder{
			err: err,
		}
	}
	return DefaultStatementWithUpdateBuilder{
		defaultBuilder: defaultBuilder,
		distinct:       false,
		builder:        builder,
	}
}

func (d DefaultStatementWithUpdateBuilder) isNotNil() bool {
	return d.notNil
}

func (d DefaultStatementWithUpdateBuilder) and(expression Expression) TerminalOngoingOrderDefinition {
	d.orderBuilder.And(expression)
	return d
}

func (d DefaultStatementWithUpdateBuilder) descending() OngoingMatchAndReturnWithOrder {
	d.orderBuilder.Descending()
	return d
}

func (d DefaultStatementWithUpdateBuilder) ascending() OngoingMatchAndReturnWithOrder {
	d.orderBuilder.Ascending()
	return d
}

func (d DefaultStatementWithUpdateBuilder) orderBySortItem(sortItem ...SortItem) OngoingMatchAndReturnWithOrder {
	d.orderBuilder.OrderBySortItem(sortItem...)
	return d
}

func (d DefaultStatementWithUpdateBuilder) orderBy(expression Expression) TerminalOngoingOrderDefinition {
	d.orderBuilder.OrderByExpression(expression)
	return d
}

func (d DefaultStatementWithUpdateBuilder) skip(number int) TerminalExposesLimitAndBuildableStatement {
	d.orderBuilder.Skip(number)
	return d
}

func (d DefaultStatementWithUpdateBuilder) limit(number int) BuildableStatement {
	d.orderBuilder.Limit(number)
	return d
}

func (d DefaultStatementWithUpdateBuilder) build() Statement {
	d.defaultBuilder.addUpdatingClause(d.builder.build())
	var returning Return
	if len(d.returnList) > 0 {
		returnItems := ExpressionList{expressions: d.returnList}
		returning = ReturnCreate1(d.distinct, returnItems, d.orderBuilder.BuildOrder(), d.orderBuilder.skip,
			d.orderBuilder.limit)
	}
	return d.defaultBuilder.BuildImpl(false, returning)
}

func (d DefaultStatementWithUpdateBuilder) returningByString(variables ...string) OngoingReadingAndReturn {
	return d.returning(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementWithUpdateBuilder) returningByNamed(variables ...Named) OngoingReadingAndReturn {
	return d.returning(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementWithUpdateBuilder) returning(expression ...Expression) OngoingReadingAndReturn {
	d.returnList = append(d.returnList, expression...)
	return d
}

func (d DefaultStatementWithUpdateBuilder) returningDistinctByString(variables ...string) OngoingReadingAndReturn {
	return d.returningDistinct(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementWithUpdateBuilder) returningDistinctByNamed(variables ...Named) OngoingReadingAndReturn {
	return d.returningDistinct(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementWithUpdateBuilder) returningDistinct(expression ...Expression) OngoingReadingAndReturn {
	d.returning(expression...)
	d.distinct = true
	return d
}

func (d DefaultStatementWithUpdateBuilder) withByString(variables ...string) OrderableOngoingReadingAndWithWithoutWhere {
	return d.with(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementWithUpdateBuilder) withByNamed(variables ...Named) OrderableOngoingReadingAndWithWithoutWhere {
	return d.with(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementWithUpdateBuilder) with(expressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere {
	return d.withDefault(false, expressions...)
}

func (d DefaultStatementWithUpdateBuilder) withDefault(distinct bool, returnedExpressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere {
	d.defaultBuilder.addUpdatingClause(d.builder.build())
	return d.defaultBuilder.withDefault(distinct, returnedExpressions...)
}

func (d DefaultStatementWithUpdateBuilder) withDistinctByString(variables ...string) OrderableOngoingReadingAndWithWithoutWhere {
	return d.withDistinct(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementWithUpdateBuilder) withDistinctByNamed(variables ...Named) OrderableOngoingReadingAndWithWithoutWhere {
	return d.withDistinct(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementWithUpdateBuilder) withDistinct(expressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere {
	return d.withDefault(true, expressions...)
}

func (d DefaultStatementWithUpdateBuilder) deleteByString(variables ...string) OngoingUpdate {
	return d.delete(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementWithUpdateBuilder) deleteByNamed(variables ...Named) OngoingUpdate {
	return d.delete(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementWithUpdateBuilder) delete(expressions ...Expression) OngoingUpdate {
	return d.deleteDefault(false, expressions...)
}

func (d DefaultStatementWithUpdateBuilder) deleteDefault(nextDetach bool, deletedExpressions ...Expression) OngoingUpdate {
	d.defaultBuilder.addUpdatingClause(d.builder.build())
	var deleteType UpdateType
	if nextDetach {
		deleteType = UPDATE_TYPE_DETACH_DELETE
	} else {
		deleteType = UPDATE_TYPE_DELETE
	}
	visitables := make([]Visitable, len(deletedExpressions))
	for i := range deletedExpressions {
		visitables[i] = deletedExpressions[i]
	}
	return d.defaultBuilder.update(deleteType, visitables)
}

func (d DefaultStatementWithUpdateBuilder) detachDeleteByString(variables ...string) OngoingUpdate {
	return d.detachDelete(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementWithUpdateBuilder) detachDeleteByNamed(variables ...Named) OngoingUpdate {
	return d.detachDelete(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementWithUpdateBuilder) detachDelete(expressions ...Expression) OngoingUpdate {
	return d.deleteDefault(true, expressions...)
}

func (d DefaultStatementWithUpdateBuilder) merge(pattern ...PatternElement) OngoingUpdate {
	d.defaultBuilder.addUpdatingClause(d.builder.build())
	return d.defaultBuilder.merge(pattern...)
}

func (d DefaultStatementWithUpdateBuilder) set(expressions ...Expression) BuildableStatementAndOngoingMatchAndUpdate {
	d.defaultBuilder.addUpdatingClause(d.builder.build())
	return DefaultStatementWithUpdateBuilderCreate2(d.defaultBuilder, UPDATE_TYPE_SET, expressions...)
}

func (d DefaultStatementWithUpdateBuilder) setWithNamed(variable Named, expression Expression) BuildableStatementAndOngoingMatchAndUpdate {
	return d.set(variable.getSymbolicName(), expression)
}

func (d DefaultStatementWithUpdateBuilder) setByNode(node Node, labels ...string) BuildableStatementAndOngoingMatchAndUpdate {
	d.defaultBuilder.addUpdatingClause(d.builder.build())
	return DefaultStatementWithUpdateBuilderCreate2(d.defaultBuilder, UPDATE_TYPE_SET, set1(node, labels...))
}

func (d DefaultStatementWithUpdateBuilder) removeByNode(node Node, labels ...string) BuildableStatementAndOngoingMatchAndUpdate {
	d.defaultBuilder.addUpdatingClause(d.builder.build())
	return DefaultStatementWithUpdateBuilderCreate2(d.defaultBuilder, UPDATE_TYPE_REMOVE, set1(node, labels...))
}

func (d DefaultStatementWithUpdateBuilder) remove(properties ...Property) BuildableStatementAndOngoingMatchAndUpdate {
	expressions := make([]Expression, len(properties))
	for i := range properties {
		expressions[i] = properties[i]
	}
	d.defaultBuilder.addUpdatingClause(d.builder.build())
	return DefaultStatementWithUpdateBuilderCreate2(d.defaultBuilder, UPDATE_TYPE_REMOVE, expressions...)
}

func (d DefaultStatementWithUpdateBuilder) create(element ...PatternElement) OngoingUpdate {
	d.defaultBuilder.addUpdatingClause(d.builder.build())
	return d.defaultBuilder.create(element...)
}
