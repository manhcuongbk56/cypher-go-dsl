package cypher_go_dsl

import "fmt"

type Create struct {
	pattern Pattern
	key     string
	notNil  bool
	err error
}

func CreateCreate(pattern Pattern) Create {
	return Create{
		pattern: pattern,
		notNil:  true,
	}
}

func (c Create) getError() error {
	return c.err
}

func (c Create) accept(visitor *CypherRenderer) {
	c.key = fmt.Sprint(&c)
	visitor.enter(c)
	c.pattern.accept(visitor)
	visitor.leave(c)
}

func (c Create) enter(renderer *CypherRenderer) {
	renderer.builder.WriteString("CREATE ")
}

func (c Create) leave(renderer *CypherRenderer) {
	renderer.builder.WriteString(" ")
}

func (c Create) getKey() string {
	return c.key
}

func (c Create) isNotNil() bool {
	return c.notNil
}

func (c Create) isUpdatingClause() bool {
	panic("implement me")
}
