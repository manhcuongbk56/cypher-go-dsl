package cypher_go_dsl

type EnterMap struct {
	visitableEnterMap map[VisitableType]func(Visitor, Visitable)
	visitableLeaveMap map[VisitableType]func(Visitor, Visitable)
}



func testOOOO()  {
	mapFunc := map[VisitableType]enter {
		MatchVisitable: EnterMatch,
	}
	match := Match{}
	a := ExampleVisitor{}
	mapFunc[match.GetType()](a, match)
}

func EnterMatch(visitor Visitor, visitable interface{})  {

}
