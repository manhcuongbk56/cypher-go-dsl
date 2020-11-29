package cypher_go_dsl

import "fmt"

type RelationshipDetails struct {
	direction    *Direction
	symbolicName *SymbolicName
	types        *RelationshipTypes
	length       *RelationshipLength
	properties   *Properties
	key          string
}

func (r RelationshipDetails) getKey() string {
	return r.key
}

func CreateRelationshipDetail(direction Direction, symbolicName SymbolicName, types RelationshipTypes) RelationshipDetails {
	return RelationshipDetails{
		direction:    &direction,
		symbolicName: &symbolicName,
		types:        &types,
	}
}

func CreateRelationshipDetailFull(direction Direction, symbolicName SymbolicName,
	types RelationshipTypes, length RelationshipLength, properties Properties) RelationshipDetails {
	return RelationshipDetails{
		direction:    &direction,
		symbolicName: &symbolicName,
		types:        &types,
		length:       &length,
		properties:   &properties,
	}
}

func (r RelationshipDetails) hasContent() bool {
	return r.direction != nil ||
		r.symbolicName != nil ||
		r.types != nil ||
		r.length != nil ||
		r.properties != nil
}

func (r RelationshipDetails) accept(visitor *CypherRenderer) {
	r.key = fmt.Sprint(&r)
	(*visitor).enter(r)
	VisitIfNotNull(r.symbolicName, visitor)
	VisitIfNotNull(r.types, visitor)
	VisitIfNotNull(r.length, visitor)
	VisitIfNotNull(r.properties, visitor)
	(*visitor).Leave(r)
}

func (r RelationshipDetails) enter(renderer *CypherRenderer) {
	direction := r.direction
	renderer.builder.WriteString(direction.symbolLeft)
	if r.hasContent() {
		renderer.builder.WriteString("[")
	}
}

func (r RelationshipDetails) leave(renderer *CypherRenderer) {
	direction := r.direction
	if r.hasContent() {
		renderer.builder.WriteString("]")
	}
	renderer.builder.WriteString(direction.symbolRight)
}
