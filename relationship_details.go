package cypher_go_dsl

import "fmt"

type RelationshipDetails struct {
	direction    Direction
	symbolicName SymbolicName
	types        RelationshipTypes
	length       RelationshipLength
	properties   Properties
	key          string
	notNil       bool
}

func RelationshipDetailsCreate1(direction Direction, types RelationshipTypes) RelationshipDetails {
	return RelationshipDetails{
		direction: direction,
		types:     types,
		notNil:    true,
	}
}

func (r RelationshipDetails) isNotNil() bool {
	return r.notNil
}

func (r RelationshipDetails) getKey() string {
	return r.key
}

func RelationshipDetailsCreate2(direction Direction, symbolicName SymbolicName, types RelationshipTypes) RelationshipDetails {
	return RelationshipDetails{
		direction:    direction,
		symbolicName: symbolicName,
		types:        types,
		notNil:       true,
	}
}

func RelationshipDetailsCreate(direction Direction, symbolicName SymbolicName,
	types RelationshipTypes, length RelationshipLength, properties Properties) RelationshipDetails {
	return RelationshipDetails{
		direction:    direction,
		symbolicName: symbolicName,
		types:        types,
		length:       length,
		properties:   properties,
		notNil:       true,
	}
}

func (r RelationshipDetails) hasContent() bool {
	return r.direction.notNil ||
		r.symbolicName.isNotNil() ||
		r.types.isNotNil() ||
		r.length.isNotNil() ||
		r.properties.isNotNil()
}

func (r RelationshipDetails) accept(visitor *CypherRenderer) {
	r.key = fmt.Sprint(&r)
	visitor.enter(r)
	VisitIfNotNull(r.symbolicName, visitor)
	VisitIfNotNull(r.types, visitor)
	VisitIfNotNull(r.length, visitor)
	VisitIfNotNull(r.properties, visitor)
	visitor.leave(r)
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
