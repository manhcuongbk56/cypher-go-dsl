package cypher_go_dsl

import (
	"fmt"
)

type Node struct {
	symbolicName *SymbolicName
	labels       []string
	properties   *Properties
}

type NodeLabel struct {
	value string
}

func (n NodeLabel) Accept(visitor Visitor) {
	visitor.Enter(n)
	visitor.Leave(n)
}

func (node Node) Accept(visitor Visitor) {
	visitor.Enter(node)
	visitor.Leave(node)
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
		*properties, err = NewMapExpression(keysAndValues)
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




