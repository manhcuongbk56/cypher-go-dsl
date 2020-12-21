package cypher

import "fmt"

func ExpressionsToVisitables(expressions []Expression) []Visitable {
	visitables := make([]Visitable, len(expressions))
	for i := range expressions {
		visitables[i] = expressions[i]
	}
	return visitables
}

func PatternElementsToVisitables(patterns []PatternElement) []Visitable {
	visitables := make([]Visitable, len(patterns))
	for i := range patterns {
		visitables[i] = patterns[i]
	}
	return visitables
}

func getAddress(a interface{}) string {
	return fmt.Sprintf("%p", a)
}
