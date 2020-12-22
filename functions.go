package cypher

import (
	"errors"
)

func FunctionIdByNode(node Node) FunctionInvocation {
	if node.getError() != nil {
		return FunctionInvocationError(node.err)
	}
	if !node.isNotNil() {
		return FunctionInvocationError(errors.New("node is required"))
	}
	return FunctionInvocationCreate(ID, node.getSymbolicName())
}

func FunctionIdByRelationship(relationship Relationship) FunctionInvocation {
	if relationship.getError() != nil {
		return FunctionInvocationError(relationship.err)
	}
	if !relationship.isNotNil() {
		return FunctionInvocationError(errors.New("relationship is required"))
	}
	return FunctionInvocationCreate(ID, relationship.getSymbolicName())
}

func FunctionLabels(node Node) FunctionInvocation {
	if node.getError() != nil {
		return FunctionInvocationError(node.err)
	}
	if !node.isNotNil() {
		return FunctionInvocationError(errors.New("node is required"))
	}
	return FunctionInvocationCreate(LABELS, node.getSymbolicName())
}

func FunctionType(relationship Relationship) FunctionInvocation {
	if relationship.getError() != nil {
		return FunctionInvocationError(relationship.err)
	}
	if !relationship.isNotNil() {
		return FunctionInvocationError(errors.New("relationship is required"))
	}
	return FunctionInvocationCreate(TYPE, relationship.getSymbolicName())
}

func FunctionCount(node Node) FunctionInvocation {
	return FunctionInvocationCreate(COUNT, node.getSymbolicName())
}

func FunctionCountByExpression(expression Expression) FunctionInvocation {
	return FunctionInvocationCreate(COUNT, expression)
}

func FunctionCountDistinct(node Node) FunctionInvocation {
	return FunctionInvocationCreateDistinct(COUNT, node.getSymbolicName())
}

func FunctionCountDistinctByExpression(expression Expression) FunctionInvocation {
	return FunctionInvocationCreateDistinct(COUNT, expression)
}

func FunctionProperties(node Node) FunctionInvocation {
	if node.getError() != nil {
		return FunctionInvocationError(node.err)
	}
	if !node.isNotNil() {
		return FunctionInvocationError(errors.New("node is required"))
	}
	return FunctionInvocationCreate(PROPERTIES, node.getSymbolicName())
}

func FunctionPropertiesByRelationship(relationship Relationship) FunctionInvocation {
	if relationship.getError() != nil {
		return FunctionInvocationError(relationship.err)
	}
	if !relationship.isNotNil() {
		return FunctionInvocationError(errors.New("relationship is required"))
	}
	return FunctionInvocationCreate(PROPERTIES, relationship.getSymbolicName())
}

func FunctionPropertiesByMapExpression(mapExpression MapExpression) FunctionInvocation {
	return FunctionInvocationCreate(PROPERTIES, mapExpression)
}

func FunctionCoalesce(expression ...Expression) FunctionInvocation {
	return FunctionInvocationCreate(COALESCE, expression...)
}

func FunctionToLower(expression Expression) FunctionInvocation {
	return FunctionInvocationCreate(TO_LOWER, expression)
}

func FunctionSize(expression Expression) FunctionInvocation {
	return FunctionInvocationCreate(SIZE, expression)
}

func FunctionSizeByPattern(pattern RelationshipPattern) FunctionInvocation {
	return FunctionInvocationCreateWithPatternElement(SIZE, pattern)
}

func FunctionExists(expression Expression) FunctionInvocation {
	return FunctionInvocationCreate(EXISTS, expression)
}

func FunctionDistance(point1 Expression, point2 Expression) FunctionInvocation {
	if point1 != nil && point1.getError() != nil {
		return FunctionInvocationError(point1.getError())
	}
	if point2 != nil && point2.getError() != nil {
		return FunctionInvocationError(point2.getError())
	}
	if point1 == nil || !point1.isNotNil() {
		return FunctionInvocationError(errors.New("two points is required"))
	}
	if point2 == nil || !point2.isNotNil() {
		return FunctionInvocationError(errors.New("two points is required"))
	}
	return FunctionInvocationCreate(DISTANCE, point1, point2)
}

func FunctionPoint(parameterMap MapExpression) FunctionInvocation {
	if parameterMap.getError() != nil {
		return FunctionInvocationError(parameterMap.getError())
	}
	return FunctionInvocationCreate(POINT, parameterMap)
}

func FunctionPointByParameter(parameter Parameter) FunctionInvocation {
	if parameter.getError() != nil {
		return FunctionInvocationError(parameter.getError())
	}
	return FunctionInvocationCreate(POINT, parameter)
}

