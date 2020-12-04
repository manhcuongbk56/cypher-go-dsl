package cypher_go_dsl

type DefaultStatementWithReturnBuilder struct {
	defaultBuilder DefaultStatementBuilder
	distinct       bool
	returnList     []Expression
	orderBuilder   OrderBuilder
}

func (d *DefaultStatementWithReturnBuilder) AddExpression(expression ...Expression) {
	d.returnList = append(d.returnList, expression...)
}

func (d DefaultStatementWithReturnBuilder) Build() Statement {
	var returning Return
	if len(d.returnList) > 0 {
		returnItems := ExpressionList{expressions: d.returnList}
		returning = ReturnCreate1(d.distinct, returnItems, d.orderBuilder.BuildOrder(), d.orderBuilder.skip,
			d.orderBuilder.limit)
	}
	return d.defaultBuilder.BuildImpl(false, returning)
}

func (d DefaultStatementWithReturnBuilder) and(expression Expression) TerminalOngoingOrderDefinition {
	d.orderBuilder.And(expression)
	return d
}

func (d DefaultStatementWithReturnBuilder) descending() OngoingMatchAndReturnWithOrder {
	d.orderBuilder.Descending()
	return d
}

func (d DefaultStatementWithReturnBuilder) ascending() OngoingMatchAndReturnWithOrder {
	d.orderBuilder.Ascending()
	return d
}

func (d DefaultStatementWithReturnBuilder) orderBySortItem(sortItem ...SortItem) OngoingMatchAndReturnWithOrder {
	d.orderBuilder.OrderBySortItem(sortItem...)
	return d
}

func (d DefaultStatementWithReturnBuilder) orderBy(expression Expression) TerminalOngoingOrderDefinition {
	d.orderBuilder.OrderByExpression(expression)
	return d
}

func (d DefaultStatementWithReturnBuilder) skip(number int) TerminalExposesLimitAndBuildableStatement {
	d.orderBuilder.Skip(number)
	return d
}

func (d DefaultStatementWithReturnBuilder) limit(number int) BuildableStatement {
	d.orderBuilder.Limit(number)
	return d
}
