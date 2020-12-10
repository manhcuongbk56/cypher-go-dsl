package cypher_go_dsl

import (
	"fmt"
	"strconv"
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
	notNil bool
	err error
}

func RelationshipTypesCreate(types []string) RelationshipTypes {
	return RelationshipTypes{
		values: types,
		notNil: true,
	}
}

func (r RelationshipTypes) getError() error {
	return r.err
}

func (r RelationshipTypes) isNotNil() bool {
	return r.notNil
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
	minimum   *int
	maximum   *int
	unbounded bool
	key       string
	notNil    bool
	err error
}

func (relationshipLength RelationshipLength) getError() error {
	return relationshipLength.err
}

func (relationshipLength RelationshipLength) accept(visitor *CypherRenderer) {
	relationshipLength.key = fmt.Sprint(&relationshipLength)
	visitor.enter(relationshipLength)
	visitor.leave(relationshipLength)
}

func (relationshipLength RelationshipLength) enter(renderer *CypherRenderer) {
	minimum := relationshipLength.minimum
	maximum := relationshipLength.maximum
	if relationshipLength.unbounded {
		renderer.builder.WriteString("*")
		return
	}
	if minimum == nil && maximum == nil {
		return
	}
	renderer.builder.WriteString("*")
	if minimum != nil {
		renderer.builder.WriteString(strconv.Itoa(*minimum))
	}
	renderer.builder.WriteString("..")
	if maximum != nil {
		renderer.builder.WriteString(strconv.Itoa(*maximum))
	}
}

func (relationshipLength RelationshipLength) leave(renderer *CypherRenderer) {
}

func (relationshipLength RelationshipLength) getKey() string {
	return relationshipLength.key
}

func (relationshipLength RelationshipLength) isNotNil() bool {
	return relationshipLength.notNil
}

func (relationshipLength RelationshipLength) Unbounded() RelationshipLength {
	return RelationshipLength{unbounded: true}
}

type Direction struct {
	symbolLeft  string
	symbolRight string
	notNil      bool
}

func DirectionCreate(symbolLeft string, symbolRight string) Direction {
	return Direction{symbolLeft, symbolRight, true}
}

func LTR() Direction {
	return DirectionCreate("-", "->")
}

func RTL() Direction {
	return DirectionCreate("<-", "-")
}

func UNI() Direction {
	return DirectionCreate("-", "-")
}

func CreateRelationship(left Node, direction Direction, right Node, types ...string) Relationship {
	typeSlice := make([]string, 0)
	typeSlice = append(typeSlice, types...)
	relationshipTypes := RelationshipTypesCreate(typeSlice)
	details := RelationshipDetailsCreate1(direction, relationshipTypes)
	return Relationship{
		left:    &left,
		right:   &right,
		details: &details,
		notNil:  true,
	}
}
