package cypher_go_dsl

func NameOrExpression(expression Expression) Expression {
	named, isNamed := (expression).(Named)
	if isNamed && named.getSymbolicName() != nil {
		return named.getSymbolicName()
	}
	return expression
}

func CreateSymbolicNames(namedOrString interface{}) Expression {
	if named, isNamed := namedOrString.(Named); isNamed {
		return named.getSymbolicName()
	}
	if stringArg, isString := namedOrString.(string); isString {
		return SymbolicNameCreate(stringArg)
	}
	return nil
}
