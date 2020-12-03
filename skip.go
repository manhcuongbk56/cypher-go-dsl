package cypher_go_dsl

import "fmt"

type Skip struct {
	skipAmount NumberLiteral
	key        string
	notNil     bool
}

func (s Skip) isNotNil() bool {
	return s.notNil
}

func (s Skip) getKey() string {
	return s.key
}

func (s Skip) accept(visitor *CypherRenderer) {
	s.key = fmt.Sprint(&s)
	(*visitor).enter(s)
	s.skipAmount.accept(visitor)
	(*visitor).leave(s)
}

func CreateSkip(number int) Skip {
	if number == 0 {
		return Skip{}
	}
	literal := NumberLiteral{
		content: number,
	}
	return Skip{skipAmount: literal}
}

func (s Skip) enter(renderer *CypherRenderer) {
	renderer.builder.WriteString(" SKIP ")
}

func (s Skip) leave(renderer *CypherRenderer) {
}
