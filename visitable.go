package cypher_go_dsl

import (
	"fmt"
	"reflect"
)

func VisitIfNotNull(dest interface{}, visitor *CypherRenderer) {
	if !reflect.ValueOf(dest).IsNil() {
		visitable, isVisitable := dest.(Visitable)
		if isVisitable {
			if visitable == nil {
				fmt.Print("io")
			}
			visitable.accept(visitor)
		}
	}
}

func VisitIfNotNullA(dest Visitable, visitor *CypherRenderer) {
	if reflect.ValueOf(dest).IsNil() {
		visitable, isVisitable := dest.(Visitable)
		if isVisitable {
			if visitable == nil {
				fmt.Print("io")
			}
			visitable.accept(visitor)
		}
	}
}

type Visitable interface {
	accept(visitor *CypherRenderer)
	enter(renderer *CypherRenderer)
	leave(renderer *CypherRenderer)
	getKey() string
}

type SubVisitable interface {
	Visitable
	PrepareVisit(visitable Visitable) Visitable
}

type SubsVisitable struct {
	subs []Visitable
}
