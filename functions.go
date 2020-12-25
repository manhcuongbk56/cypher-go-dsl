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

/**
 * Creates a function invocation for the {@code max()} function.
 * See <a href="https://neo4j.com/docs/cypher-manual/current/functions/aggregating/#functions-max">max</a>.
 *
 * @param expression A list from which the maximum element value is returned
 * @return A function call for {@code max()}
 */
func FunctionMax(expression Expression) FunctionInvocation {
	if expression != nil && expression.getError() != nil {
		return FunctionInvocationError(expression.getError())
	}
	return FunctionInvocationCreate(MAX, expression)
}

/**
 * Creates a function invocation for the {@code max()} function with {@code DISTINCT} added.
 * See <a href="https://neo4j.com/docs/cypher-manual/current/functions/aggregating/#functions-max">max</a>.
 *
 * @param expression A list from which the maximum element value is returned
 * @return A function call for {@code max()}
 */
func FunctionMaxDistinct(expression Expression) FunctionInvocation {
	if expression != nil && expression.getError() != nil {
		return FunctionInvocationError(expression.getError())
	}
	return FunctionInvocationCreateDistinct(MAX, expression)
}

/**
 * Creates a function invocation for the {@code min()} function.
 * See <a href="https://neo4j.com/docs/cypher-manual/current/functions/aggregating/#functions-min">min</a>.
 *
 * @param expression A list from which the minimum element value is returned
 * @return A function call for {@code min()}
 */
func FunctionMin(expression Expression) FunctionInvocation {
	if expression != nil && expression.getError() != nil {
		return FunctionInvocationError(expression.getError())
	}
	return FunctionInvocationCreate(MIN, expression)
}

/**
 * Creates a function invocation for the {@code min()} function with {@code DISTINCT} added.
 * See <a href="https://neo4j.com/docs/cypher-manual/current/functions/aggregating/#functions-min">min</a>.
 *
 * @param expression A list from which the minimum element value is returned
 * @return A function call for {@code min()}
 */
func FunctionMinDistinct(expression Expression) FunctionInvocation {
	if expression != nil && expression.getError() != nil {
		return FunctionInvocationError(expression.getError())
	}
	return FunctionInvocationCreateDistinct(MIN, expression)
}

/**
 * Creates a function invocation for the {@code percentileCont()} function.
 * See <a href="https://neo4j.com/docs/cypher-manual/current/functions/aggregating/#functions-percentilecont">percentileCont</a>.
 *
 * @param expression A numeric expression
 * @param percentile A numeric value between 0.0 and 1.0
 * @return A function call for {@code percentileCont()}
 */
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

/**
 * Creates a function invocation for the {@code percentileCont()} function with {@code DISTINCT} added.
 * See <a href="https://neo4j.com/docs/cypher-manual/current/functions/aggregating/#functions-percentilecont">percentileCont</a>.
 *
 * @param expression A numeric expression
 * @param percentile A numeric value between 0.0 and 1.0
 * @return A function call for {@code percentileCont()}
 */
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

/**
 * Creates a function invocation for the {@code percentileDisc()} function.
 * See <a href="https://neo4j.com/docs/cypher-manual/current/functions/aggregating/#functions-percentiledisc">percentileDisc</a>.
 *
 * @param expression A numeric expression
 * @param percentile A numeric value between 0.0 and 1.0
 * @return A function call for {@code percentileDisc()}
 */
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

/**
 * Creates a function invocation for the {@code percentileDisc()} function with {@code DISTINCT} added.
 * See <a href="https://neo4j.com/docs/cypher-manual/current/functions/aggregating/#functions-percentiledisc">percentileDisc</a>.
 *
 * @param expression A numeric expression
 * @param percentile A numeric value between 0.0 and 1.0
 * @return A function call for {@code percentileDisc()}
 */
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

/**
 * Creates a function invocation for the {@code stDev()} function.
 * See <a href="https://neo4j.com/docs/cypher-manual/current/functions/aggregating/#functions-stdev">stDev</a>.
 *
 * @param expression A numeric expression
 * @return A function call for {@code stDev()}
 */
func FunctionStDev(expression Expression) FunctionInvocation {
	if expression != nil && expression.getError() != nil {
		return FunctionInvocationError(expression.getError())
	}
	return FunctionInvocationCreate(ST_DEV, expression)
}

/**
 * Creates a function invocation for the {@code stDev()} function with {@code DISTINCT} added.
 * See <a href="https://neo4j.com/docs/cypher-manual/current/functions/aggregating/#functions-stdev">stDev</a>.
 *
 * @param expression A numeric expression
 * @return A function call for {@code stDev()}
 */
