package cypher

type PatternElement interface {
	IsPatternElement() bool
	Visitable
}
