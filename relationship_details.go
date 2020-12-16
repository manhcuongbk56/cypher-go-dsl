package cypher_go_dsl

import "errors"

type RelationshipDetails struct {
	direction    Direction
	symbolicName SymbolicName
	types        RelationshipTypes
	length       RelationshipLength
	properties   Properties
	key          string
	notNil       bool
	err          error
}

func RelationshipDetailsCreate1(direction Direction, types RelationshipTypes) RelationshipDetails {
	if types.getError() != nil {
		return RelationshipDetailsError(types.getError())
	}
	r := RelationshipDetails{
		direction: direction,
		types:     types,
		notNil:    true,
	}
	r.key = getAddress(&r)
	return r
}

func RelationshipDetailsCreate2(direction Direction, symbolicName SymbolicName, types RelationshipTypes) RelationshipDetails {
	if symbolicName.getError() != nil {
		return RelationshipDetailsError(symbolicName.getError())
	}
	if types.getError() != nil {
		return RelationshipDetailsError(types.getError())
	}
	r := RelationshipDetails{
		direction:    direction,
		symbolicName: symbolicName,
		types:        types,
		notNil:       true,
	}
	r.key = getAddress(&r)
	return r
}

func RelationshipDetailsCreate(direction Direction, symbolicName SymbolicName,
	types RelationshipTypes, length RelationshipLength, properties Properties) RelationshipDetails {
	if symbolicName.getError() != nil {
		return RelationshipDetailsError(symbolicName.getError())
	}
	if types.getError() != nil {
		return RelationshipDetailsError(types.getError())
	}
	if length.getError() != nil {
		return RelationshipDetailsError(length.getError())
	}
	if properties.getError() != nil {
		return RelationshipDetailsError(properties.getError())
	}
	return RelationshipDetails{
		direction:    direction,
		symbolicName: symbolicName,
		types:        types,
		length:       length,
		properties:   properties,
		notNil:       true,
	}
}

func RelationshipDetailsError(err error) RelationshipDetails {
	return RelationshipDetails{
		err: err,
	}
}

func (r RelationshipDetails) getError() error {
	return r.err
}

func (r RelationshipDetails) isNotNil() bool {
	return r.notNil
}

func (r RelationshipDetails) getKey() string {
	return r.key
}

func (r RelationshipDetails) namedByString(newNamed string) RelationshipDetails {
	if newNamed == "" {
		return RelationshipDetailsError(errors.New("symbolic name is required"))
	}
	return r.named(SymbolicNameCreate(newNamed))
}

func (r RelationshipDetails) named(newSymbolicName SymbolicName) RelationshipDetails {
	if !newSymbolicName.isNotNil() {
		return RelationshipDetailsError(errors.New("symbolic name is required"))
	}
	return RelationshipDetailsCreate(r.direction, newSymbolicName, r.types, r.length, r.properties)
}

func (r RelationshipDetails) with(newProperties Properties) RelationshipDetails {
	return RelationshipDetailsCreate(r.direction, r.symbolicName, r.types, r.length, newProperties)
}

func (r RelationshipDetails) unbounded() RelationshipDetails {
	return RelationshipDetailsCreate(r.direction, r.symbolicName, r.types, RelationshipLengthCreate(true), r.properties)
}

func (r RelationshipDetails) min(minimum int) RelationshipDetails {
	if !r.length.isNotNil() || r.length.minimum == nil {
		return r
	}
	newLength := RelationshipLengthCreate1(&minimum, nil)
	if r.length.isNotNil() {
		newLength = RelationshipLengthCreate1(&minimum, r.length.maximum)
	}
	return RelationshipDetailsCreate(r.direction, r.symbolicName, r.types, newLength, r.properties)
}

func (r RelationshipDetails) max(maximum int) RelationshipDetails {
	if !r.length.isNotNil() || r.length.minimum == nil {
		return r
	}
	newLength := RelationshipLengthCreate1(nil, &maximum)
	if r.length.isNotNil() {
		newLength = RelationshipLengthCreate1(r.length.minimum, &maximum)
	}
	return RelationshipDetailsCreate(r.direction, r.symbolicName, r.types, newLength, r.properties)
}

func (r RelationshipDetails) hasContent() bool {
	return r.direction.notNil ||
		r.symbolicName.isNotNil() ||
		r.types.isNotNil() ||
		r.length.isNotNil() ||
		r.properties.isNotNil()
}

func (r RelationshipDetails) accept(visitor *CypherRenderer) {
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
