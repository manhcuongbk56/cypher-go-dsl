package cypher

type ExposesUnwind interface {
	Unwinds(expression ...Expression) OngoingUnwind
	UnwindByString(variable string) OngoingUnwind
	Unwind(expression Expression) OngoingUnwind
}
