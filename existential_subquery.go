package cypher_go_dsl

type ExistentialSubquery struct {
	fragment Match
	key      string
	notNil   bool
	err      error
}

func ExistentialSubqueryCreate(fragment Match) ExistentialSubquery {
	if fragment.getError() != nil {
		return ExistentialSubqueryError(fragment.getError())
	}
	e := ExistentialSubquery{
		fragment: fragment,
		notNil:   true,
	}
	e.key = getAddress(&e)
	return e
}

func ExistentialSubqueryError(err error) ExistentialSubquery {
	return ExistentialSubquery{
		err: err,
	}
}

func ExistentialSubqueryExists(fragment Match) ExistentialSubquery {
	return ExistentialSubqueryCreate(fragment)
}

func (e ExistentialSubquery) getError() error {
	return e.err
}

func (e ExistentialSubquery) accept(visitor *CypherRenderer) {
	visitor.enter(e)
	e.fragment.accept(visitor)
	visitor.leave(e)
}

func (e ExistentialSubquery) enter(renderer *CypherRenderer) {
	renderer.append("EXISTS {")
}

func (e ExistentialSubquery) leave(renderer *CypherRenderer) {
	//FIXME: It may be wrong.
	renderer.append("}")
}

func (e ExistentialSubquery) getKey() string {
	return e.key
}

func (e ExistentialSubquery) isNotNil() bool {
	return e.notNil
}

func (e ExistentialSubquery) GetExpressionType() ExpressionType {
	panic("implement me")
}

func (e ExistentialSubquery) getConditionType() string {
	return "ExistentialSubquery"
}
