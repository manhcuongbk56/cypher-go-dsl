package cypher

type Where struct {
	key       string
	condition Condition
	notNil    bool
	err       error
}

func WhereCreate(condition Condition) Where {
	if condition != nil && condition.GetError() != nil {
		return WhereError(condition.GetError())
	}
	where := Where{
		condition: condition,
		notNil:    true,
	}
	where.key = getAddress(&where)
	return where
}

func WhereError(err error) Where {
	return Where{
		err: err,
	}
}

func (w Where) GetError() error {
	return w.err
}

func (w Where) isNotNil() bool {
	return w.notNil
}

func (w Where) getKey() string {
	return w.key
}

func (w Where) accept(visitor *CypherRenderer) {
	visitor.enter(w)
	w.condition.accept(visitor)
	visitor.leave(w)
}

func (w Where) enter(renderer *CypherRenderer) {
	renderer.append(" WHERE ")
}

func (w Where) leave(renderer *CypherRenderer) {
}
