package cypher

type Order struct {
	sortItems []SortItem
	key       string
	notNil    bool
	err       error
}

func OrderCreate(sortItems []SortItem) Order {
	for _, item := range sortItems {
		if item.GetError() != nil {
			return Order{err: item.GetError()}
		}
	}
	o := Order{sortItems: sortItems, notNil: true}
	o.key = getAddress(&o)
	return o
}

func (o Order) GetError() error {
	return o.err
}

func (o Order) isNotNil() bool {
	return o.notNil
}

func (o Order) getKey() string {
	return o.key
}

func (o Order) accept(visitor *CypherRenderer) {
	(*visitor).enter(o)
	for _, sortItem := range o.sortItems {
		o.PrepareVisit(sortItem).accept(visitor)
	}
	(*visitor).leave(o)
}

func (o Order) PrepareVisit(visitable Visitable) Visitable {
	return visitable
}

func (o Order) enter(renderer *CypherRenderer) {
	renderer.append(" ORDER BY ")
}

func (o Order) leave(renderer *CypherRenderer) {
}