func FunctionAvg(expression Expression) FunctionInvocation {
	if expression != nil && expression.getError() != nil {
		return FunctionInvocationError(expression.getError())
	}
	return FunctionInvocationCreate(AVG, expression)
}

func FunctionAvgDistinct(expression Expression) FunctionInvocation {
	if expression != nil && expression.getError() != nil {
		return FunctionInvocationError(expression.getError())
	}
	return FunctionInvocationCreateDistinct(AVG, expression)
}

/**
 * Creates a function invocation for the {@code collect()} function.
 * See <a href="https://neo4j.com/docs/cypher-manual/current/functions/aggregating/#functions-collect">collect</a>.
 *
 * @param expression The things to collect
 * @return A function call for {@code collect()}
 */
func FunctionCollect(expression Expression) FunctionInvocation {
	if expression != nil && expression.getError() != nil {
		return FunctionInvocationError(expression.getError())
	}
	return FunctionInvocationCreate(AVG, expression)
}

/**
 * Creates a function invocation for the {@code collect()} function with {@code DISTINCT} added.
 * See <a href="https://neo4j.com/docs/cypher-manual/current/functions/aggregating/#functions-collect">collect</a>.
 *
 * @param expression The things to collect
 * @return A function call for {@code collect()}
 */
func FunctionCollectDistinct(expression Expression) FunctionInvocation {
	if expression != nil && expression.getError() != nil {
		return FunctionInvocationError(expression.getError())
	}
	return FunctionInvocationCreateDistinct(AVG, expression)
}

func FunctionMax(expression Expression) FunctionInvocation {
	if expression != nil && expression.getError() != nil {
		return FunctionInvocationError(expression.getError())
	}
	return FunctionInvocationCreate(MAX, expression)
}

func FunctionMaxDistinct(expression Expression) FunctionInvocation {
	if expression != nil && expression.getError() != nil {
		return FunctionInvocationError(expression.getError())
	}
	return FunctionInvocationCreateDistinct(MAX, expression)
}

func FunctionMin(expression Expression) FunctionInvocation {
	if expression != nil && expression.getError() != nil {
		return FunctionInvocationError(expression.getError())
	}
	return FunctionInvocationCreate(MIN, expression)
}

func FunctionMinDistinct(expression Expression) FunctionInvocation {
	if expression != nil && expression.getError() != nil {
		return FunctionInvocationError(expression.getError())
	}
	return FunctionInvocationCreateDistinct(MIN, expression)
}

func FunctionPercentileCont(expression Expression, percentile float64) FunctionInvocation {
	if expression != nil && expression.getError() != nil {
		return FunctionInvocationError(expression.getError())
	}
	if expression == nil || !expression.isNotNil() {
		return FunctionInvocationError(errors.New("functions percentile cont: numeric expression for " + PERCENTILE_CONT.getImplementationName() + " is required."))
	}
	if percentile < 0.0 || percentile > 1.0 {
		return FunctionInvocationError(errors.New("functions percentile cont: the percentile for " + PERCENTILE_CONT.getImplementationName() + " must be between 0.0 and 1.0."))
	}
	return FunctionInvocationCreate(PERCENTILE_CONT, expression, NumberLiteralCreate2(percentile))
}

func FunctionPercentileContDistinct(expression Expression, percentile float64) FunctionInvocation {
	if expression != nil && expression.getError() != nil {
		return FunctionInvocationError(expression.getError())
	}
	if expression == nil || !expression.isNotNil() {
		return FunctionInvocationError(errors.New("functions percentile cont: numeric expression for " + PERCENTILE_CONT.getImplementationName() + " is required."))
	}
	if percentile < 0.0 || percentile > 1.0 {
		return FunctionInvocationError(errors.New("functions percentile cont: the percentile for " + PERCENTILE_CONT.getImplementationName() + " must be between 0.0 and 1.0."))
	}
	return FunctionInvocationCreateDistinct(PERCENTILE_CONT, expression, NumberLiteralCreate2(percentile))
}

func FunctionPercentileDisc(expression Expression, percentile float64) FunctionInvocation {
	if expression != nil && expression.getError() != nil {
		return FunctionInvocationError(expression.getError())
	}
	if expression == nil || !expression.isNotNil() {
		return FunctionInvocationError(errors.New("functions percentile cont: numeric expression for " + PERCENTILE_DISC.getImplementationName() + " is required."))
	}
	if percentile < 0.0 || percentile > 1.0 {
		return FunctionInvocationError(errors.New("functions percentile cont: the percentile for " + PERCENTILE_DISC.getImplementationName() + " must be between 0.0 and 1.0."))
	}
	return FunctionInvocationCreate(PERCENTILE_DISC, expression, NumberLiteralCreate2(percentile))
}

