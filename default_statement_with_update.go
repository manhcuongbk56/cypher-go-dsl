package cypher

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
		notNil:         true,
	}
}

func DefaultStatementWithUpdateBuilderCreate2(defaultBuilder *DefaultStatementBuilder, updateType UpdateType, expressions ...Expression) DefaultStatementWithUpdateBuilder {
	if defaultBuilder != nil && defaultBuilder.err != nil {
		return DefaultStatementWithUpdateBuilder{
			err: defaultBuilder.err,
		}
	}
	visitables := make([]Visitable, len(expressions))
	for i := range expressions {
		if expressions[i] != nil && expressions[i].GetError() != nil {
			return DefaultStatementWithUpdateBuilder{
				err: expressions[i].GetError(),
			}
		}
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
		notNil:         true,
	}
}

func DefaultStatementWithUpdateBuilderError(err error) DefaultStatementWithUpdateBuilder {
	return DefaultStatementWithUpdateBuilder{err: err}
}

func (d DefaultStatementWithUpdateBuilder) isNotNil() bool {
	return d.notNil
}

func (d DefaultStatementWithUpdateBuilder) And(expression Expression) TerminalOngoingOrderDefinition {
	d.orderBuilder.And(expression)
	return d
}

func (d DefaultStatementWithUpdateBuilder) Descending() OngoingMatchAndReturnWithOrder {
	d.orderBuilder.Descending()
	return d
}

func (d DefaultStatementWithUpdateBuilder) Ascending() OngoingMatchAndReturnWithOrder {
	d.orderBuilder.Ascending()
	return d
}

func (d DefaultStatementWithUpdateBuilder) OrderBySortItem(sortItem ...SortItem) OngoingMatchAndReturnWithOrder {
	d.orderBuilder.OrderBySortItem(sortItem...)
	return d
}

func (d DefaultStatementWithUpdateBuilder) OrderBy(expression Expression) TerminalOngoingOrderDefinition {
	d.orderBuilder.OrderByExpression(expression)
	return d
}

func (d DefaultStatementWithUpdateBuilder) Skip(number int) TerminalExposesLimitAndBuildableStatement {
	d.orderBuilder.Skip(number)
	return d
}

func (d DefaultStatementWithUpdateBuilder) Limit(number int) BuildableStatement {
	d.orderBuilder.Limit(number)
	return d
}

func (d DefaultStatementWithUpdateBuilder) Build() (Statement, error) {
	if d.err != nil {
		return nil, d.err
	}
	d.defaultBuilder.addUpdatingClause(d.builder.build())
	var returning Return
	if len(d.returnList) > 0 {
		returnItems := ExpressionList{expressions: d.returnList}
		returning = ReturnCreate1(d.distinct, returnItems, d.orderBuilder.BuildOrder(), d.orderBuilder.skip,
			d.orderBuilder.limit)
	}
	return d.defaultBuilder.BuildImpl(false, returning), nil
}

