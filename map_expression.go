package cypher

import (
	errors "golang.org/x/xerrors"
)

type MapExpression struct {
	ExpressionContainer
	expressions []Expression
	key         string
	notNil      bool
	err         error
}

func MapExpressionCreate(newContents []Expression) MapExpression {
	for _, content := range newContents {
		if content != nil && content.getError() != nil {
			return MapExpressionError(content.getError())
		}
	}
	m := MapExpression{
		expressions: newContents,
		notNil:      true,
	}
	m.key = getAddress(&m)
	m.ExpressionContainer = ExpressionWrap(m)
	return m
}

func MapExpressionError(err error) MapExpression {
	return MapExpression{
		err: err,
	}
}

func (m MapExpression) GetExpressionType() ExpressionType {
	return "MapExpression"
}

func (m MapExpression) getError() error {
	return m.err
}

func (m MapExpression) isNotNil() bool {
	return m.notNil
}

func (m MapExpression) getKey() string {
	return m.key
}

func NewMapExpression(objects ...interface{}) MapExpression {
	if len(objects)%2 != 0 {
		return MapExpressionError(errors.Errorf("number of object input should be product of 2 but it is %defaultBuilder", len(objects)))
	}
	var newContents = make([]Expression, len(objects)/2)
	var knownKeys = make(map[string]int)
	for i := 0; i < len(objects); i += 2 {
		key, isString := objects[i].(string)
		if !isString {
			return MapExpressionError(errors.Errorf("key must be string"))
		}
		value, isExpression := objects[i+1].(Expression)
		if !isExpression {
			return MapExpressionError(errors.Errorf("object must be expression"))
		}
		if knownKeys[key] == 1 {
			return MapExpressionError(errors.Errorf("duplicate key"))
		}
		knownKeys[key] = 1
		newContents[i/2] = EntryExpressionCreate(key, value)
	}
	return MapExpressionCreate(newContents)
}

func (m MapExpression) accept(visitor *CypherRenderer) {
	(*visitor).enter(m)
	for _, child := range m.expressions {
		m.PrepareVisit(child).accept(visitor)
	}
	(*visitor).leave(m)
}

func (m MapExpression) enter(renderer *CypherRenderer) {
	renderer.append("{")
}

func (m MapExpression) leave(renderer *CypherRenderer) {
	renderer.append("}")
}

func (m MapExpression) PrepareVisit(visitable Visitable) Visitable {
	expression := visitable.(Expression)
	return NameOrExpression(expression)
}
