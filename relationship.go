package cypher

import "errors"

type Relationship struct {
	left    Node
	right   Node
	details RelationshipDetails
	key     string
	notNil  bool
	err     error
}

func RelationshipCreate(left Node, direction Direction, right Node, types ...string) Relationship {
	relationshipTypes := RelationshipTypes{}
	if len(types) > 0 {
		typeSlice := make([]string, 0)
		typeSlice = append(typeSlice, types...)
		relationshipTypes = RelationshipTypesCreate(typeSlice)
	}
	details := RelationshipDetailsCreate1(direction, relationshipTypes)
	return RelationshipCreate3(left, details, right)
}

func RelationshipCreate3(left Node, details RelationshipDetails, right Node) Relationship {
	if left.GetError() != nil {
		return RelationshipError(left.GetError())
	}
	if details.GetError() != nil {
		return RelationshipError(details.GetError())
	}
	if right.GetError() != nil {
		return RelationshipError(right.GetError())
	}
	r := Relationship{
		left:    left,
		right:   right,
		details: details,
		notNil:  true,
	}
	r.key = getAddress(&r)
	return r
}

func RelationshipError(err error) Relationship {
	return Relationship{
		err: err,
	}
}

func (r Relationship) GetError() error {
	return r.err
}

func (r Relationship) isNotNil() bool {
	return r.notNil
}

func (r Relationship) GetRequiredSymbolicName() SymbolicName {
	if r.details.symbolicName.isNotNil() {
		return r.details.symbolicName
	}
	return SymbolicNameError(errors.New("relationship get symbolic name:  no name present"))
}

func (r Relationship) GetSymbolicName() SymbolicName {
	return r.details.symbolicName
}

func (r Relationship) getKey() string {
	return r.key
}

func (r Relationship) NamedByString(newSymbolicName string) Relationship {
	return RelationshipCreate3(r.left, r.details.namedByString(newSymbolicName), r.right)
}

func (r Relationship) Unbounded() Relationship {
	return RelationshipCreate3(r.left, r.details.unbounded(), r.right)
}

func (r Relationship) Min(minimum int) Relationship {
	return RelationshipCreate3(r.left, r.details.min(minimum), r.right)
}

func (r Relationship) Length(minimum int, maximum int) Relationship {
	return RelationshipCreate3(r.left, r.details.min(minimum).max(maximum), r.right)
}

func (r Relationship) Max(maximum int) Relationship {
	return RelationshipCreate3(r.left, r.details.max(maximum), r.right)
}

func (r Relationship) accept(visitor *CypherRenderer) {
	visitor.enter(r)
	r.left.accept(visitor)
	r.details.accept(visitor)
	r.right.accept(visitor)
	visitor.leave(r)
}

func (r Relationship) enter(renderer *CypherRenderer) {
}

func (r Relationship) leave(renderer *CypherRenderer) {
}

func (r Relationship) IsPatternElement() bool {
	return true
}

func (r Relationship) RelationshipTo(node Node, types ...string) RelationshipChain {
	return RelationshipChainCreate(r).Add(r.right.RelationshipTo(node, types...))
}

func (r Relationship) RelationshipFrom(node Node, types ...string) RelationshipChain {
	return RelationshipChainCreate(r).Add(r.right.RelationshipFrom(node, types...))
}

func (r Relationship) RelationshipBetween(node Node, types ...string) RelationshipChain {
	return RelationshipChainCreate(r).Add(r.right.RelationshipBetween(node, types...))
}

func (r Relationship) Named(name string) Relationship {
	return RelationshipCreate3(r.left, r.details.namedByString(name), r.right)
}

func (r Relationship) NamedC(name string) RelationshipChain {
	return RelationshipChainError("can not use namedC for relationship")
}

func (r Relationship) Property(name string) Property {
	return PropertyCreate(r, name)
}

func (r Relationship) WithRawProperties(keysAndValues ...interface{}) Relationship {
	properties := MapExpression{}
	if keysAndValues != nil && len(keysAndValues) != 0 {
		properties = NewMapExpression(keysAndValues...)
		if properties.GetError() != nil {
			return RelationshipError(properties.GetError())
		}
	}
	return r.WithProperties(properties)
}

func (r Relationship) WithProperties(newProperties MapExpression) Relationship {
	if !newProperties.isNotNil() && !r.details.properties.isNotNil() {
		return r
	}
	property := Properties{}
	if newProperties.isNotNil() {
		property = PropertiesCreate(newProperties)
	}
	return RelationshipCreate3(r.left, r.details.with(property), r.right)
}

func (r Relationship) Project(entries ...interface{}) MapProjection {
	return MapProjectionCreate(r.GetRequiredSymbolicName(), entries...)
}