func FunctionStDevDistinct(expression Expression) FunctionInvocation {
	if expression != nil && expression.getError() != nil {
		return FunctionInvocationError(expression.getError())
	}
	return FunctionInvocationCreateDistinct(ST_DEV, expression)
}

/**
 * Creates a function invocation for the {@code stDevP()} function.
 * See <a href="https://neo4j.com/docs/cypher-manual/current/functions/aggregating/#functions-stdevp">stDevP</a>.
 *
 * @param expression A numeric expression
 * @return A function call for {@code stDevP()}
 */
func FunctionStDevP(expression Expression) FunctionInvocation {
	if expression != nil && expression.getError() != nil {
		return FunctionInvocationError(expression.getError())
	}
	return FunctionInvocationCreate(ST_DEV_P, expression)
}

/**
 * Creates a function invocation for the {@code stDevP()} function with {@code DISTINCT} added.
 * See <a href="https://neo4j.com/docs/cypher-manual/current/functions/aggregating/#functions-stdevp">stDevP</a>.
 *
 * @param expression A numeric expression
 * @return A function call for {@code stDevP()}
 */
func FunctionStDevPDistinct(expression Expression) FunctionInvocation {
	if expression != nil && expression.getError() != nil {
		return FunctionInvocationError(expression.getError())
	}
	return FunctionInvocationCreateDistinct(ST_DEV_P, expression)
}

/**
 * Creates a function invocation for the {@code sum()} function.
 * See <a href="https://neo4j.com/docs/cypher-manual/current/functions/aggregating/#functions-sum">sum</a>.
 *
 * @param expression An expression returning a set of numeric values
 * @return A function call for {@code sum()}
 */
func FunctionSum(expression Expression) FunctionInvocation {
	if expression != nil && expression.getError() != nil {
		return FunctionInvocationError(expression.getError())
	}
	return FunctionInvocationCreate(SUM, expression)
}

/**
 * Creates a function invocation for the {@code sum()} function  with {@code DISTINCT} added.
 * See <a href="https://neo4j.com/docs/cypher-manual/current/functions/aggregating/#functions-sum">sum</a>.
 *
 * @param expression An expression returning a set of numeric values
 * @return A function call for {@code sum()}
 */
func FunctionSumDistinct(expression Expression) FunctionInvocation {
	if expression != nil && expression.getError() != nil {
		return FunctionInvocationError(expression.getError())
	}
	return FunctionInvocationCreateDistinct(SUM, expression)
}

/**
 * @param start the range's start
 * @param end   the range's end
 * @return A function call for {@code range()}
 * @see #range(Expression, Expression)
 */
func FunctionRange2Raw(start int, end int) FunctionInvocation {
	return FunctionRange2(LiteralOf(start), LiteralOf(end))
}

/**
 * @param start the range's start
 * @param end   the range's end
 * @return A function call for {@code range()}
 * @see #range(Expression, Expression, Expression)
 */
func FunctionRange2(start Expression, end Expression) FunctionInvocation {
	return FunctionRange3(start, end, nil)
}

/**
 * Creates a function invocation for the {@code range()} function.
 * See <a href="https://neo4j.com/docs/cypher-manual/current/functions/list/#functions-range">range</a>.
 *
 * @param start the range's start
 * @param end   the range's end
 * @param step  the range's step
 * @return A function call for {@code range()}
 * @see #range(Expression, Expression, Expression)
 */
func FunctionRange3Raw(start int, end int, step int) FunctionInvocation {
	return FunctionRange3(LiteralOf(start), LiteralOf(end), LiteralOf(step))
}

/**
 * Creates a function invocation for the {@code range()} function.
 * See <a href="https://neo4j.com/docs/cypher-manual/current/functions/list/#functions-range">range</a>.
 *
 * @param start the range's start
 * @param end   the range's end
 * @param step  the range's step
 * @return A function call for {@code range()}
 */
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

/**
 * Creates a function invocation for the {@code head()} function.
 * See <a href="https://neo4j.com/docs/cypher-manual/current/functions/scalar/#functions-head">head</a>.
 *
 * @param expression A list from which the head element is returned
 * @return A function call for {@code head()}
 */
func FunctionHead(expression Expression) FunctionInvocation {
	if expression != nil && expression.getError() != nil {
		return FunctionInvocationError(expression.getError())
	}
	return FunctionInvocationCreate(HEAD, expression)
}