func FunctionPercentileDiscDistinct(expression Expression, percentile float64) FunctionInvocation {
	if expression != nil && expression.getError() != nil {
		return FunctionInvocationError(expression.getError())
	}
	if expression == nil || !expression.isNotNil() {
		return FunctionInvocationError(errors.New("functions percentile cont: numeric expression for " + PERCENTILE_DISC.getImplementationName() + " is required."))
	}
	if percentile < 0.0 || percentile > 1.0 {
		return FunctionInvocationError(errors.New("functions percentile cont: the percentile for " + PERCENTILE_DISC.getImplementationName() + " must be between 0.0 and 1.0."))
	}
	return FunctionInvocationCreateDistinct(PERCENTILE_DISC, expression, NumberLiteralCreate2(percentile))
}

func FunctionStDev(expression Expression) FunctionInvocation {
	if expression != nil && expression.getError() != nil {
		return FunctionInvocationError(expression.getError())
	}
	return FunctionInvocationCreate(ST_DEV, expression)
}

func FunctionStDevDistinct(expression Expression) FunctionInvocation {
	if expression != nil && expression.getError() != nil {
		return FunctionInvocationError(expression.getError())
	}
	return FunctionInvocationCreateDistinct(ST_DEV, expression)
}

func FunctionStDevP(expression Expression) FunctionInvocation {
	if expression != nil && expression.getError() != nil {
		return FunctionInvocationError(expression.getError())
	}
	return FunctionInvocationCreate(ST_DEV_P, expression)
}

func FunctionStDevPDistinct(expression Expression) FunctionInvocation {
	if expression != nil && expression.getError() != nil {
		return FunctionInvocationError(expression.getError())
	}
	return FunctionInvocationCreateDistinct(ST_DEV_P, expression)
}

func FunctionSum(expression Expression) FunctionInvocation {
	if expression != nil && expression.getError() != nil {
		return FunctionInvocationError(expression.getError())
	}
	return FunctionInvocationCreate(SUM, expression)
}

func FunctionSumDistinct(expression Expression) FunctionInvocation {
	if expression != nil && expression.getError() != nil {
		return FunctionInvocationError(expression.getError())
	}
	return FunctionInvocationCreateDistinct(SUM, expression)
}

func FunctionRange2Raw(start int, end int) FunctionInvocation {
	return FunctionRange2(CypherLiteralOf(start), CypherLiteralOf(end))
}

func FunctionRange2(start Expression, end Expression) FunctionInvocation {
	return FunctionRange3(start, end, nil)
}

func FunctionRange3Raw(start int, end int, step int) FunctionInvocation {
	return FunctionRange3(CypherLiteralOf(start), CypherLiteralOf(end), CypherLiteralOf(step))
}

func FunctionRange3(start Expression, end Expression, step Expression) FunctionInvocation {
	if start != nil && start.getError() != nil {
		return FunctionInvocationError(start.getError())
	}
	if end != nil && end.getError() != nil {
		return FunctionInvocationError(end.getError())
	}
	if step != nil && step.getError() != nil {
		return FunctionInvocationError(step.getError())
	}
	if start == nil || !start.isNotNil() {
		return FunctionInvocationError(errors.New("functions range: start for range is required"))
	}
	if end == nil || !end.isNotNil() {
		return FunctionInvocationError(errors.New("functions range: end for range is required"))
	}
	if step == nil || !step.isNotNil() {
		return FunctionInvocationCreate(RANGE, start, end)
	} else {
		return FunctionInvocationCreate(RANGE, start, end, step)
	}
}

func FunctionHead(expression Expression) FunctionInvocation {
	if expression != nil && expression.getError() != nil {
		return FunctionInvocationError(expression.getError())
	}
	return FunctionInvocationCreate(HEAD, expression)
}

func FunctionLast(expression Expression) FunctionInvocation {
	if expression != nil && expression.getError() != nil {
		return FunctionInvocationError(expression.getError())
	}
	return FunctionInvocationCreate(LAST, expression)
}

func FunctionNodes(path NamedPath) FunctionInvocation {
	if path.getError() != nil {
		return FunctionInvocationError(path.getError())
	}
	if !path.isNotNil() {
		return FunctionInvocationError(errors.New("functions nodes : path for nodes is required"))
	}
	symbolicName := path.getRequiredSymbolicName()
	if symbolicName.getError() != nil {
		return FunctionInvocationError(errors.New("functions nodes : path need to be named"))
	}
	return FunctionInvocationCreate(NODES, symbolicName)
}

//TODO: implement more create FunctionInvocation function
