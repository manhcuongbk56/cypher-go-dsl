package cypher_go_dsl

type Case struct {
	caseElse       CaseElse
	caseExpression Expression
	caseWhenThens  []CaseWhenThen
	caseType       string
	key            string
	notNil         bool
	err            error
}

var SIMPLE_CASE = "SimpleCase"
var ENDING_SIMPLE_CASE = "EndingSimpleCase"
var GENERIC_CASE = "GenericCase"
var ENDING_GENERIC_CASE = "EndingGenericCase"

func SimpleCaseCreate(caseExpression Expression, caseWhenThens []CaseWhenThen) Case {
	if caseExpression != nil && caseExpression.getError() != nil {
		return Case{err: caseExpression.getError()}
	}
	for _, caseWhenThen := range caseWhenThens {
		if caseWhenThen.getError() != nil {
			return Case{err: caseWhenThen.getError()}
		}
	}
	simpleCase := Case{
		caseExpression: caseExpression,
		caseWhenThens:  caseWhenThens,
		caseType:       SIMPLE_CASE,
		notNil:         true,
	}
	simpleCase.key = getAddress(&simpleCase)
	return simpleCase
}

func SimpleCaseCreate1(caseExpression Expression) Case {
	if caseExpression != nil && caseExpression.getError() != nil {
		return Case{err: caseExpression.getError()}
	}
	simpleCase := Case{
		caseExpression: caseExpression,
		caseWhenThens:  make([]CaseWhenThen, 0),
		caseType:       SIMPLE_CASE,
		notNil:         true,
	}
	simpleCase.key = getAddress(&simpleCase)
	return simpleCase
}

func EndingSimpleCaseCreate(caseExpression Expression, caseWhenThens []CaseWhenThen) Case {
	if caseExpression != nil && caseExpression.getError() != nil {
		return Case{err: caseExpression.getError()}
	}
	for _, caseWhenThen := range caseWhenThens {
		if caseWhenThen.getError() != nil {
			return Case{err: caseWhenThen.getError()}
		}
	}
	simpleCase := Case{
		caseExpression: caseExpression,
		caseWhenThens:  caseWhenThens,
		caseType:       ENDING_SIMPLE_CASE,
		notNil:         true,
	}
	simpleCase.key = getAddress(&simpleCase)
	return simpleCase
}

func GenericCaseCreate1() Case {
	genericCase := Case{
		caseWhenThens: make([]CaseWhenThen, 0),
		notNil:        true,
		caseType:      GENERIC_CASE,
	}
	genericCase.key = getAddress(&genericCase)
	return genericCase
}

func GenericCaseCreate(caseWhenThens []CaseWhenThen) Case {
	for _, caseWhenThen := range caseWhenThens {
		if caseWhenThen.getError() != nil {
			return Case{err: caseWhenThen.getError()}
		}
	}
	genericCase := Case{
		caseWhenThens: caseWhenThens,
		notNil:        true,
		caseType:      GENERIC_CASE,
	}
	genericCase.key = getAddress(&genericCase)
	return genericCase
}

func EndingGenericCaseCreate(caseWhenThens []CaseWhenThen) Case {
	for _, caseWhenThen := range caseWhenThens {
		if caseWhenThen.getError() != nil {
			return Case{err: caseWhenThen.getError()}
		}
	}
	genericCase := Case{
		caseWhenThens: caseWhenThens,
		notNil:        true,
		caseType:      ENDING_GENERIC_CASE,
	}
	genericCase.key = getAddress(&genericCase)
	return genericCase
}

func (c Case) When(nextExpression Expression) OngoingWhenThen {
	return OngoingWhenThenCreate(&c, nextExpression)
}

func (c Case) getError() error {
	return c.err
}

func (c Case) accept(visitor *CypherRenderer) {
	visitor.enter(c)
	if c.caseExpression != nil && c.caseExpression.isNotNil() {
		c.caseExpression.accept(visitor)
	}
	for _, caseWhenThen := range c.caseWhenThens {
		caseWhenThen.accept(visitor)
	}
	if c.caseElse.isNotNil() {
		c.caseElse.accept(visitor)
	}
	visitor.leave(c)
}

func (c Case) enter(renderer *CypherRenderer) {
	if c.caseType == SIMPLE_CASE {
		renderer.append("CASE ")
		return
	}
	if c.caseType == GENERIC_CASE {
		renderer.append("CASE")
	}
}

func (c Case) leave(renderer *CypherRenderer) {
	if c.caseType == GENERIC_CASE {
		renderer.append(" END")
	}
}

func (c Case) getKey() string {
	return c.key
}

