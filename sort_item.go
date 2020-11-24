package cypher_go_dsl

type SortItem struct {
	expression IsExpression
	direction SortDirection
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

func CreateSortItem(expression IsExpression, direction SortDirectionRaw) SortItem{
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

func (item SortItem) Accept(visitor Visitor) {
	visitor.Enter(item)
	NameOrExpression(item.expression).Accept(visitor)
	if item.direction.value == ASC || item.direction.value == DESC {
		item.direction.Accept(visitor)
	}
	visitor.Leave(item)
}

func (s SortDirection) Accept(visitor Visitor) {
	visitor.Enter(s)
	visitor.Leave(s)
}

func (item SortItem) GetType() VisitableType {
	return SortItemVisitable
}

func (s SortDirection) GetType() VisitableType {
	return SortDirectionVisitable
}





