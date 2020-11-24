package cypher_go_dsl

type Where struct {
	condition Condition
}

func (w Where) Accept(visitor Visitor) {
	visitor.Enter(w)
	w.Accept(visitor)
	visitor.Leave(w)
}

func (w Where) GetType() VisitableType {
	return WhereVisitable
}


