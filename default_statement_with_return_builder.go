package cypher_go_dsl

type DefaultStatementWithReturnBuilder struct {
	defaultStatementBuilder DefaultStatementBuilder
	distinct                bool
	returnList              []Expression
	orderBuilder            OrderBuilder
}

func (builder *DefaultStatementWithReturnBuilder) AddExpression(expression ...Expression) {
	builder.returnList = append(builder.returnList, expression...)
}

func (builder DefaultStatementWithReturnBuilder) Build() Statement {
	var returning Return
	if len(builder.returnList) > 0 {
		returnItems := ExpressionList{expressions: builder.returnList}
		returning = ReturnByMultiVariable(builder.distinct, returnItems, builder.orderBuilder.BuildOrder(), builder.orderBuilder.skip,
			builder.orderBuilder.limit)
	}
	return builder.defaultStatementBuilder.BuildImpl(returning)
}

func (builder DefaultStatementWithReturnBuilder) and(expression Expression) TerminalOngoingOrderDefinition {
	panic("implement me")
}

func (builder DefaultStatementWithReturnBuilder) descending() OngoingMatchAndReturnWithOrder {
	panic("implement me")
}

func (builder DefaultStatementWithReturnBuilder) ascending() OngoingMatchAndReturnWithOrder {
	panic("implement me")
}

func (builder DefaultStatementWithReturnBuilder) orderBySortItem(sortItem ...SortItem) OngoingMatchAndReturnWithOrder {
	panic("implement me")
}

func (builder DefaultStatementWithReturnBuilder) orderBy(expression Expression) {
	panic("implement me")
}

func (builder DefaultStatementWithReturnBuilder) skip(number int) TerminalExposesLimitAndBuildableStatement {
	panic("implement me")
}

func (builder DefaultStatementWithReturnBuilder) limit(number int) BuildableStatement {
	panic("implement me")
}
