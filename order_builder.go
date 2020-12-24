package cypher

type OrderBuilder struct {
	sortItemList []SortItem
	lastSortItem SortItem
	skip         Skip
	limit        Limit
	notNil       bool
}

func OrderBuilderCreate() OrderBuilder {
	return OrderBuilder{
		sortItemList: make([]SortItem, 0),
	}
}

func (o *OrderBuilder) OrderBySortItem(item ...SortItem) {
	o.sortItemList = append(o.sortItemList, item...)
}

func (o *OrderBuilder) OrderByExpression(expression Expression) {
	o.lastSortItem = CypherSort(expression)
}

func (o *OrderBuilder) And(expression Expression) {
	o.OrderByExpression(expression)
}

func (o *OrderBuilder) Descending() {
	o.sortItemList = append(o.sortItemList, o.lastSortItem.Descending())
	o.lastSortItem = SortItem{}
}

func (o *OrderBuilder) Ascending() {
	o.sortItemList = append(o.sortItemList, o.lastSortItem.Ascending())
	o.lastSortItem = SortItem{}
}

func (o *OrderBuilder) Skip(number int) {
	o.skip = SkipCreate(number)
}

func (o *OrderBuilder) Limit(number int) {
	o.limit = LimitCreate(number)
}

func (o *OrderBuilder) BuildOrder() Order {
	if o.lastSortItem.isNotNil() {
		o.sortItemList = append(o.sortItemList, o.lastSortItem)
	}
	if o.sortItemList != nil && len(o.sortItemList) > 0 {
		return OrderCreate(o.sortItemList)
	}
	return Order{}
}

func (o *OrderBuilder) reset() {
	o.sortItemList = nil
	o.lastSortItem = SortItem{}
	o.skip = Skip{}
	o.limit = Limit{}
}
