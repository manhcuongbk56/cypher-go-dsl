package cypher_go_dsl

import "errors"

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

func FunctionPoint(parameterMap MapExpression) FunctionInvocation {
	return FunctionInvocationCreate(POINT, parameterMap)
}

func FunctionPointByParameter(parameter Parameter) FunctionInvocation {
	return FunctionInvocationCreate(POINT, parameter)
}

//TODO: implement more create FunctionInvocation function
