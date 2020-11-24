package cypher_go_dsl

type OrderBuilder struct {
	sortItemList []SortItem
	lastSortItem *SortItem
	skip Skip
	limit Limit
}

func (o OrderBuilder) OrderByItem(item ...SortItem)  {
	o.sortItemList = append(o.sortItemList, item...)
}

func (o OrderBuilder) OrderByExpression(expression Expression)  {
	o.lastSortItem = Sort(expression)
}

func (o OrderBuilder) And(expression Expression)  {
	o.OrderByExpression(expression)
}

func (o OrderBuilder) Descending()  {
	o.sortItemList = append(o.sortItemList, o.lastSortItem.Descending())
	o.lastSortItem = nil
}

func (o OrderBuilder) Ascending()  {
	o.sortItemList = append(o.sortItemList, o.lastSortItem.Ascending())
	o.lastSortItem = nil
}

func (o OrderBuilder) Skip(number int)  {
	o.skip = CreateSkip(number)
}

func (o OrderBuilder) Limit(number int)  {
	o.limit = CreateLimit(number)
}

func (o OrderBuilder) BuildOrder() *Order {
	if o.lastSortItem != nil {
		o.sortItemList = append(o.sortItemList, *o.lastSortItem)
	}
	if o.sortItemList != nil && len(o.sortItemList) > 0 {
		return &Order{o.sortItemList}
	}
	return nil
}

