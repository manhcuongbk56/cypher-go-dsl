package cypher

import "errors"

type HasLabelCondition struct {
	ConditionContainer
	nodeName   SymbolicName
	nodeLabels []NodeLabel
	key        string
	err        error
	notNil     bool
}

func HasLabelConditionCreate(nodeName SymbolicName, labels ...string) HasLabelCondition {
	if nodeName.getError() != nil {
		return HasLabelConditionError(nodeName.getError())
	}
	if !nodeName.isNotNil() {
		return HasLabelConditionError(errors.New("symbolic name is required"))
	}
	if labels == nil || len(labels) == 0 {
		return HasLabelConditionError(errors.New("labels is required"))
	}
	nodeLabels := make([]NodeLabel, len(labels))
	for i := range labels {
		nodeLabels[i] = NodeLabelCreate(labels[i])
	}
	hasLabelCondition := HasLabelCondition{
		nodeName:   nodeName,
		nodeLabels: nodeLabels,
		notNil:     true,
	}
	hasLabelCondition.key = getAddress(&hasLabelCondition)
	hasLabelCondition.ConditionContainer = ConditionWrap(hasLabelCondition)
	return hasLabelCondition
}

func HasLabelConditionCreate1(nodeName SymbolicName, labels []NodeLabel) HasLabelCondition {
	if nodeName.getError() != nil {
		return HasLabelConditionError(nodeName.getError())
	}
	if !nodeName.isNotNil() {
		return HasLabelConditionError(errors.New("symbolic name is required"))
	}
	if labels == nil || len(labels) == 0 {
		return HasLabelConditionError(errors.New("labels is required"))
	}
	hasLabelCondition := HasLabelCondition{
		nodeName:   nodeName,
		nodeLabels: labels,
		notNil:     true,
	}
	hasLabelCondition.key = getAddress(&hasLabelCondition)
	hasLabelCondition.ConditionContainer = ConditionWrap(hasLabelCondition)
	return hasLabelCondition
}

func HasLabelConditionError(err error) HasLabelCondition {
	return HasLabelCondition{
		err: err,
	}
}

func (h HasLabelCondition) getError() error {
	return h.err
}

func (h HasLabelCondition) accept(visitor *CypherRenderer) {
	visitor.enter(h)
	h.nodeName.accept(visitor)
	for _, label := range h.nodeLabels {
		label.accept(visitor)
	}
	visitor.leave(h)
}

func (h HasLabelCondition) enter(renderer *CypherRenderer) {
	panic("implement me")
}

func (h HasLabelCondition) leave(renderer *CypherRenderer) {
	panic("implement me")
}

func (h HasLabelCondition) getKey() string {
	return h.key
}

func (h HasLabelCondition) isNotNil() bool {
	return h.notNil
}

func (h HasLabelCondition) GetExpressionType() ExpressionType {
	return "HasLabelCondition"
}

func (h HasLabelCondition) getConditionType() string {
	return "HasLabelCondition"
}
