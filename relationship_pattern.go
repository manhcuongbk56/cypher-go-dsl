package cypher_go_dsl

import (
	"fmt"
	"strings"
)

type RelationshipPattern interface {
	ExposesRelationship
	PatternElement
}

/**
RelationshipTypes
*/
type RelationshipTypes struct {
	values []string
	key    string
}

func (r RelationshipTypes) getKey() string {
	return r.key
}

func (r RelationshipTypes) accept(visitor *CypherRenderer) {
	r.key = fmt.Sprint(&r)
	(*visitor).enter(r)
	(*visitor).leave(r)
}

func (r RelationshipTypes) enter(renderer *CypherRenderer) {
	typeWithPrefix := make([]string, 0)
	for _, typeRaw := range r.values {
		if typeRaw == "" {
			continue
		}
		typeWithPrefix = append(typeWithPrefix, RelTypeStart+escapeName(typeRaw))
	}
	renderer.builder.WriteString(strings.Join(typeWithPrefix, RelTypSeparator))
}

func (r RelationshipTypes) leave(renderer *CypherRenderer) {
}

/**
Relationship length
*/

type RelationshipLength struct {
	minimum   *int8
	maximum   *int8
	unbounded bool
}

func (relationshipLength RelationshipLength) Unbounded() RelationshipLength {
	return RelationshipLength{nil, nil, true}
}

type Direction struct {
	symbolLeft  string
	symbolRight string
}

func LTR() Direction {
	return Direction{symbolLeft: "-", symbolRight: "->"}
}

func RTL() Direction {
	return Direction{symbolLeft: "<-", symbolRight: "-"}
}

func UNI() Direction {
	return Direction{symbolLeft: "-", symbolRight: "-"}
}

func CreateRelationship(left Node, direction Direction, right Node, types ...string) Relationship {
	typeSlice := make([]string, 0)
	typeSlice = append(typeSlice, types...)
	relationshipTypes := RelationshipTypes{values: typeSlice}
	details := RelationshipDetails{
		direction:    &direction,
		symbolicName: nil,
		types:        &relationshipTypes,
	}
	return Relationship{
		left:    &left,
		right:   &right,
		details: &details,
	}
}
