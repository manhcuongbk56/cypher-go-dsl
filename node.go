package cypher_go_dsl

import (
	"fmt"
)

type Node struct {
	symbolicName *SymbolicName
	labels       []NodeLabel
	properties   *Properties
}

func (node Node) hasSymbolic() bool {
	return node.symbolicName != nil
}

func (node Node) Accept(visitor *CypherRenderer) {
	(*visitor).Enter(&node)
	VisitIfNotNull(node.symbolicName, visitor)
	for _, label := range node.labels {
		label.Accept(visitor)
	}
	VisitIfNotNull(node.properties, visitor)
	(*visitor).Leave(&node)
}

func (node Node) GetType() VisitableType {
	return NodeVisitable
}

func (n NodeLabel) GetType() VisitableType {
	return NodeLabelVisitable
}

func (node Node) RelationshipTo(nodeDest Node, types ...string) Relationship {
	panic("implement me")
}

func (node Node) RelationshipFrom(other Node, types ...string) Relationship {
	return CreateLTR(node, RTL(), other, types...)
}

func (node Node) RelationshipBetween(nodeDest Node, types ...string) Relationship {
	panic("implement me")
}

func (node Node) WithRawProperties(keysAndValues ...interface{}) (Node, error){
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
	return Node{symbolicName: node.symbolicName, labels: node.labels, properties: &Properties{newProperties}}
}

func (node Node) Property(name string) {
	fmt.Print(name)
}

func (node Node) Named(name string) Node {
	node.symbolicName = &SymbolicName{Value: name}
	return node
}

func (node Node) getSymbolicName() *SymbolicName {
	return node.symbolicName
}

func (node Node) Enter(renderer *CypherRenderer) {
	renderer.builder.WriteString("(")
	if !node.hasSymbolic(){
		return
	}
	var named interface{} = &node
	_, renderer.skipNodeContent = renderer.visitedNamed[named]
	if renderer.skipNodeContent {
		renderer.builder.WriteString(node.symbolicName.Value)
	}
}

func (node Node) Leave(renderer *CypherRenderer) {
	panic("implement me")
}

type NodeLabel struct {
	value string
}

func (n NodeLabel) Accept(visitor *CypherRenderer) {
	(*visitor).Enter(n)
	(*visitor).Leave(n)
}

func (n NodeLabel) Enter(renderer *CypherRenderer) {
	if n.value == "" {
		return
	}
	renderer.builder.WriteString(NODE_LABEL_START)
	renderer.builder.WriteString(escapeName(n.value))
}

func (n NodeLabel) Leave(renderer *CypherRenderer) {
}







