package cypher_go_dsl

import errors "golang.org/x/xerrors"

type MapExpression struct {
	expressions []Expression
}

func NewMapExpression(objects ...interface{}) (MapExpression, error) {
	if len(objects) %2 != 0 {
		err := errors.Errorf("number of object input should be product of 2 but it is %defaultStatementBuilder", len(objects))
		return MapExpression{}, err
	}
	var newContents = make([]Expression, len(objects)/2)
	var knownKeys = make(map[string]int)
	for i := 0; i < len(objects); i+=2 {
		key, isString := objects[i].(string)
		if !isString{
			err := errors.Errorf("key must be string")
			return MapExpression{}, err
		}
		value, isExpression := objects[i + 1].(Expression)
		if !isExpression{
			err := errors.Errorf("object must be expression")
			return MapExpression{}, err
		}
		if knownKeys[key] == 1  {
			err := errors.Errorf("duplicate key")
			return MapExpression{}, err
		}
		knownKeys[key] = 1
		newContents[i/2] = EntryExpression{
			Value: value,
			Key: key,
		}
	}
	return MapExpression{newContents}, nil
}

func (m MapExpression) Accept(visitor *CypherRenderer) {
	(*visitor).Enter(m)
	for _, child := range m.expressions {
		m.PrepareVisit(child).Accept(visitor)
	}
	(*visitor).Leave(m)
}

func (m MapExpression) GetType() VisitableType {
	return MapExpressionVisitable
}

func (m MapExpression) Enter(renderer *CypherRenderer) {
	renderer.builder.WriteString("{")
}

func (m MapExpression) Leave(renderer *CypherRenderer) {
}

func (m MapExpression) PrepareVisit(visitable Visitable) Visitable {
	expression := visitable.(Expression)
	return NameOrExpression(expression)
}


