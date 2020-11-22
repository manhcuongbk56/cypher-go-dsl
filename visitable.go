package cypher_go_dsl

type Visitor interface {
	Enter(visitable Visitable)
	Leave(visitable Visitable)
}

func VisitIfNotNull(dest interface{}, visitor Visitor) {
	if dest != nil {
		visitable, isVisitable := dest.(Visitable)
		if isVisitable {
			visitable.Accept(visitor)
		}
	}
}

type Visitable interface {
	Accept(visitor Visitor)

}

type SubVisitable interface {
	PrepareVisit(visitable Visitable) Visitable
}

type SubsVisitable struct {
	subs []Visitable
}