/**
 * Creates a function invocation for the {@code last()} function.
 * See <a href="https://neo4j.com/docs/cypher-manual/current/functions/scalar/#functions-last">last</a>.
 *
 * @param expression A list from which the last element is returned
 * @return A function call for {@code last()}
 */
func FunctionLast(expression Expression) FunctionInvocation {
	if expression != nil && expression.getError() != nil {
		return FunctionInvocationError(expression.getError())
	}
	return FunctionInvocationCreate(LAST, expression)
}

/**
 * Creates a function invocation for {@code nodes{}}.
 * See <a href="https://neo4j.com/docs/cypher-manual/current/functions/list/#functions-nodes">nodes</a>.
 *
 * @param path The path for which the number of nodes should be retrieved
 * @return A function call for {@code nodes()} on a path.
 * @since 1.1
 */
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

/**
 * Creates a function invocation for {@code relationships{}}.
 * See <a href="https://neo4j.com/docs/cypher-manual/current/functions/list/#functions-relationships">relationships</a>.
 *
 * @param path The path for which the relationships should be retrieved
 * @return A function call for {@code relationships()} on a path.
 * @since 2020.0.2
 */
func FunctionRelationships(path NamedPath) FunctionInvocation {
	if path.getError() != nil {
		return FunctionInvocationError(path.getError())
	}
	if !path.isNotNil() {
		return FunctionInvocationError(errors.New("functions relationships : path for relationships is required"))
	}
	symbolicName := path.getRequiredSymbolicName()
	if symbolicName.getError() != nil {
		return FunctionInvocationError(errors.New("functions relationships : path need to be named"))
	}
	return FunctionInvocationCreate(RELATIONSHIPS, symbolicName)
}

/**
 * Creates a function invocation for {@code startNode{}}.
 * See <a href="https://neo4j.com/docs/cypher-manual/current/functions/scalar/#functions-startnode">startNode</a>.
 *
 * @param relationship The relationship for which the start node be retrieved
 * @return A function call for {@code startNode()} on a path.
 * @since 2020.0.2
 */
func FunctionStartNode(relationship Relationship) FunctionInvocation {
	if relationship.getError() != nil {
		return FunctionInvocationError(relationship.getError())
	}
	if !relationship.isNotNil() {
		return FunctionInvocationError(errors.New("functions start node : relationship for start node is required"))
	}
	symbolicName := relationship.getRequiredSymbolicName()
	if symbolicName.getError() != nil {
		return FunctionInvocationError(errors.New("functions start node : relationship need to be named"))
	}
	return FunctionInvocationCreate(START_NODE, symbolicName)
}

/**
 * Creates a function invocation for {@code endNode{}}.
 * See <a href="https://neo4j.com/docs/cypher-manual/current/functions/scalar/#functions-endnode">endNode</a>.
 *
 * @param relationship The relationship for which the end node be retrieved
 * @return A function call for {@code endNode()} on a path.
 * @since 2020.0.2
 */
func FunctionEndNode(relationship Relationship) FunctionInvocation {
	if relationship.getError() != nil {
		return FunctionInvocationError(relationship.getError())
	}
	if !relationship.isNotNil() {
		return FunctionInvocationError(errors.New("functions start node : relationship for end node is required"))
	}
	symbolicName := relationship.getRequiredSymbolicName()
	if symbolicName.getError() != nil {
		return FunctionInvocationError(errors.New("functions start node : relationship need to be named"))
	}
	return FunctionInvocationCreate(END_NODE, symbolicName)
}

/**
 * Creates a function invocation for {@code date()}.
 * See <a href="https://neo4j.com/docs/cypher-manual/current/functions/temporal/date/">date</a>.
 * This is the most simple form.
 *
 * @return A function call for {@code date()}.
 * @since 2020.1.0
 */
func FunctionDate() FunctionInvocation {
	return FunctionInvocationCreate1(DATE)
}

/**
 * Creates a function invocation for {@code date({})}.
 * See <a href="https://neo4j.com/docs/cypher-manual/current/functions/temporal/date/">date</a>.
 *
 * @param year  The year
 * @param month The month
 * @param day   The day
 * @return A function call for {@code date({})}.
 * @since 2020.1.0
 */
func FunctionCalendarDate(year int, month int, day int) FunctionInvocation {
	return FunctionInvocationCreate(DATE, MapOf("year", LiteralOf(year), "month", LiteralOf(month), "day", LiteralOf(day)))
}

//TODO: implement more create FunctionInvocation function
