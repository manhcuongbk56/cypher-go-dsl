package cypher_go_dsl

import "strings"

func NewNode(primaryLabel string) Node {
	var labels = make([]NodeLabel, 0)
	labels = append(labels, NodeLabel{value: primaryLabel})
	return Node{
		labels: labels,
		notNil: true,
	}
}

func AnyNode() Node {
	return NodeCreate()
}
func AnyNode1(symbolicName string) Node {
	return NodeCreate()
}

func NewNodeWithLabels(primaryLabel string, additionalLabel ...string) Node {
	var labels = make([]NodeLabel, 0)
	labels = append(labels, NodeLabel{value: primaryLabel})
	for _, label := range additionalLabel {
		labels = append(labels, NodeLabel{value: label})
	}
	return Node{
		labels: labels,
	}
}

func Matchs(element ...PatternElement) OngoingReadingWithoutWhere {
	return NewDefaultBuilder().match(element...)
}

func Sort(expression Expression) SortItem {
	return SortItem{
		expression: expression,
		direction:  SortDirection{value: UNDEFINED},
	}
}

func escapeName(name string) string {
	return "`" + strings.ReplaceAll(name, "`", "``") + "`"
}

func ListOf(expressions ...Expression) ExpressionList {
	return ExpressionListCreate(expressions)
}

func Name(value string) SymbolicName {
	return SymbolicNameCreate(value)
}
