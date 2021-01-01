package cypher

func NameOrExpression(expression Expression) Expression {
	named, isNamed := (expression).(Named)
	if isNamed && named.GetSymbolicName().isNotNil() {
		return named.GetSymbolicName()
	}
	return expression
}

func CreateSymbolicNameByNamed(nameds ...Named) []Expression {
	expressions := make([]Expression, 0)
	for _, named := range nameds {
		if named != nil {
			expressions = append(expressions, named.GetSymbolicName())
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
