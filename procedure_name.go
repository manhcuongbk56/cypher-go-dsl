package cypher_go_dsl

import (
	"fmt"
	"strings"
)

type ProcedureName struct {
	optionalNamespace Namespace
	value             string
	key               string
	notNil            bool
	err error
}

func ProcedureNameCreate(namespaceAndProcedure ...string) ProcedureName {
	if len(namespaceAndProcedure) == 1 {
		return ProcedureNameCreate1(namespaceAndProcedure[0])
	} else {
		values := make([]string, 0)
		inputLength := len(namespaceAndProcedure)
		copy(values, namespaceAndProcedure[1:inputLength-1])
		namespace := NameSpaceCreate(values)
		return ProcedureNameCreate2(namespace, namespaceAndProcedure[inputLength-1])
	}
}

func ProcedureNameCreate1(value string) ProcedureName {
	return ProcedureName{
		value:  value,
		notNil: true,
	}
}

func ProcedureNameCreate2(namespace Namespace, value string) ProcedureName {
	return ProcedureName{
		value:             value,
		optionalNamespace: namespace,
		notNil:            true,
	}
}

func (p ProcedureName) getQualifiedName() string {
	namespace := ""
	if p.optionalNamespace.isNotNil() {
		namespace = strings.Join(p.optionalNamespace.content[:], ".")
	}
	return namespace + "." + p.value
}

func (p ProcedureName) getError() error {
	return p.err
}

func (p ProcedureName) accept(visitor *CypherRenderer) {
	p.key = fmt.Sprint(&p)
	visitor.enter(p)
	VisitIfNotNull(p.optionalNamespace, visitor)
	visitor.leave(p)
}

func (p ProcedureName) enter(renderer *CypherRenderer) {
}

func (p ProcedureName) leave(renderer *CypherRenderer) {
}

func (p ProcedureName) getKey() string {
	return p.key
}

func (p ProcedureName) isNotNil() bool {
	return p.notNil
}