func (c Case) isNotNil() bool {
	return c.notNil
}

func (c Case) GetExpressionType() ExpressionType {
	return "Case"
}

func (c Case) getConditionType() string {
	return "Case"
}

func (c Case) ElseDefault(defaultExpression Expression) CaseEnding {
	caseElse := CaseElseCreate(defaultExpression)
	if caseElse.getError() != nil {
		return Case{err: caseElse.getError()}
	}
	return c
}

type CaseEnding interface {
	Condition
	/**
	 * Adds a new {@code WHEN} block.
	 *
	 * @param expression A new when expression.
	 * @return An ongoing when builder.
	 */
	When(expression Expression) OngoingWhenThen
	/**
	 * Ends this case expression with a default expression to evaluate.
	 *
	 * @param defaultExpression The new default expression
	 * @return An ongoing When builder.
	 */
	ElseDefault(defaultExpression Expression) CaseEnding
}

//OngoingWhenThen
type OngoingWhenThen struct {
	caseInstance   *Case
	whenExpression Expression
	err            error
}

func OngoingWhenThenCreate(caseInstance *Case, whenExpression Expression) OngoingWhenThen {
	if whenExpression != nil && whenExpression.getError() != nil {
		return OngoingWhenThen{err: whenExpression.getError()}
	}
	return OngoingWhenThen{whenExpression: whenExpression, caseInstance: caseInstance}
}

func (o OngoingWhenThen) then(expression Expression) CaseEnding {
	caseWhenThen := CaseWhenThenCreate(o.whenExpression, expression)
	o.caseInstance.caseWhenThens = append(o.caseInstance.caseWhenThens, caseWhenThen)
	if o.caseInstance.caseExpression != nil && o.caseInstance.caseExpression.isNotNil() {
		return EndingSimpleCaseCreate(o.caseInstance.caseExpression, o.caseInstance.caseWhenThens)
	} else {
		return EndingGenericCaseCreate(o.caseInstance.caseWhenThens)
	}
}

//CaseWhenThen
type CaseWhenThen struct {
	whenExpression Expression
	thenExpression Expression
	caseInstance   *Case
	key            string
	err            error
	notNil         bool
}

func CaseWhenThenCreate(thenExpression Expression, whenExpression Expression) CaseWhenThen {
	if thenExpression != nil && thenExpression.getError() != nil {
		return CaseWhenThen{err: thenExpression.getError()}
	}
	if whenExpression != nil && whenExpression.getError() != nil {
		return CaseWhenThen{err: whenExpression.getError()}
	}
	caseWhenThen := CaseWhenThen{
		thenExpression: thenExpression,
		whenExpression: whenExpression,
		notNil:         true,
	}
	caseWhenThen.key = getAddress(&caseWhenThen)
	return caseWhenThen
}

func (c CaseWhenThen) When(nextExpression Expression) OngoingWhenThen {
	return OngoingWhenThenCreate(c.caseInstance, nextExpression)
}

func (c CaseWhenThen) getError() error {
	return c.err
}

func (c CaseWhenThen) accept(visitor *CypherRenderer) {
	visitor.enter(c)
	c.whenExpression.accept(visitor)
	visitor.leave(c)
	c.thenExpression.accept(visitor)
}

func (c CaseWhenThen) enter(renderer *CypherRenderer) {
	renderer.append(" WHEN ")
}

func (c CaseWhenThen) leave(renderer *CypherRenderer) {
	renderer.append(" THEN ")
}

func (c CaseWhenThen) getKey() string {
	return c.key
}

func (c CaseWhenThen) isNotNil() bool {
	return c.notNil
}

//CaseElse
type CaseElse struct {
	elseExpression Expression
	key            string
	err            error
	notNil         bool
}

func CaseElseCreate(elseExpression Expression) CaseElse {
	if elseExpression != nil && elseExpression.getError() != nil {
		return CaseElse{err: elseExpression.getError()}
	}
	caseElse := CaseElse{elseExpression: elseExpression, notNil: true}
	caseElse.key = getAddress(&caseElse)
	return caseElse
}

func (c CaseElse) getError() error {
	return c.err
}

func (c CaseElse) accept(visitor *CypherRenderer) {
	visitor.enter(c)
	c.elseExpression.accept(visitor)
	visitor.leave(c)
}

func (c CaseElse) enter(renderer *CypherRenderer) {
	renderer.append(" ELSE ")
}

func (c CaseElse) leave(renderer *CypherRenderer) {
}

func (c CaseElse) getKey() string {
	return c.key
}

func (c CaseElse) isNotNil() bool {
	return c.notNil
}
