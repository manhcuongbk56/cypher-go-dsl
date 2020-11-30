package cypher_go_dsl

import "fmt"

type Relationship struct {
	left    *Node
	right   *Node
	details *RelationshipDetails
	key     string
}

func (r Relationship) getSymbolicName() *SymbolicName {
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
	r.key = fmt.Sprint(&r)
	(*visitor).enter(r)
	r.left.accept(visitor)
	r.details.accept(visitor)
	r.right.accept(visitor)
	(*visitor).leave(r)
}

func (r Relationship) enter(renderer *CypherRenderer) {
}

func (r Relationship) leave(renderer *CypherRenderer) {
}
