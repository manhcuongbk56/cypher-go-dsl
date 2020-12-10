package cypher_go_dsl

import "fmt"

type Where struct {
	key       string
	condition Condition
	notNil    bool
	err       error
}

func WhereCreate(condition Condition) Where {
	where := Where{
		condition: condition,
		notNil:    true,
	}
	where.key = getAddress(&where)
	return where
}

func (w Where) getError() error {
	return w.err
}

func (w Where) isNotNil() bool {
	return w.notNil
}

func (w Where) getKey() string {
	return w.key
}

func (w Where) accept(visitor *CypherRenderer) {
	w.key = fmt.Sprint(&w)
	(*visitor).enter(w)
	w.condition.accept(visitor)
	(*visitor).leave(w)
}

func (w Where) enter(renderer *CypherRenderer) {
	renderer.append(" WHERE ")
}

func (w Where) leave(renderer *CypherRenderer) {
}
