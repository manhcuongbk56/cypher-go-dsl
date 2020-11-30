package cypher_go_dsl

func idByNode(node Node) FunctionInvocation {
	functionInvocation, _ := FunctionInvocationCreate(ID, node.getSymbolicName())
	return functionInvocation
}

func idByRelationship(relationship Relationship) FunctionInvocation {
	functionInvocation, _ := FunctionInvocationCreate(ID, relationship.getSymbolicName())
	return functionInvocation
}

func labels(node Node) FunctionInvocation {
	functionInvocation, _ := FunctionInvocationCreate(ID, node.getSymbolicName())
	return functionInvocation
}

func size(expression Expression) FunctionInvocation {
	functionInvocation, _ := FunctionInvocationCreate(SIZE, expression)
	return functionInvocation
}
