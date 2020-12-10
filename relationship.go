package cypher_go_dsl

import "fmt"

type Relationship struct {
	left    *Node
	right   *Node
	details *RelationshipDetails
	key     string
	notNil  bool
	err     error
}

func RelationshipCreate(left Node, direction Direction, right Node, types ...string) Relationship {
	typeSlice := make([]string, 0)
	typeSlice = append(typeSlice, types...)
	relationshipTypes := RelationshipTypesCreate(typeSlice)
	details := RelationshipDetailsCreate1(direction, relationshipTypes)
	r := Relationship{
		left:    &left,
		right:   &right,
		details: &details,
		notNil:  true,
	}
	r.key = getAddress(&r)
	return r
}

func (r Relationship) getError() error {
	return r.err
}

func (r Relationship) isNotNil() bool {
	return r.notNil
}

func (r Relationship) getSymbolicName() SymbolicName {
	return r.details.symbolicName
}

func (r Relationship) getKey() string {
	return r.key
}

func (r Relationship) RelationshipTo(node Node, types ...string) RelationshipPattern {
	var relationshipI interface{} = (*(r.right)).RelationshipTo(node, types...)
	relationship, _ := relationshipI.(Relationship)
	return FirstElement(r).Add(relationship)
}

func (r Relationship) RelationshipFrom(node Node, types ...string) RelationshipPattern {
	panic("implement me")
}

func (r Relationship) RelationshipBetween(node Node, types ...string) RelationshipPattern {
	panic("implement me")
}

func (r Relationship) IsPatternElement() bool {
	return true
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
