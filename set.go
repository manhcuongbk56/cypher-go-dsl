package cypher_go_dsl

type Set struct {
	setItems ExpressionList
	key      string
	notNil   bool
	err      error
}

func SetCreate(setItems ExpressionList) Set {
	if setItems.getError() != nil {
		return SetError(setItems.getError())
	}
	set := Set{
		setItems: setItems,
		notNil:   true,
	}
	set.key = getAddress(&set)
	return set
}

func SetError(err error) Set {
	return Set{
		err: err,
	}
}

func (s Set) isUpdatingClause() bool {
	return true
}

func (s Set) getError() error {
	return s.err
}

func (s Set) accept(visitor *CypherRenderer) {
	visitor.enter(s)
	s.setItems.accept(visitor)
	visitor.leave(s)
}

func (s Set) enter(renderer *CypherRenderer) {
	renderer.append("SET ")
}

func (s Set) leave(renderer *CypherRenderer) {
	renderer.append(" ")
}

func (s Set) getKey() string {
	return s.key
}

func (s Set) isNotNil() bool {
	return s.notNil
}
