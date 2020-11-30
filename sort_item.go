package cypher_go_dsl

import "fmt"

type SortItem struct {
	expression Expression
	direction  SortDirection
	key        string
}



func (item SortItem) getKey() string {
	return item.key
}

type SortDirection struct {
	value SortDirectionRaw
	key   string
}

type SortDirectionRaw string

const (
	UNDEFINED SortDirectionRaw = ""
	ASC                        = "ASC"
	DESC                       = "DESC"
)

func CreateSortItem(expression Expression, direction SortDirectionRaw) SortItem {
	return SortItem{
		expression: expression,
		direction:  SortDirection{value: direction},
	}
}

func CreateAscendingSortItem(expression Expression) SortItem {
	return SortItem{
		expression: expression,
		direction: SortDirection{
			value: ASC,
		},
	}
}

func CreateDescendingSortItem(expression Expression) SortItem {
	return SortItem{
		expression: expression,
		direction: SortDirection{
			value: ASC,
		},
	}
}

func (item SortItem) Ascending() SortItem {
	return SortItem{
		expression: item.expression,
		direction:  SortDirection{value: ASC},
	}
}

func (item SortItem) Descending() SortItem {
	return SortItem{
		expression: item.expression,
		direction:  SortDirection{value: DESC},
	}
}

func (item SortItem) accept(visitor *CypherRenderer) {
	item.key = fmt.Sprint(&item)
	(*visitor).enter(item)
	NameOrExpression(item.expression).accept(visitor)
	if item.direction.value == ASC || item.direction.value == DESC {
		item.direction.accept(visitor)
	}
	(*visitor).leave(item)
}

func (s SortDirection) accept(visitor *CypherRenderer) {
	(*visitor).enter(s)
	(*visitor).leave(s)
}

func (s SortDirection) getKey() string {
	return s.key
}

func (item SortItem) enter(renderer *CypherRenderer) {
	panic("implement me")
}

func (item SortItem) leave(renderer *CypherRenderer) {
	panic("implement me")
}

func (s SortDirection) enter(renderer *CypherRenderer) {
	panic("implement me")
}

func (s SortDirection) leave(renderer *CypherRenderer) {
	panic("implement me")
}
