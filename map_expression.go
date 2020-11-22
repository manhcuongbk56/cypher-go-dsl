package cypher_go_dsl

import errors "golang.org/x/xerrors"

type MapExpression struct {
	Expressions []IsExpression
}

func NewMapExpression(objects ...interface{}) (MapExpression, error) {
	if len(objects) %2 != 0 {
		err := errors.Errorf("number of object input should be product of 2 but it is %d", len(objects))
		return MapExpression{}, err
	}
	var newContents = make([]IsExpression, len(objects)/2)
	var knownKeys = make(map[string]int)
	for i := 0; i < len(objects); i+=2 {
		key, isString := objects[i].(string)
		if !isString{
			err := errors.Errorf("key must be string")
			return MapExpression{}, err
		}
		value, isExpression := objects[i + 1].(IsExpression)
		if !isExpression{
			err := errors.Errorf("object must be expression")
			return MapExpression{}, err
		}
		if knownKeys[key] == 0  {
			err := errors.Errorf("duplicate key")
			return MapExpression{}, err
		}
		newContents = append(newContents, EntryExpression{
			Value: value,
			Key: key,
		})
	}
	return MapExpression{newContents}, nil
}

