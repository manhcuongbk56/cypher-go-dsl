package cypher

type Remove struct {
	setItems ExpressionList
	key      string
	notNil   bool
	err      error
}

func RemoveCreate(setItems ExpressionList) Remove {
	if setItems.getError() != nil {
		return RemoveError(setItems.getError())
	}
	r := Remove{
		setItems: setItems,
		notNil:   true,
	}
	r.key = getAddress(&r)
	return r
}

func RemoveError(err error) Remove {
	return Remove{
		err: err,
	}
}

func (r Remove) getError() error {
	return r.err
}

func (r Remove) accept(visitor *CypherRenderer) {
	visitor.enter(r)
	r.setItems.accept(visitor)
	visitor.leave(r)
}

func (r Remove) enter(renderer *CypherRenderer) {
	renderer.append("REMOVE ")
}

func (r Remove) leave(renderer *CypherRenderer) {
	renderer.append(" ")
}

func (r Remove) getKey() string {
	return r.key
}

func (r Remove) isNotNil() bool {
	return r.notNil
}

func (r Remove) isUpdatingClause() bool {
	panic("implement me")
}
