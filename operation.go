package cypher_go_dsl

import "fmt"

type Operation struct {
	left     Expression
	operator Operator
	right    Visitable
	key      string
	notNil   bool
	err      error
}

func OperationCreate(left Expression, operator Operator, right Expression) Operation {
	o := Operation{
		left:     left,
		operator: operator,
		right:    right,
	}
	o.key = getAddress(&o)
	return o
}

func OperationCreate1(left Expression, operator Operator, right NodeLabel) Operation {
	o := Operation{
		left:     left,
		operator: operator,
		right:    right,
	}
	o.key = getAddress(&o)
	return o
}

func OperationCreate2(op1 Node, operator Operator, nodeLabels ...string) Operation {
	labels := make([]NodeLabel, 0)
	for _, nodeLabel := range nodeLabels {
		labels = append(labels, NodeLabelCreate(nodeLabel))
	}
	o := Operation{
		left:     op1.getSymbolicName(),
		operator: operator,
		right:    NodeLabelsCreate(labels),
	}
	o.key = getAddress(&o)
	return o
}

func (o Operation) getError() error {
	return o.err
}

func (o Operation) accept(visitor *CypherRenderer) {
	NameOrExpression(o.left).accept(visitor)
	o.operator.accept(visitor)
	o.right.accept(visitor)
	visitor.leave(o)
}

func (o Operation) enter(renderer *CypherRenderer) {
	if o.needsGrouping() {
		renderer.append("(")
	}
}

func (o Operation) leave(renderer *CypherRenderer) {
	if o.needsGrouping() {
		renderer.append(")")
	}
}

func (o Operation) getKey() string {
	return o.key
}

func (o Operation) isNotNil() bool {
	return o.notNil
}

func (o Operation) GetExpressionType() ExpressionType {
	return "Operation"
}

func (o Operation) needsGrouping() bool {
	return (o.operator.operatorType == PROPERTY || o.operator.operatorType == LABEL) &&
		(o.operator != EXPONENTIATION && o.operator != PIPE)
}
