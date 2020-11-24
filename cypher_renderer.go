package cypher_go_dsl

type CypherRenderer struct {
	visitableEnterMap map[VisitableType]enter
	visitableLeaveMap map[VisitableType]leave
}

type enter func(visitor Visitor, visitable Visitable)
type leave func(visitor Visitor, visitable Visitable)

func (renderer CypherRenderer) FindEnterFunc(visitable Visitable)  {

}

func (renderer CypherRenderer) FindLeaveFunc(visitable Visitable)  {

}

func (renderer CypherRenderer) Enter(visitable Visitable) {
	panic("implement me")
}

func (renderer CypherRenderer) Leave(visitable Visitable) {
	panic("implement me")
}

func (renderer CypherRenderer) getAliasIfSeen(visitable Visitable)  {

}
