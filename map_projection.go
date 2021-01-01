package cypher

import (
	"golang.org/x/xerrors"
	"reflect"
)

type MapProjection struct {
	ExpressionContainer
	name          SymbolicName
	mapExpression MapExpression
	key           string
	err           error
	notNil        bool
}

func MapProjectionCreate(name SymbolicName, content ...interface{}) MapProjection {
	if name.GetError() != nil {
		return MapProjectionError(name.GetError())
	}
	newContent, err := createNewContent(content...)
	if err != nil {
		return MapProjectionError(err)
	}
	mapProjection := MapProjection{
		name:          name,
		mapExpression: MapExpressionCreate(newContent),
		notNil:        true,
	}
	mapProjection.key = getAddress(&mapProjection)
	mapProjection.ExpressionContainer = ExpressionWrap(mapProjection)
	return mapProjection
}

func MapProjectionCreateWithMapExpression(name SymbolicName, mapExpression MapExpression) MapProjection {
	if name.GetError() != nil {
		return MapProjectionError(name.GetError())
	}
	if mapExpression.GetError() != nil {
		return MapProjectionError(mapExpression.GetError())
	}
	mapProjection := MapProjection{
		name:          name,
		mapExpression: mapExpression,
		notNil:        true,
	}
	mapProjection.key = getAddress(&mapProjection)
	mapProjection.ExpressionContainer = ExpressionWrap(mapProjection)
	return mapProjection
}

func MapProjectionError(err error) MapProjection {
	return MapProjection{err: err}
}

func (m MapProjection) GetError() error {
	return m.err
}

func (m MapProjection) accept(visitor *CypherRenderer) {
	visitor.enter(m)
	m.name.accept(visitor)
	m.mapExpression.accept(visitor)
	visitor.leave(m)
}

func (m MapProjection) enter(_ *CypherRenderer) {
}

func (m MapProjection) leave(_ *CypherRenderer) {
}

func (m MapProjection) getKey() string {
	return m.key
}

func (m MapProjection) isNotNil() bool {
	return m.notNil
}

func (m MapProjection) GetExpressionType() ExpressionType {
	return "MapProjection"
}

func (m MapProjection) And(content ...interface{}) MapProjection {
	if m.GetError() != nil {
		return m
	}
	newContent, err := createNewContent(content...)
	if err != nil {
		return MapProjectionError(err)
	}
	return MapProjectionCreateWithMapExpression(m.name, m.mapExpression.AddEntries(newContent))
}

func contentAt(content []interface{}, i int) interface{} {
	currentObject := content[i]
	if expression, isExpression := currentObject.(Expression); isExpression {
		return NameOrExpression(expression)
	} else if named, isNamed := currentObject.(Named); isNamed {
		symbolicName := named.getSymbolicName()
		if symbolicName.isNotNil() {
			return symbolicName
		}
	}
	return currentObject
}

func createNewContent(content ...interface{}) ([]Expression, error) {
	newContent := make([]Expression, 0)
	knownKeys := make(map[string]int)
	lastKey := ""
	var lastExpression Expression
	i := 0
	for i < len(content) {
		var next interface{}
		if i+1 >= len(content) {
			next = nil
		} else {
			next = contentAt(content, i+1)
		}
		current := contentAt(content, i)
		if aString, isString := current.(string); isString {
			if anExpression, isExpression := next.(Expression); isExpression {
				lastKey = aString
				lastExpression = anExpression
				i += 2
			} else {
				lastKey = ""
				lastExpression = PropertyLookupCreate(aString)
				i += 1
			}
		} else if anExpression, isExpression := current.(Expression); isExpression {
			lastKey = ""
			lastExpression = anExpression
			i += 1
		}

		if _, isAsterisk := lastExpression.(Asterisk); isAsterisk {
			lastExpression = PropertyLookupCreate("*")
		}
		var entry Expression
		if lastKey != "" {
			_, isDuplicateKey := knownKeys[lastKey]
			if isDuplicateKey {
				return nil, xerrors.Errorf("map projection create new content: duplicate key %s", lastKey)
			}
			entry = KeyValueMapEntryCreate(lastKey, lastExpression)
			knownKeys[lastKey] = 1
		} else if _, isSymbolicName := lastExpression.(SymbolicName); isSymbolicName {
			entry = lastExpression
		} else if _, isPropertyLookup := lastExpression.(PropertyLookup); isPropertyLookup {
			entry = lastExpression
		} else if aProperty, isProperty := lastExpression.(Property); isProperty {
			entry = aProperty.getName()
		} else if anAliasedExpression, isAliasedExpression := lastExpression.(AliasedExpression); isAliasedExpression {
			entry = KeyValueMapEntryCreate(anAliasedExpression.GetAlias(), anAliasedExpression.delegate)
		} else if lastExpression == nil {
			return nil, xerrors.Errorf("map projection create new content: could not determine an expression from the given content!")
		} else {
			return nil, xerrors.Errorf("map projection create new content: unknown type %s" +
				" cannot be used with an implicit name as map entry", reflect.TypeOf(lastExpression))
		}
		if entry != nil && entry.GetError() != nil {
			return nil, entry.GetError()
		}
		newContent = append(newContent, entry)
		lastKey = ""
		lastExpression = nil
	}
	return newContent, nil
}
