package cypher_go_dsl

import "fmt"

type Remove struct {
	setItems ExpressionList
	key      string
	notNil   bool
}

func RemoveCreate(setItems ExpressionList) Remove {
	return Remove{
		setItems: setItems,
		notNil:   true,
	}
}

func (r Remove) accept(visitor *CypherRenderer) {
	r.key = fmt.Sprint(&r)
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
