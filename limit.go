package cypher_go_dsl

type Limit struct {
	limitAmount NumberLiteral
}

func (l Limit) Accept(visitor *CypherRenderer) {
	(*visitor).Enter(l)
	l.limitAmount.Accept(visitor)
	(*visitor).Leave(l)
}

func CreateLimit(number int) Limit {
	literal := NumberLiteral{
		content: number,
	}
	return Limit{limitAmount: literal}
}

func (l Limit) GetType() VisitableType {
	return LimitVisitable
}

func (l Limit) Enter(renderer *CypherRenderer) {
	renderer.builder.WriteString(" LIMIT ")}

func (l Limit) Leave(renderer *CypherRenderer) {
}



