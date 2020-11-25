package cypher_go_dsl

import (
	"fmt"
	"reflect"
)



func VisitIfNotNull(dest interface{}, visitor *CypherRenderer) {
	if  !reflect.ValueOf(dest).IsNil() {
		visitable, isVisitable := dest.(Visitable)
		if isVisitable {
			if visitable == nil {
				 fmt.Print("io")
			}
			visitable.Accept(visitor)
		}
	}
}

func VisitIfNotNullA(dest Visitable, visitor *CypherRenderer) {
	if  reflect.ValueOf(dest).IsNil() {
		visitable, isVisitable := dest.(Visitable)
		if isVisitable {
			if visitable == nil {
				fmt.Print("io")
			}
			visitable.Accept(visitor)
		}
	}
}


type VisitableType int;

const (
	MatchVisitable VisitableType = 1
	NodeVisitable = 2
	NodeLabelVisitable = 3
	NumberLiteralVisitable = 4
	StringLiteralVisitable = 5
	WhereVisitable = 6
	SymbolicNameVisitable = 7
	SortItemVisitable = 8
	SortDirectionVisitable = 9
	SkipVisitable = 10
	ReturnBodyVisitable = 11
	ReturnVisitable = 12
	RelationshipTypesVisitable = 13
	PatternVisitable = 14
	OrderVisitable = 15
	EntryExpressionVisitable = 16
	DistinctVisitable = 17
	SinglePartQueryVisitable = 18
	RelationshipVisitable = 19
	RelationshipChainVisitable = 20
	LimitVisitable = 21
	ExpressionListVisitable = 22
	AliasedExpressionVisitable = 23
	PropertiesVisitable = 24
	MapExpressionVisitable = 24
	RelationshipDetailsVisitable = 25
	LiteralVisitable = 26
)

type Visitable interface {
	Accept(visitor *CypherRenderer)
	Enter(renderer *CypherRenderer)
	Leave(renderer *CypherRenderer)
	GetType() VisitableType
}


type SubVisitable interface {
	Visitable
	PrepareVisit(visitable Visitable) Visitable
}

type SubsVisitable struct {
	subs []Visitable
}
