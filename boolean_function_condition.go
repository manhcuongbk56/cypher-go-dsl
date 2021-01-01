package cypher

type BooleanFunctionCondition struct {
	ConditionContainer
	delegate FunctionInvocation
	key      string
	notNil   bool
	err      error
}

func BooleanFunctionConditionCreate(delegate FunctionInvocation) BooleanFunctionCondition {
	function := BooleanFunctionCondition{
		delegate: delegate,
		notNil:   true,
	}
	function.key = getAddress(&function)
	function.ConditionContainer = ConditionWrap(function)
	return function
}

func BooleanFunctionConditionError(err error) BooleanFunctionCondition {
	function := BooleanFunctionCondition{
		err: err,
	}
	return function
}

func (b BooleanFunctionCondition) GetError() error {
	return b.err
}

func (b BooleanFunctionCondition) accept(visitor *CypherRenderer) {
	b.delegate.accept(visitor)
}

func (b BooleanFunctionCondition) enter(renderer *CypherRenderer) {
}

func (b BooleanFunctionCondition) leave(renderer *CypherRenderer) {
}

func (b BooleanFunctionCondition) getKey() string {
	return b.key
}

func (b BooleanFunctionCondition) isNotNil() bool {
	return b.notNil
}

func (b BooleanFunctionCondition) GetExpressionType() ExpressionType {
	return "BooleanFunctionCondition"
}

func (b BooleanFunctionCondition) getConditionType() string {
	return "BooleanFunctionCondition"
}
