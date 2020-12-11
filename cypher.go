package cypher_go_dsl

import "strings"

func NewNode(primaryLabel string) Node {
	return NodeCreate2(primaryLabel)
}

func AnyNode() Node {
	return NodeCreate()
}
func AnyNode1(symbolicName string) Node {
	return NodeCreate().Named(symbolicName)
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
	return DefaultStatementBuilderCreate().Match(element...)
}

func Sort(expression Expression) SortItem {
	return SortItemCreate(expression, UNDEFINED)
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
