package cypher_go_dsl

import (
	"fmt"
	"testing"
)

func TestType(t *testing.T) {
	var got interface{} = Comparison{}
	visitable, ok := got.(Visitable)
	hasexpression, ok1 := got.(Expression)
	fmt.Print(ok)
	fmt.Print(ok1)
	fmt.Print(visitable)
	fmt.Print(hasexpression)
}


