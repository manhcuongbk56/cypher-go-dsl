package cypher

type DefaultStatementWithReturnBuilder struct {
	defaultBuilder *DefaultStatementBuilder
	distinct       bool
	returnList     []Expression
	orderBuilder   OrderBuilder
	err            error
}

func DefaultStatementWithReturnBuilderError(err error) DefaultStatementWithReturnBuilder {
	return DefaultStatementWithReturnBuilder{err: err}
}

func (d *DefaultStatementWithReturnBuilder) AddExpression(expression ...Expression) {
	d.returnList = append(d.returnList, expression...)
}

func (d DefaultStatementWithReturnBuilder) Build() (Statement, error) {
	if d.err != nil {
		return nil, d.err
	}
	var returning Return
	if len(d.returnList) > 0 {
		returnItems := ExpressionList{expressions: d.returnList}
		returning = ReturnCreate1(d.distinct, returnItems, d.orderBuilder.BuildOrder(), d.orderBuilder.skip,
			d.orderBuilder.limit)
	}
	return d.defaultBuilder.BuildImpl(false, returning), nil
}

func (d DefaultStatementWithReturnBuilder) And(expression Expression) TerminalOngoingOrderDefinition {
	d.orderBuilder.And(expression)
	return d
}

func (d DefaultStatementWithReturnBuilder) Descending() OngoingMatchAndReturnWithOrder {
	d.orderBuilder.Descending()
	return d
}

func (d DefaultStatementWithReturnBuilder) Ascending() OngoingMatchAndReturnWithOrder {
	d.orderBuilder.Ascending()
	return d
}

func (d DefaultStatementWithReturnBuilder) OrderBySortItem(sortItem ...SortItem) OngoingMatchAndReturnWithOrder {
	d.orderBuilder.OrderBySortItem(sortItem...)
	return d
}

func (d DefaultStatementWithReturnBuilder) OrderBy(expression Expression) TerminalOngoingOrderDefinition {
	d.orderBuilder.OrderByExpression(expression)
	return d
}

func (d DefaultStatementWithReturnBuilder) Skip(number int) TerminalExposesLimitAndBuildableStatement {
	d.orderBuilder.Skip(number)
	return d
}

func (d DefaultStatementWithReturnBuilder) Limit(number int) BuildableStatement {
	d.orderBuilder.Limit(number)
	return d
}
