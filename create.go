package cypher

type create struct {
	pattern Pattern
	key     string
	notNil  bool
	err     error
}

func createcreate(pattern Pattern) create {
	if pattern.GetError() != nil {
		return createError(pattern.GetError())
	}
	c := create{
		pattern: pattern,
		notNil:  true,
	}
	c.key = getAddress(&c)
	return c
}

func createError(err error) create {
	return create{
		err: err,
	}
}

func (c create) GetError() error {
	return c.err
}

func (c create) accept(visitor *CypherRenderer) {
	visitor.enter(c)
	c.pattern.accept(visitor)
	visitor.leave(c)
}

func (c create) enter(renderer *CypherRenderer) {
	renderer.append("CREATE ")
}

func (c create) leave(renderer *CypherRenderer) {
	renderer.append(" ")
}

func (c create) getKey() string {
	return c.key
}

func (c create) isNotNil() bool {
	return c.notNil
}

func (c create) isUpdatingClause() bool {
	return true
}
