package cypher

type DefaultOngoingUnwind struct {
	defaultBuilder     *DefaultStatementBuilder
	expressionToUnwind Expression
	err                error
}

func DefaultOngoingUnwindCreate(defaultBuilder *DefaultStatementBuilder, expressionToUnwind Expression) DefaultOngoingUnwind {
	return DefaultOngoingUnwind{
		defaultBuilder:     defaultBuilder,
		expressionToUnwind: expressionToUnwind,
	}
}

func DefaultOngoingUnwindError(err error) DefaultOngoingUnwind {
	return DefaultOngoingUnwind{err: err}
}

func (d DefaultOngoingUnwind) GetError() error {
	return d.err
}

func (d DefaultOngoingUnwind) As(variable string) OngoingReading {
	d.defaultBuilder.currentSinglePartElements = append(d.defaultBuilder.currentSinglePartElements, unwindCreate(d.expressionToUnwind, variable))
	return d.defaultBuilder
}
