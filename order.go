package cypher_go_dsl

import "fmt"

type Order struct {
	sortItems []SortItem
	key       string
	notNil    bool
	err error
}

func (o Order) getError() error {
	return o.err
}

func (o Order) isNotNil() bool {
	return o.notNil
}

func (o Order) getKey() string {
	return o.key
}

func (o Order) accept(visitor *CypherRenderer) {
	o.key = fmt.Sprint(&o)
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
	panic("implement me")
}

func (o Order) leave(renderer *CypherRenderer) {
	panic("implement me")
}
