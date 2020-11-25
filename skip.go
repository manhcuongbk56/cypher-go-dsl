package cypher_go_dsl

type Skip struct {
	skipAmount NumberLiteral
}

func (s Skip) Accept(visitor *CypherRenderer) {
	(*visitor).Enter(s)
	s.skipAmount.Accept(visitor)
	(*visitor).Leave(s)
}

func (s Skip) GetType() VisitableType {
	return SkipVisitable
}

func CreateSkip(number int)  Skip{
	literal := NumberLiteral{
		content: number,
	}
	return Skip{skipAmount: literal}
}

func (s Skip) Enter(renderer *CypherRenderer) {
	renderer.builder.WriteString(" SKIP ")
}

func (s Skip) Leave(renderer *CypherRenderer) {
}


