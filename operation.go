package cypher

import (
	"errors"
	"golang.org/x/xerrors"
)

type Operation struct {
	ExpressionContainer
	left     Expression
	operator Operator
	right    Visitable
	key      string
	notNil   bool
	err      error
}

func OperationCreate(left Expression, operator Operator, right Expression) Operation {
	if left != nil && left.getError() != nil {
		return OperationError(left.getError())
	}
	if operator.getError() != nil {
		return OperationError(operator.getError())
	}
	if right != nil && right.getError() != nil {
		return OperationError(right.getError())
	}
	if left == nil || !left.isNotNil() {
		return OperationError(errors.New("operation: left can not be nil"))
	}
	if !operator.isNotNil() {
		return OperationError(errors.New("operation: operator can not be nil"))
	}
	if right == nil || !right.isNotNil() {
		return OperationError(errors.New("operation: right can not be nil"))
	}
	o := Operation{
		left:     left,
		operator: operator,
		right:    right,
		notNil:   true,
	}
	o.key = getAddress(&o)
	o.ExpressionContainer = ExpressionWrap(o)
	return o
}

func OperationCreate1(left Expression, operator Operator, right NodeLabel) Operation {
	if left != nil && left.getError() != nil {
		return OperationError(left.getError())
	}
	if operator.getError() != nil {
		return OperationError(operator.getError())
	}
	if right.getError() != nil {
		return OperationError(right.getError())
	}
	if left == nil || left.isNotNil() {
		return OperationError(errors.New("left can not be nil"))
	}
	if operator.isNotNil() {
		return OperationError(errors.New("operator can not be nil"))
	}
	if right.isNotNil() {
		return OperationError(errors.New("right can not be nil"))
	}
	o := Operation{
		left:     left,
		operator: operator,
		right:    right,
		notNil:   true,
	}
	o.key = getAddress(&o)
	o.ExpressionContainer = ExpressionWrap(o)
	return o
}

func OperationCreate2(op1 Node, operator Operator, nodeLabels ...string) Operation {
	if op1.getError() != nil {
		return OperationError(op1.getError())
	}
	if operator.getError() != nil {
		return OperationError(operator.getError())
	}
	if !op1.isNotNil() {
		return OperationError(errors.New("left can not be nil"))
	}
	if !operator.isNotNil() {
		return OperationError(errors.New("operator can not be nil"))
	}
	if !(operator.representation == SET_LABEL.representation || operator.representation == REMOVE_LABEL.representation) {
		return OperationError(xerrors.Errorf("operator %s can not use to modify label", operator.representation))
	}
	if nodeLabels == nil || len(nodeLabels) == 0 {
		return OperationError(errors.New("labels can not be nil or empty"))
	}
	labels := make([]NodeLabel, 0)
	for _, nodeLabel := range nodeLabels {
		labels = append(labels, NodeLabelCreate(nodeLabel))
	}
	o := Operation{
		left:     op1.getRequiredSymbolicName(),
		operator: operator,
		right:    NodeLabelsCreate(labels),
		notNil:   true,
	}
	o.key = getAddress(&o)
	o.ExpressionContainer = ExpressionWrap(o)
	return o
}

func OperationError(err error) Operation {
	return Operation{err: err}
}

func (o Operation) getError() error {
	return o.err
}

func (o Operation) accept(visitor *CypherRenderer) {
	visitor.enter(o)
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
	return (o.operator.operatorType != PROPERTY && o.operator.operatorType != LABEL) &&
		(o.operator != EXPONENTIATION && o.operator != PIPE)
}
