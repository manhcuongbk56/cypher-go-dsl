package cypher

type Create struct {
	pattern Pattern
	key     string
	notNil  bool
	err     error
}

func CreateCreate(pattern Pattern) Create {
	if pattern.GetError() != nil {
		return CreateError(pattern.GetError())
	}
	c := Create{
		pattern: pattern,
		notNil:  true,
	}
	c.key = getAddress(&c)
	return c
}

func CreateError(err error) Create {
	return Create{
		err: err,
	}
}

func (c Create) GetError() error {
	return c.err
}

func (c Create) accept(visitor *CypherRenderer) {
	visitor.enter(c)
	c.pattern.accept(visitor)
	visitor.leave(c)
}

func (c Create) enter(renderer *CypherRenderer) {
	renderer.append("CREATE ")
}

func (c Create) leave(renderer *CypherRenderer) {
	renderer.append(" ")
}

func (c Create) getKey() string {
	return c.key
}

func (c Create) isNotNil() bool {
	return c.notNil
}

func (c Create) isUpdatingClause() bool {
	return true
}
