package cypher

type ExpressionContainer struct {
	expression Expression
}

func ExpressionWrap(expression Expression) ExpressionContainer {
	return ExpressionContainer{
		expression: expression,
	}
}

func (e ExpressionContainer) Get() Expression {
	return e.expression
}

func (e ExpressionContainer) As(alias string) ExpressionContainer {
	e.expression = AliasedExpressionCreate(e.expression, alias)
	return e
}

func (e ExpressionContainer) IsEqualTo(rhs Expression) ConditionContainer {
	e.expression = ComparisonCreate(e.expression, EQUALITY, rhs)
	return ConditionContainer{e}
}

func (e ExpressionContainer) IsNotEqualTo(rhs Expression) ConditionContainer {
	e.expression = ComparisonCreate(e.expression, INEQUALITY, rhs)
	return ConditionContainer{e}
}

func (e ExpressionContainer) Lt(rhs Expression) ConditionContainer {
	e.expression = ComparisonCreate(e.expression, LESS_THAN, rhs)
	return ConditionContainer{e}
}

func (e ExpressionContainer) Lte(rhs Expression) ConditionContainer {
	e.expression = ComparisonCreate(e.expression, LESS_THAN_OR_EQUAL_TO, rhs)
	return ConditionContainer{e}
}

func (e ExpressionContainer) Gt(rhs Expression) ConditionContainer {
	e.expression = ComparisonCreate(e.expression, GREATER_THAN, rhs)
	return ConditionContainer{e}
}

func (e ExpressionContainer) Gte(rhs Expression) ConditionContainer {
	e.expression = ComparisonCreate(e.expression, GREATER_THAN_OR_EQUAL_TO, rhs)
	return ConditionContainer{e}
}

func (e ExpressionContainer) IsTrue() ConditionContainer {
	return e.IsEqualTo(BooleanLiteralCreate(true))
}

func (e ExpressionContainer) IsFalse() ConditionContainer {
	return e.IsEqualTo(BooleanLiteralCreate(false))
}

func (e ExpressionContainer) Matches(expression Expression) ConditionContainer {
	e.expression = ComparisonCreate(e.expression, MATCHES, expression)
	return ConditionContainer{e}
}

func (e ExpressionContainer) MatchesPattern(pattern string) ConditionContainer {
	e.expression = ComparisonCreate(e.expression, MATCHES, StringLiteralCreate(pattern))
	return ConditionContainer{e}
}

func (e ExpressionContainer) StartWiths(rhs Expression) ConditionContainer {
	e.expression = ComparisonCreate(e.expression, STARTS_WITH, rhs)
	return ConditionContainer{e}
}

func (e ExpressionContainer) Contains(rhs Expression) ConditionContainer {
	e.expression = ComparisonCreate(e.expression, CONTAINS, rhs)
	return ConditionContainer{e}
}

func (e ExpressionContainer) EndsWith(rhs Expression) ConditionContainer {
	e.expression = ComparisonCreate(e.expression, ENDS_WITH, rhs)
	return ConditionContainer{e}
}

func (e ExpressionContainer) Concat(rhs Expression) ExpressionContainer {
	e.expression = OperationConcat(e.expression, rhs)
	return e
}

func (e ExpressionContainer) Add(rhs Expression) ExpressionContainer {
	e.expression = OperationAdd(e.expression, rhs)
	return e
}

func (e ExpressionContainer) Subtract(rhs Expression) ExpressionContainer {
	e.expression = OperationSubtract(e.expression, rhs)
	return e
}

func (e ExpressionContainer) Multiply(rhs Expression) ExpressionContainer {
	e.expression = OperationMultiply(e.expression, rhs)
	return e
}

func (e ExpressionContainer) Divide(rhs Expression) ExpressionContainer {
	e.expression = OperationDivide(e.expression, rhs)
	return e
}

func (e ExpressionContainer) Remainder(rhs Expression) ExpressionContainer {
	e.expression = OperationRemainder(e.expression, rhs)
	return e
}

func (e ExpressionContainer) Pow(rhs Expression) ExpressionContainer {
	e.expression = OperationPow(e.expression, rhs)
	return e
}

func (e ExpressionContainer) IsNull() ConditionContainer {
	e.expression = ComparisonCreate1(IS_NULL, e.expression)
	return ConditionContainer{e}
}

func (e ExpressionContainer) IsNotNull() ConditionContainer {
	e.expression = ComparisonCreate1(IS_NOT_NULL, e.expression)
	return ConditionContainer{e}
}

func (e ExpressionContainer) In(haystack Expression) ConditionContainer {
	e.expression = ComparisonCreate(e.expression, IN, haystack)
	return ConditionContainer{e}
}

func (e ExpressionContainer) IsEmpty() ConditionContainer {
	e.expression = FunctionSize(e.expression)
	return e.IsEqualTo(NumberLiteralCreate1(0))
}

func (e ExpressionContainer) Descending() SortItem {
	return CreateDescendingSortItem(e.expression)
}

func (e ExpressionContainer) Ascending() SortItem {
	return CreateAscendingSortItem(e.expression)
}
