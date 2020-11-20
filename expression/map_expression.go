package expression

import errors "golang.org/x/xerrors"

type MapExpression struct {
	Expressions []Expression
}

func NewMapExpression(objects ...interface{}) (MapExpression, error) {
	if len(objects) %2 != 0 {
		err := errors.Errorf("number of object input should be product of 2 but it is %d", len(objects))
		return MapExpression{}, err
	}
	var knownKeys = make(map[string]int)
	for i := 0; i < len(objects); i+=2 {
		key, isString := objects[i].(string)
		if !isString{
			err := errors.Errorf("key must be string")
			return MapExpression{}, err
		}
		value, isExpression := objects[i + 1].(IExpression)
		if !isExpression{
			err := errors.Errorf("object must be expression")
			return MapExpression{}, err
		}
		if knownKeys[key] {

		}
	}
	return MapExpression{}, err
}

func
