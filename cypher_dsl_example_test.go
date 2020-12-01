package cypher_go_dsl

import (
	"fmt"
	"testing"
)

func TestFindAllMovies(t *testing.T) {
	movie := NewNode("Device").Named("d")
	statement := Matchs(movie).
		returning(nil).
		Build()
	query := NewRenderer().Render(statement)
	fmt.Println(query)
}

func TestExposesReturning(t *testing.T)  {
	exposesStruct := ExposesReturningStruct{}
	exposesStruct.returning("a", "b", "c")
}
