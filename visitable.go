package cypher_go_dsl

type Visitor interface {
	Enter(visitable Visitable)
	Leave(visitable Visitable)
}

type ExampleVisitor struct {

}

func (e ExampleVisitor) Enter(visitable Visitable) {
	panic("implement me")
}

func (e ExampleVisitor) Leave(visitable Visitable) {
	panic("implement me")
}

func VisitIfNotNull(dest interface{}, visitor Visitor) {
	if dest != nil {
		visitable, isVisitable := dest.(Visitable)
		if isVisitable {
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
)

type Visitable interface {
	Accept(visitor Visitor)
	GetType() VisitableType
}


type SubVisitable interface {
	Visitable
	PrepareVisit(visitable Visitable) Visitable
}

type SubsVisitable struct {
	subs []Visitable
}
