package cypher

func NameOrExpression(expression Expression) Expression {
	named, isNamed := (expression).(Named)
	if isNamed && named.getSymbolicName().isNotNil() {
		return named.getSymbolicName()
	}
	return expression
}

func CreateSymbolicNameByNamed(nameds ...Named) []Expression {
	expressions := make([]Expression, 0)
	for _, named := range nameds {
		if named != nil {
			expressions = append(expressions, named.getSymbolicName())
		}
	}
	return expressions
}
func CreateSymbolicNameByString(args ...string) []Expression {
	expressions := make([]Expression, 0)
	for _, arg := range args {
		if arg != "" {
			expressions = append(expressions, SymbolicNameCreate(arg))
		}
	}
	return expressions
}
