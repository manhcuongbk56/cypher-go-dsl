package cypher_go_dsl

import "fmt"

type UnionPart struct {
	all    bool
	query  SingleQuery
	key    string
	notNil bool
}

func UnionPartCreate(all bool, query SingleQuery) UnionPart {
	return UnionPart{
		all:    all,
		query:  query,
		notNil: true,
	}
}

func (u UnionPart) isAll() bool {
	return u.all
}

func (u UnionPart) accept(visitor *CypherRenderer) {
	u.key = fmt.Sprint(&u)
	visitor.enter(u)
	u.query.accept(visitor)
	visitor.leave(u)
}

func (u UnionPart) enter(renderer *CypherRenderer) {
	renderer.append(" UNION ")
	if u.isAll() {
		renderer.append("ALL ")
	}
}

func (u UnionPart) leave(renderer *CypherRenderer) {
}

func (u UnionPart) getKey() string {
	return u.key
}

func (u UnionPart) isNotNil() bool {
	return u.notNil
}
