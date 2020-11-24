package cypher_go_dsl

type Order struct {
	sortItems []SortItem
}

func (o Order) Accept(visitor Visitor) {
	visitor.Enter(o)
	for _, sortItem := range o.sortItems{
		o.PrepareVisit(sortItem).Accept(visitor)
	}
	visitor.Leave(o)
}

func (o Order) GetType() VisitableType {
	return OrderVisitable
}

func (o Order) PrepareVisit(visitable Visitable) Visitable {
	return visitable
}



