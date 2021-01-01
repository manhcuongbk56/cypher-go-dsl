package cypher

type EntryExpression struct {
	ExpressionContainer
	Key    string
	Value  Expression
	key    string
	notNil bool
	err    error
}

func EntryExpressionCreate(key string, value Expression) EntryExpression {
	if value != nil && value.GetError() != nil {
		return EntryExpression{
			err: value.GetError(),
		}
	}
	e := EntryExpression{
		Value:  value,
		Key:    key,
		notNil: true,
	}
	e.key = getAddress(&e)
	e.ExpressionContainer = ExpressionWrap(e)
	return e
}

func (e EntryExpression) GetError() error {
	return e.err
}

func (e EntryExpression) isNotNil() bool {
	return e.notNil
}

func (e EntryExpression) getKey() string {
	return e.key
}

func (e EntryExpression) GetExpressionType() ExpressionType {
	return EXPRESSION
}

func (e EntryExpression) accept(visitor *CypherRenderer) {
	visitor.enter(e)
	e.Value.accept(visitor)
	visitor.leave(e)
}

func (e EntryExpression) enter(renderer *CypherRenderer) {
	renderer.append(escapeIfNecessary(e.Key))
	renderer.append(": ")
}

func (e EntryExpression) leave(renderer *CypherRenderer) {
}
