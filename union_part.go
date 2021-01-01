package cypher

type UnionPart struct {
	all    bool
	query  SingleQuery
	key    string
	notNil bool
	err    error
}

func UnionPartCreate(all bool, query SingleQuery) UnionPart {
	if query != nil && query.GetError() != nil {
		return UnionPartError(query.GetError())
	}
	unionPart := UnionPart{
		all:    all,
		query:  query,
		notNil: true,
	}
	unionPart.key = getAddress(&unionPart)
	return unionPart
}

func UnionPartError(err error) UnionPart {
	return UnionPart{
		err: err,
	}
}

func (u UnionPart) isAll() bool {
	return u.all
}

func (u UnionPart) GetError() error {
	return u.err
}

func (u UnionPart) accept(visitor *CypherRenderer) {
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
