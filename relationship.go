package cypher_go_dsl

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
	typeSlice := make([]string, 0)
	typeSlice = append(typeSlice, types...)
	relationshipTypes := RelationshipTypesCreate(typeSlice)
	details := RelationshipDetailsCreate1(direction, relationshipTypes)
	return RelationshipCreate3(left, details, right)
}

func RelationshipCreate3(left Node, details RelationshipDetails, right Node) Relationship {
	if left.getError() != nil {
		return RelationshipError(left.getError())
	}
	if details.getError() != nil {
		return RelationshipError(details.getError())
	}
	if right.getError() != nil {
		return RelationshipError(right.getError())
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

func (r Relationship) getError() error {
	return r.err
}

func (r Relationship) isNotNil() bool {
	return r.notNil
}

func (r Relationship) getRequiredSymbolicName() SymbolicName {
	if r.details.symbolicName.isNotNil() {
		return r.details.symbolicName
	}
	return SymbolicNameError(errors.New("no name present"))
}

func (r Relationship) getSymbolicName() SymbolicName {
	return r.details.symbolicName
}

func (r Relationship) getKey() string {
	return r.key
}

func (r Relationship) namedByString(newSymbolicName string) Relationship {
	return RelationshipCreate3(r.left, r.details.namedByString(newSymbolicName), r.right)
}

func (r Relationship) named(newSymbolicName SymbolicName) Relationship {
	return RelationshipCreate3(r.left, r.details.named(newSymbolicName), r.right)
}

func (r Relationship) unbounded() Relationship {
	return RelationshipCreate3(r.left, r.details.unbounded(), r.right)
}

func (r Relationship) min(minimum int) Relationship {
	return RelationshipCreate3(r.left, r.details.min(minimum), r.right)
}

func (r Relationship) max(minimum int, maximum int) Relationship {
	return RelationshipCreate3(r.left, r.details.min(minimum).max(maximum), r.right)
}

func (r Relationship) length(maximum int) Relationship {
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

func (r Relationship) RelationshipTo(node Node, types ...string) RelationshipPattern {
	return RelationshipChainCreate(r).Add(r.right.RelationshipTo(node, types...))
}

func (r Relationship) RelationshipFrom(node Node, types ...string) RelationshipPattern {
	return RelationshipChainCreate(r).Add(r.right.RelationshipFrom(node, types...))
}

func (r Relationship) RelationshipBetween(node Node, types ...string) RelationshipPattern {
	return RelationshipChainCreate(r).Add(r.right.RelationshipBetween(node, types...))
}

func (r Relationship) Property(name string) Property {
	return PropertyCreate(r, name)
}

func (r Relationship) WithRawProperties(keysAndValues ...interface{}) Relationship {
	properties := MapExpression{}
	if keysAndValues != nil && len(keysAndValues) != 0 {
		properties = NewMapExpression(keysAndValues...)
		if properties.getError() != nil {
			return RelationshipError(properties.getError())
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
