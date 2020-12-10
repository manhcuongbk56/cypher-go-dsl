package cypher_go_dsl

import (
	"fmt"
	errors "golang.org/x/xerrors"
)

type MapExpression struct {
	expressions []Expression
	key         string
	notNil      bool
	err         error
}

func MapExpressionCreate(newContents []Expression) MapExpression {
	m := MapExpression{
		expressions: newContents,
		notNil:      true,
	}
	m.key = getAddress(&m)
	return m
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

func NewMapExpression(objects ...interface{}) (MapExpression, error) {
	if len(objects)%2 != 0 {
		err := errors.Errorf("number of object input should be product of 2 but it is %defaultBuilder", len(objects))
		return MapExpression{}, err
	}
	var newContents = make([]Expression, len(objects)/2)
	var knownKeys = make(map[string]int)
	for i := 0; i < len(objects); i += 2 {
		key, isString := objects[i].(string)
		if !isString {
			err := errors.Errorf("key must be string")
			return MapExpression{}, err
		}
		value, isExpression := objects[i+1].(Expression)
		if !isExpression {
			err := errors.Errorf("object must be expression")
			return MapExpression{}, err
		}
		if knownKeys[key] == 1 {
			err := errors.Errorf("duplicate key")
			return MapExpression{}, err
		}
		knownKeys[key] = 1
		newContents[i/2] = EntryExpressionCreate(key, value)
	}
	return MapExpressionCreate(newContents), nil
}

func (m MapExpression) accept(visitor *CypherRenderer) {
	(*visitor).enter(m)
	for _, child := range m.expressions {
		m.PrepareVisit(child).accept(visitor)
	}
	(*visitor).leave(m)
}

func (m MapExpression) enter(renderer *CypherRenderer) {
	renderer.builder.WriteString("{")
}

func (m MapExpression) leave(renderer *CypherRenderer) {
}

func (m MapExpression) PrepareVisit(visitable Visitable) Visitable {
	expression := visitable.(Expression)
	return NameOrExpression(expression)
}
