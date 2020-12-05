package cypher_go_dsl

import "fmt"

type ProcedureCall struct {
	name          ProcedureName
	arguments     Arguments
	yieldItems    YieldItems
	optionalWhere Where
	key           string
	notNil        bool
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
	panic("implement me")
}

func (p ProcedureCall) leave(renderer *CypherRenderer) {
	panic("implement me")
}

func (p ProcedureCall) getKey() string {
	panic("implement me")
}

func (p ProcedureCall) isNotNil() bool {
	return p.notNil
}
