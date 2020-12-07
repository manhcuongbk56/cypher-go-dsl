package cypher_go_dsl

import (
	"fmt"
)

type Node struct {
	symbolicName SymbolicName
	labels       []NodeLabel
	properties   Properties
	key          string
	notNil       bool
}

func NodeCreate() Node {
	return Node{
		notNil: true,
	}
}

func NodeCreate1(symbolicName string) Node {
	node := Node{
		notNil: true,
	}
	return node.Named(symbolicName)
}

func (node Node) getSymbolicName() SymbolicName {
	return node.symbolicName
}

func (node Node) isNotNil() bool {
	return node.notNil
}

func (node Node) IsPatternElement() bool {
	return true
}

func (node Node) getKey() string {
	return node.key
}

func (node Node) hasSymbolic() bool {
	return node.symbolicName.isNotNil()
}

func (node Node) accept(visitor *CypherRenderer) {
	node.key = fmt.Sprint(&node)
	(*visitor).enter(node)
	VisitIfNotNull(node.symbolicName, visitor)
	for _, label := range node.labels {
		label.accept(visitor)
	}
	VisitIfNotNull(node.properties, visitor)
	(*visitor).leave(node)
}

func (node Node) RelationshipTo(other Node, types ...string) Relationship {
	return CreateRelationship(node, LTR(), other, types...)
}

func (node Node) RelationshipFrom(other Node, types ...string) Relationship {
	return CreateRelationship(node, RTL(), other, types...)
}

func (node Node) RelationshipBetween(nodeDest Node, types ...string) Relationship {
	panic("implement me")
}

func (node Node) WithRawProperties(keysAndValues ...interface{}) (Node, error) {
	var properties = &MapExpression{}
	if keysAndValues != nil && len(keysAndValues) != 0 {
		var err error
		*properties, err = NewMapExpression(keysAndValues...)
		if err != nil {
			return Node{}, err
		}
	}
	return node.WithProperties(*properties), nil
}

func (node Node) WithProperties(newProperties MapExpression) Node {
	return Node{symbolicName: node.symbolicName, labels: node.labels, notNil: true, properties: Properties{properties: newProperties, notNil: true}}
}

func (node Node) Property(name string) {
	fmt.Print(name)
}

func (node Node) Named(name string) Node {
	node.symbolicName = SymbolicNameCreate(name)
	return node
}

func (node Node) enter(renderer *CypherRenderer) {
	renderer.builder.WriteString("(")
	if !node.hasSymbolic() {
		return
	}
	var named interface{} = &node
	_, renderer.skipNodeContent = renderer.visitedNamed[named]
	if renderer.skipNodeContent {
		renderer.builder.WriteString(node.symbolicName.value)
	}
}

func (node Node) leave(renderer *CypherRenderer) {
	renderer.builder.WriteString(")")
	renderer.skipNodeContent = false
}

type NodeLabel struct {
	value  string
	key    string
	notNil bool
}

func NodeLabelCreate(value string) NodeLabel {
	return NodeLabel{
		value:  value,
		notNil: true,
	}
}

func (n NodeLabel) isNotNil() bool {
	return n.notNil
}

func (n NodeLabel) getKey() string {
	return n.key
}

func (n NodeLabel) accept(visitor *CypherRenderer) {
	n.key = fmt.Sprint(&n)
	visitor.enter(n)
	visitor.leave(n)
}

func (n NodeLabel) enter(renderer *CypherRenderer) {
	if n.value == "" {
		return
	}
	renderer.builder.WriteString(NodeLabelStart)
	renderer.builder.WriteString(escapeName(n.value))
}

func (n NodeLabel) leave(renderer *CypherRenderer) {
}