func (d DefaultStatementWithUpdateBuilder) ReturningByString(variables ...string) OngoingReadingAndReturn {
	return d.Returning(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementWithUpdateBuilder) ReturningByNamed(variables ...Named) OngoingReadingAndReturn {
	return d.Returning(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementWithUpdateBuilder) Returning(expression ...Expression) OngoingReadingAndReturn {
	d.returnList = append(d.returnList, expression...)
	return d
}

func (d DefaultStatementWithUpdateBuilder) ReturningDistinctByString(variables ...string) OngoingReadingAndReturn {
	return d.ReturningDistinct(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementWithUpdateBuilder) ReturningDistinctByNamed(variables ...Named) OngoingReadingAndReturn {
	return d.ReturningDistinct(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementWithUpdateBuilder) ReturningDistinct(expression ...Expression) OngoingReadingAndReturn {
	d.Returning(expression...)
	d.distinct = true
	return d
}

func (d DefaultStatementWithUpdateBuilder) WithByString(variables ...string) OrderableOngoingReadingAndWithWithoutWhere {
	return d.With(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementWithUpdateBuilder) WithByNamed(variables ...Named) OrderableOngoingReadingAndWithWithoutWhere {
	return d.With(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementWithUpdateBuilder) With(expressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere {
	return d.WithDefault(false, expressions...)
}

func (d DefaultStatementWithUpdateBuilder) WithDefault(distinct bool, returnedExpressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere {
	d.defaultBuilder.addUpdatingClause(d.builder.build())
	return d.defaultBuilder.WithDefault(distinct, returnedExpressions...)
}

func (d DefaultStatementWithUpdateBuilder) WithDistinctByString(variables ...string) OrderableOngoingReadingAndWithWithoutWhere {
	return d.WithDistinct(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementWithUpdateBuilder) WithDistinctByNamed(variables ...Named) OrderableOngoingReadingAndWithWithoutWhere {
	return d.WithDistinct(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementWithUpdateBuilder) WithDistinct(expressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere {
	return d.WithDefault(true, expressions...)
}

func (d DefaultStatementWithUpdateBuilder) DeleteByString(variables ...string) OngoingUpdate {
	return d.Delete(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementWithUpdateBuilder) DeleteByNamed(variables ...Named) OngoingUpdate {
	return d.Delete(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementWithUpdateBuilder) Delete(expressions ...Expression) OngoingUpdate {
	return d.DeleteDefault(false, expressions...)
}

func (d DefaultStatementWithUpdateBuilder) DeleteDefault(nextDetach bool, deletedExpressions ...Expression) OngoingUpdate {
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
	return d.defaultBuilder.Update(deleteType, visitables)
}

func (d DefaultStatementWithUpdateBuilder) DetachDeleteByString(variables ...string) OngoingUpdate {
	return d.DetachDelete(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementWithUpdateBuilder) DetachDeleteByNamed(variables ...Named) OngoingUpdate {
	return d.DetachDelete(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementWithUpdateBuilder) DetachDelete(expressions ...Expression) OngoingUpdate {
	return d.DeleteDefault(true, expressions...)
}

func (d DefaultStatementWithUpdateBuilder) Merge(pattern ...PatternElement) OngoingUpdateAndExposesSet {
	d.defaultBuilder.addUpdatingClause(d.builder.build())
	return d.defaultBuilder.Merge(pattern...)
}

func (d DefaultStatementWithUpdateBuilder) Set(expressions ...Expression) BuildableStatementAndOngoingMatchAndUpdate {
	d.defaultBuilder.addUpdatingClause(d.builder.build())
	return DefaultStatementWithUpdateBuilderCreate2(d.defaultBuilder, UPDATE_TYPE_SET, expressions...)
}

func (d DefaultStatementWithUpdateBuilder) SetWithNamed(variable Named, expression Expression) BuildableStatementAndOngoingMatchAndUpdate {
	return d.Set(variable.getSymbolicName(), expression)
}

func (d DefaultStatementWithUpdateBuilder) SetByNode(node Node, labels ...string) BuildableStatementAndOngoingMatchAndUpdate {
	d.defaultBuilder.addUpdatingClause(d.builder.build())
	return DefaultStatementWithUpdateBuilderCreate2(d.defaultBuilder, UPDATE_TYPE_SET, OperationSetLabel(node, labels...))
}

func (d DefaultStatementWithUpdateBuilder) RemoveByNode(node Node, labels ...string) BuildableStatementAndOngoingMatchAndUpdate {
	if d.err != nil {
		return d
	}
	d.defaultBuilder.addUpdatingClause(d.builder.build())
	return DefaultStatementWithUpdateBuilderCreate2(d.defaultBuilder, UPDATE_TYPE_REMOVE, OperationSetLabel(node, labels...))
}

func (d DefaultStatementWithUpdateBuilder) Remove(properties ...Property) BuildableStatementAndOngoingMatchAndUpdate {
	expressions := make([]Expression, len(properties))
	for i := range properties {
		expressions[i] = properties[i]
	}
	d.defaultBuilder.addUpdatingClause(d.builder.build())
	return DefaultStatementWithUpdateBuilderCreate2(d.defaultBuilder, UPDATE_TYPE_REMOVE, expressions...)
}

func (d DefaultStatementWithUpdateBuilder) Create(element ...PatternElement) OngoingUpdateAndExposesSet {
	d.defaultBuilder.addUpdatingClause(d.builder.build())
	return d.defaultBuilder.Create(element...)
}
