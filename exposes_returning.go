package cypher_go_dsl

type ExposesReturning interface {
	returning(expression ...Expression) OngoingReadingAndReturn
	returningDistinct(expression ...Expression) OngoingReadingAndReturn
}

type ExposesReturningOne interface {
	returning(expression ...Expression) OngoingReadingAndReturn
}

type ExposesReturningStruct struct {
	ExposesReturningOne
}

func (e ExposesReturningStruct) returning(args ...interface{}) OngoingReadingAndReturn {
	expressions := make([]Expression, 0)
	for _, arg := range args{
		if _, isString := arg.(string); isString  {
			expressions = append(expressions, CreateSymbolicNames(arg))
			continue
		}
		if _, isNamed := arg.(Named); isNamed  {
			expressions = append(expressions, CreateSymbolicNames(arg))
		}
	}
	return e.ExposesReturningOne.returning(expressions...)
}
