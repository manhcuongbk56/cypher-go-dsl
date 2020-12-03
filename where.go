package cypher_go_dsl

import "fmt"

type Where struct {
	key       string
	condition Condition
	notNil    bool
}

func WhereCreate(condition Condition) Where {
	return Where{
		condition: condition,
		notNil:    true,
	}
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
	w.accept(visitor)
	(*visitor).leave(w)
}

func (w Where) enter(renderer *CypherRenderer) {
	panic("implement me")
}

func (w Where) leave(renderer *CypherRenderer) {
	panic("implement me")
}
