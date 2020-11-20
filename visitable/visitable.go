package visitable

type Visitor interface {
	Enter(visitable Visitable)
	Leave(visitable Visitable)
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