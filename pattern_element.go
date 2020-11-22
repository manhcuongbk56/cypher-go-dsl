package cypher_go_dsl

type PatternElement interface {
	IsPatternElement() bool
	Visitable
}
