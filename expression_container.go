package cypher_go_dsl

type ExpressionContainer struct {
	expression Expression
}

func (e ExpressionContainer) accept(visitor *CypherRenderer) {
	e.expression.accept(visitor)
}

func (e ExpressionContainer) enter(renderer *CypherRenderer) {
	panic("can not enter expression container")
}

func (e ExpressionContainer) leave(renderer *CypherRenderer) {
	panic("can not leave expression container")
}

func (e ExpressionContainer) getKey() string {
	panic("Expression container have no key")
}

func (e *ExpressionContainer) As(alias string) ExpressionContainer {
	e.expression = AliasedExpressionCreate(e.expression, alias)
	return *e
}

func (e *ExpressionContainer) IsEqualTo(rhs Expression) ConditionContainer {
	e.expression = NewComparison(e.expression, EQUALITY, rhs)
	return ConditionContainer{*e}
}

func (e *ExpressionContainer) IsNotEqualTo(rhs Expression) ConditionContainer {
	e.expression = NewComparison(e.expression, INEQUALITY, rhs)
	return ConditionContainer{*e}
}

func (e *ExpressionContainer) Lt(rhs Expression) ConditionContainer {
	e.expression = NewComparison(e.expression, LESS_THAN, rhs)
	return ConditionContainer{*e}
}

func (e *ExpressionContainer) Lte(rhs Expression) ConditionContainer {
	e.expression = NewComparison(e.expression, LESS_THAN_OR_EQUAL_TO, rhs)
	return ConditionContainer{*e}
}

func (e *ExpressionContainer) Gt(rhs Expression) ConditionContainer {
	e.expression = NewComparison(e.expression, GREATER_THAN, rhs)
	return ConditionContainer{*e}
}

func (e *ExpressionContainer) Gte(rhs Expression) ConditionContainer {
	e.expression = NewComparison(e.expression, GREATER_THAN_OR_EQUAL_TO, rhs)
	return ConditionContainer{*e}
}

func (e *ExpressionContainer) IsTrue() ConditionContainer {
	return e.IsEqualTo(BooleanLiteral{
		content: true,
	})
}

func (e *ExpressionContainer) IsFalse() ConditionContainer {
	return e.IsEqualTo(BooleanLiteral{
		content: false,
	})
}

func (e *ExpressionContainer) Matches(expression Expression) ConditionContainer {
	e.expression = NewComparison(e.expression, MATCHES, expression)
	return ConditionContainer{*e}
}

func (e *ExpressionContainer) MatchesPattern(pattern string) ConditionContainer {
	e.expression = NewComparison(e.expression, MATCHES, StringLiteral{
		content: pattern,
	})
	return ConditionContainer{*e}
}

func (e *ExpressionContainer) StartWiths(rhs Expression) ConditionContainer {
	e.expression = NewComparison(e.expression, STARTS_WITH, rhs)
	return ConditionContainer{*e}
}

func (e *ExpressionContainer) Contains(rhs Expression) ConditionContainer {
	e.expression = NewComparison(e.expression, CONTAINS, rhs)
	return ConditionContainer{*e}
}

func (e *ExpressionContainer) EndsWith(rhs Expression) ConditionContainer {
	e.expression = NewComparison(e.expression, ENDS_WITH, rhs)
	return ConditionContainer{*e}
}

func (e *ExpressionContainer) Concat(rhs Expression) ExpressionContainer {
	e.expression = NewComparison(e.expression, CONCAT, rhs)
	return *e
}

func (e *ExpressionContainer) Add(rhs Expression) ExpressionContainer {
	e.expression = NewComparison(e.expression, ADDITION, rhs)
	return *e
}

func (e *ExpressionContainer) Subtract(rhs Expression) ExpressionContainer {
	e.expression = NewComparison(e.expression, SUBTRACTION, rhs)
	return *e
}

func (e *ExpressionContainer) Multiply(rhs Expression) ExpressionContainer {
	e.expression = NewComparison(e.expression, MULTIPLICATION, rhs)
	return *e
}

func (e *ExpressionContainer) Divide(rhs Expression) ExpressionContainer {
	e.expression = NewComparison(e.expression, DIVISION, rhs)
	return *e
}

func (e *ExpressionContainer) Remainder(rhs Expression) ExpressionContainer {
	e.expression = NewComparison(e.expression, MODULO_DIVISION, rhs)
	return *e
}

func (e *ExpressionContainer) Pow(rhs Expression) ExpressionContainer {
	e.expression = NewComparison(e.expression, EXPONENTIATION, rhs)
	return *e
}

func (e *ExpressionContainer) IsNull() ConditionContainer {
	e.expression = NewComparisonWithConstant(IS_NOT_NULL, e.expression)
	return ConditionContainer{*e}
}

func (e *ExpressionContainer) IsNotNull() ConditionContainer {
	e.expression = NewComparisonWithConstant(IS_NOT_NULL, e.expression)
	return ConditionContainer{*e}
}

func (e *ExpressionContainer) In(haystack Expression) ConditionContainer {
	e.expression = NewComparison(e.expression, IN, haystack)
	return ConditionContainer{*e}
}

func (e *ExpressionContainer) IsEmpty() ConditionContainer {
	e.expression = size(e.expression)
	return e.IsEqualTo(NumberLiteral{
		content: 0,
	})
}

func (e *ExpressionContainer) Descending() SortItem {
	return CreateDescendingSortItem(e.expression)
}

func (e *ExpressionContainer) Ascending() SortItem {
	return CreateAscendingSortItem(e.expression)
}
