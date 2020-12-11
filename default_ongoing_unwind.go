package cypher_go_dsl

type DefaultOngoingUnwind struct {
	defaultBuilder     *DefaultStatementBuilder
	expressionToUnwind Expression
}

func DefaultOngoingUnwindCreate(defaultBuilder *DefaultStatementBuilder, expressionToUnwind Expression) DefaultOngoingUnwind {
	return DefaultOngoingUnwind{
		defaultBuilder:     defaultBuilder,
		expressionToUnwind: expressionToUnwind,
	}
}

func (d DefaultOngoingUnwind) As(variable string) OngoingReading {
	d.defaultBuilder.currentSinglePartElements = append(d.defaultBuilder.currentSinglePartElements, UnwindCreate(d.expressionToUnwind, variable))
	return d.defaultBuilder
}
