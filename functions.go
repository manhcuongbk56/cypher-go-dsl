package cypher_go_dsl

import "errors"

func idByNode(node Node) FunctionInvocation {
	if node.getError() != nil {
		return FunctionInvocationError(node.err)
	}
	if !node.isNotNil() {
		return FunctionInvocationError(errors.New("node is required"))
	}
	return FunctionInvocationCreate(ID, node.getSymbolicName())
}

func idByRelationship(relationship Relationship) FunctionInvocation {
	if relationship.getError() != nil {
		return FunctionInvocationError(relationship.err)
	}
	if !relationship.isNotNil() {
		return FunctionInvocationError(errors.New("relationship is required"))
	}
	return FunctionInvocationCreate(ID, relationship.getSymbolicName())
}

func labels(node Node) FunctionInvocation {
	if node.getError() != nil {
		return FunctionInvocationError(node.err)
	}
	if !node.isNotNil() {
		return FunctionInvocationError(errors.New("node is required"))
	}
	return FunctionInvocationCreate(LABELS, node.getSymbolicName())
}

func functionType(relationship Relationship) FunctionInvocation {
	if relationship.getError() != nil {
		return FunctionInvocationError(relationship.err)
	}
	if !relationship.isNotNil() {
		return FunctionInvocationError(errors.New("relationship is required"))
	}
	return FunctionInvocationCreate(TYPE, relationship.getSymbolicName())
}

func Count(node Node) FunctionInvocation {
	return FunctionInvocationCreate(COUNT, node.getSymbolicName())
}

func CountByExpression(expression Expression) FunctionInvocation {
	return FunctionInvocationCreate(COUNT, expression)
}

func CountDistinct(node Node) FunctionInvocation {
	return FunctionInvocationCreateDistinct(COUNT, node.getSymbolicName())
}

func countDistinctByExpression(expression Expression) FunctionInvocation {
	return FunctionInvocationCreateDistinct(COUNT, expression)
}

func properties(node Node) FunctionInvocation {
	if node.getError() != nil {
		return FunctionInvocationError(node.err)
	}
	if !node.isNotNil() {
		return FunctionInvocationError(errors.New("node is required"))
	}
	return FunctionInvocationCreate(PROPERTIES, node.getSymbolicName())
}

func propertiesByRelationship(relationship Relationship) FunctionInvocation {
	if relationship.getError() != nil {
		return FunctionInvocationError(relationship.err)
	}
	if !relationship.isNotNil() {
		return FunctionInvocationError(errors.New("relationship is required"))
	}
	return FunctionInvocationCreate(PROPERTIES, relationship.getSymbolicName())
}

func propertiesByMapExpression(mapExpression MapExpression) FunctionInvocation {
	return FunctionInvocationCreate(PROPERTIES, mapExpression)
}

func coalesce(expression ...Expression) FunctionInvocation {
	return FunctionInvocationCreate(COALESCE, expression...)
}

func toLower(expression Expression) FunctionInvocation {
	return FunctionInvocationCreate(TO_LOWER, expression)
}

func size(expression Expression) FunctionInvocation {
	return FunctionInvocationCreate(SIZE, expression)
}

func sizeByPattern(pattern RelationshipPattern) FunctionInvocation {
	return FunctionInvocationCreateWithPatternElement(SIZE, pattern)
}

func exists(expression Expression) FunctionInvocation {
	return FunctionInvocationCreate(EXISTS, expression)
}

func distance(point1 Expression, point2 Expression) FunctionInvocation {
	if point1.getError() != nil {
		return FunctionInvocationError(point1.getError())
	}
	if !point1.isNotNil() {
		return FunctionInvocationError(errors.New("two points is required"))
	}
	if point2.getError() != nil {
		return FunctionInvocationError(point2.getError())
	}
	if !point2.isNotNil() {
		return FunctionInvocationError(errors.New("two points is required"))
	}
	return FunctionInvocationCreate(DISTANCE, point1, point2)
}

func point(parameterMap MapExpression) FunctionInvocation {
	return FunctionInvocationCreate(POINT, parameterMap)
}

func pointByParameter(parameter Parameter) FunctionInvocation {
	return FunctionInvocationCreate(POINT, parameter)
}

//TODO: implement more create FunctionInvocation function
