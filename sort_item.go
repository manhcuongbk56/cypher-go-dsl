package cypher

type SortItem struct {
	expression Expression
	direction  SortDirection
	key        string
	notNil     bool
	err        error
}

func SortItemCreate(expression Expression, direction SortDirectionRaw) SortItem {
	if expression != nil && expression.GetError() != nil {
		return SortItemError(expression.GetError())
	}
	sortItem := SortItem{
		expression: expression,
		direction:  SortDirection{value: direction},
		notNil:     true,
	}
	sortItem.key = getAddress(&sortItem)
	return sortItem
}

func SortItemError(err error) SortItem {
	return SortItem{err: err}
}

func (item SortItem) GetError() error {
	return item.err
}

func (item SortItem) isNotNil() bool {
	return item.notNil
}

func (item SortItem) getKey() string {
	return item.key
}

func (item SortItem) Ascending() SortItem {
	return SortItemCreate(item.expression, ASC)
}

func (item SortItem) Descending() SortItem {
	return SortItem{
		expression: item.expression,
		direction:  SortDirection{value: DESC},
	}
}

func (item SortItem) accept(visitor *CypherRenderer) {
	visitor.enter(item)
	NameOrExpression(item.expression).accept(visitor)
	if item.direction.value == ASC || item.direction.value == DESC {
		item.direction.accept(visitor)
	}
	visitor.leave(item)
}

func (item SortItem) enter(renderer *CypherRenderer) {
}

func (item SortItem) leave(renderer *CypherRenderer) {
}

//SORT DIRECTION

type SortDirection struct {
	value  SortDirectionRaw
	key    string
	notNil bool
	err    error
}

func (s SortDirection) GetError() error {
	return s.err
}

func (s SortDirection) isNotNil() bool {
	return s.notNil
}

type SortDirectionRaw string

const (
	UNDEFINED SortDirectionRaw = ""
	ASC                        = "ASC"
	DESC                       = "DESC"
)

func CreateAscendingSortItem(expression Expression) SortItem {
	return SortItemCreate(expression, ASC)
}

func CreateDescendingSortItem(expression Expression) SortItem {
	return SortItemCreate(expression, DESC)
}

func (s SortDirection) accept(visitor *CypherRenderer) {
	visitor.enter(s)
	visitor.leave(s)
}

func (s SortDirection) getKey() string {
	return s.key
}

func (s SortDirection) enter(renderer *CypherRenderer) {
	renderer.append(" ").append(string(s.value))
}

func (s SortDirection) leave(renderer *CypherRenderer) {
}
