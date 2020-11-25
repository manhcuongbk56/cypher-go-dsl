package cypher_go_dsl

type SortItem struct {
	expression Expression
	direction  SortDirection
}

type SortDirection struct {
	value SortDirectionRaw
}

type SortDirectionRaw string

const (
	UNDEFINED SortDirectionRaw = ""
	ASC = "ASC"
	DESC = "DESC"
)

func CreateSortItem(expression Expression, direction SortDirectionRaw) SortItem{
	return SortItem{
		expression: expression,
		direction:  SortDirection{value: direction},
	}
}

func (item SortItem) Ascending() SortItem  {
	return SortItem{
		expression: item.expression,
		direction:  SortDirection{value: ASC},
	}
}

func (item SortItem) Descending() SortItem  {
	return SortItem{
		expression: item.expression,
		direction:  SortDirection{value: DESC},
	}
}

func (item SortItem) Accept(visitor *CypherRenderer) {
	(*visitor).Enter(item)
	NameOrExpression(item.expression).Accept(visitor)
	if item.direction.value == ASC || item.direction.value == DESC {
		item.direction.Accept(visitor)
	}
	(*visitor).Leave(item)
}

func (s SortDirection) Accept(visitor *CypherRenderer) {
	(*visitor).Enter(s)
	(*visitor).Leave(s)
}

func (item SortItem) GetType() VisitableType {
	return SortItemVisitable
}

func (s SortDirection) GetType() VisitableType {
	return SortDirectionVisitable
}

func (item SortItem) Enter(renderer *CypherRenderer) {
	panic("implement me")
}

func (item SortItem) Leave(renderer *CypherRenderer) {
	panic("implement me")
}

func (s SortDirection) Enter(renderer *CypherRenderer) {
	panic("implement me")
}

func (s SortDirection) Leave(renderer *CypherRenderer) {
	panic("implement me")
}






