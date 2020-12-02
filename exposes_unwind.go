package cypher_go_dsl

type ExposesUnwind interface {

	unwinds(expression ...Expression) OngoingUnwind
	unwindByString(variable string) OngoingUnwind
	unwind(expression Expression) OngoingUnwind
}