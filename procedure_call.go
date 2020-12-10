package cypher_go_dsl

import "fmt"

type ProcedureCallBuilder interface {
	ExposesWhere
	ExposesReturning
	BuildableStatement
	isNotNil() bool
}

type ProcedureCall struct {
	name          ProcedureName
	arguments     Arguments
	yieldItems    YieldItems
	optionalWhere Where
	key           string
	notNil        bool
	err error
}

func ProcedureCallCreate(name ProcedureName, arguments Arguments, yieldItems YieldItems, optionalWhere Where) ProcedureCall {
	return ProcedureCall{
		name:          name,
		arguments:     arguments,
		yieldItems:    yieldItems,
		optionalWhere: optionalWhere,
		notNil:        true,
	}
}

func (p ProcedureCall) doesReturnElement() bool {
	return p.yieldItems.isNotNil()
}

func (p ProcedureCall) getError() error {
	return p.err
}

func (p ProcedureCall) accept(visitor *CypherRenderer) {
	p.key = fmt.Sprint(&p)
	visitor.enter(p)
	p.name.accept(visitor)
	VisitIfNotNull(p.arguments, visitor)
	VisitIfNotNull(p.yieldItems, visitor)
	VisitIfNotNull(p.optionalWhere, visitor)
	visitor.leave(p)
}

func (p ProcedureCall) enter(renderer *CypherRenderer) {
	renderer.append("CALL ")
}

func (p ProcedureCall) leave(renderer *CypherRenderer) {
	renderer.append(".")
}

func (p ProcedureCall) getKey() string {
	return p.key
}

func (p ProcedureCall) isNotNil() bool {
	return p.notNil
}

func (p ProcedureCall) doesReturnElements() bool {
	return p.yieldItems.isNotNil()
}
