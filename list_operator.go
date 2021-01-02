package cypher

import "errors"

type ListOperator struct {
	ExpressionContainer
	targetExpression Expression
	details          ListOperatorDetails
	key              string
	notNil           bool
	err              error
}

func listOperatorCreate(targetExpression Expression, optionalStart Expression, dots Dot, optionalEnd Expression) ListOperator {
	operator := ListOperator{
		targetExpression: targetExpression,
		details:          ListOperatorDetailsCreate(optionalStart, dots, optionalEnd),
		notNil:           true,
	}
	operator.key = getAddress(&operator)
	operator.ExpressionContainer = ExpressionWrap(operator)
	return operator
}

func ListOperatorError(err error) ListOperator {
	return ListOperator{err: err}
}

/**
 * Creates a closed range with given boundaries.
 *
 * @param targetExpression The target expression for the range
 * @param start            The inclusive start
 * @param end              The exclusive end
 * @return A range literal.
 */
func SubList(targetExpression Expression, start Expression, end Expression) ListOperator {
	if targetExpression != nil && targetExpression.GetError() != nil {
		return ListOperatorError(targetExpression.GetError())
	}
	if start != nil && start.GetError() != nil {
		return ListOperatorError(start.GetError())
	}
	if end != nil && end.GetError() != nil {
		return ListOperatorError(end.GetError())
	}
	if targetExpression == nil || !targetExpression.isNotNil() {
		return ListOperatorError(errors.New("the range's target expression must not be nil"))
	}
	if start == nil || !start.isNotNil() {
		return ListOperatorError(errors.New("the range's start expression must not be nil"))
	}
	if end == nil || !end.isNotNil() {
		return ListOperatorError(errors.New("the range's end expression must not be nil"))
	}
	return listOperatorCreate(targetExpression, start, DOTS, end)
}

/**
 * Creates an open range starting at {@code start}.
 *
 * @param targetExpression The target expression for the range
 * @param start            The inclusive start
 * @return A range literal.
 */
func SubListFrom(targetExpression Expression, start Expression) ListOperator {
	if targetExpression != nil && targetExpression.GetError() != nil {
		return ListOperatorError(targetExpression.GetError())
	}
	if start != nil && start.GetError() != nil {
		return ListOperatorError(start.GetError())
	}
	if targetExpression == nil || !targetExpression.isNotNil() {
		return ListOperatorError(errors.New("the range's target expression must not be nil"))
	}
	if start == nil || !start.isNotNil() {
		return ListOperatorError(errors.New("the range's start expression must not be nil"))
	}
	return listOperatorCreate(targetExpression, start, DOTS, nil)
}

/**
 * Creates an open range starting at {@code start}.
 *
 * @param targetExpression The target expression for the range
 * @param end              The exclusive end
 * @return A range literal.
 */
func SubListUntil(targetExpression Expression, end Expression) ListOperator {
	if targetExpression != nil && targetExpression.GetError() != nil {
		return ListOperatorError(targetExpression.GetError())
	}
	if end != nil && end.GetError() != nil {
		return ListOperatorError(end.GetError())
	}
	if targetExpression == nil || !targetExpression.isNotNil() {
		return ListOperatorError(errors.New("the range's target expression must not be nil"))
	}
	if end == nil || !end.isNotNil() {
		return ListOperatorError(errors.New("the range's end expression must not be nil"))
	}
	return listOperatorCreate(targetExpression, nil, DOTS, end)
}

/**
 * Creates a single valued range at {@code index}.
 *
 * @param targetExpression The target expression for the range
 * @param index            The index of the range
 * @return A range literal.
 */
func ValueAt(targetExpression Expression, index Expression) ListOperator {
	if targetExpression != nil && targetExpression.GetError() != nil {
		return ListOperatorError(targetExpression.GetError())
	}
	if index != nil && index.GetError() != nil {
		return ListOperatorError(index.GetError())
	}
	if targetExpression == nil || !targetExpression.isNotNil() {
		return ListOperatorError(errors.New("the range's target expression must not be nil"))
	}
	if index == nil || !index.isNotNil() {
		return ListOperatorError(errors.New("the index of the range must not be nil"))
	}
	return listOperatorCreate(targetExpression, index, Dot{}, nil)
}

func (l ListOperator) GetError() error {
	return l.err
}

func (l ListOperator) accept(visitor *CypherRenderer) {
	visitor.enter(l)
	l.targetExpression.accept(visitor)
	l.details.accept(visitor)
	visitor.leave(l)
}

func (l ListOperator) enter(renderer *CypherRenderer) {
}

func (l ListOperator) leave(renderer *CypherRenderer) {
}

func (l ListOperator) getKey() string {
	return l.key
}

func (l ListOperator) isNotNil() bool {
	return l.notNil
}

func (l ListOperator) GetExpressionType() ExpressionType {
	return "ListOperator"
}

//Dot struct
type Dot struct {
	ExpressionContainer
	content string
	key     string
	notNil  bool
	err     error
}

var DOTS = DotCreate("..")

func DotCreate(content string) Dot {
	dot := Dot{
		content: content,
		notNil:  true,
	}
	dot.key = getAddress(&dot)
	dot.ExpressionContainer = ExpressionWrap(dot)
	return dot
}

func (s Dot) GetError() error {
	return s.err
}

func (s Dot) isNotNil() bool {
	return s.notNil
}

func (s Dot) getKey() string {
	return s.key
}

func (s Dot) GetExpressionType() ExpressionType {
	return "Dot"
}

func (s Dot) GetContent() interface{} {
	return s.content
}

func (s Dot) AsString() string {
	return s.content
}

func (s Dot) accept(visitor *CypherRenderer) {
	visitor.enter(s)
	visitor.leave(s)
}

func (s Dot) enter(renderer *CypherRenderer) {
	renderer.append(s.AsString())
}

func (s Dot) leave(renderer *CypherRenderer) {
}

//ListOperatorDetails struct

type ListOperatorDetails struct {
	optionalStart Expression
	dots          Dot
	optionalEnd   Expression
	key           string
	err           error
	notNil        bool
}

func ListOperatorDetailsCreate(optionalStart Expression, dots Dot, optionalEnd Expression) ListOperatorDetails {
	if optionalStart != nil && optionalStart.GetError() != nil {
		return ListOperatorDetails{err: optionalStart.GetError()}
	}
	if optionalEnd != nil && optionalEnd.GetError() != nil {
		return ListOperatorDetails{err: optionalEnd.GetError()}
	}
	operator := ListOperatorDetails{
		optionalStart: optionalStart,
		dots:          dots,
		optionalEnd:   optionalEnd,
		notNil:        true,
	}
	operator.key = getAddress(&operator)
	return operator
}

func (l ListOperatorDetails) GetError() error {
	return l.err
}

func (l ListOperatorDetails) accept(visitor *CypherRenderer) {
	visitor.enter(l)
	VisitIfNotNull(l.optionalStart, visitor)
	VisitIfNotNull(l.dots, visitor)
	VisitIfNotNull(l.optionalEnd, visitor)
	visitor.leave(l)
}

func (l ListOperatorDetails) enter(renderer *CypherRenderer) {
	renderer.append("[")
}

func (l ListOperatorDetails) leave(renderer *CypherRenderer) {
	renderer.append("]")
}

func (l ListOperatorDetails) getKey() string {
	return l.key
}

func (l ListOperatorDetails) isNotNil() bool {
	return l.notNil
}
