package cypher_go_dsl

type Limit struct {
	limitAmount NumberLiteral
}

func (l Limit) Accept(visitor Visitor) {
	visitor.Enter(l)
	l.limitAmount.Accept(visitor)
	visitor.Leave(l)
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


