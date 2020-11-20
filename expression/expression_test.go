package expression

import (
	v "cypher-go-dsl/visitable"
	"fmt"
	"testing"
)

func TestType(t *testing.T) {
	var got interface{} = Comparison{}
	visitable, ok := got.(v.Visitable)
	hasexpression, ok1 := got.(HasExpression)
	fmt.Print(ok)
	fmt.Print(ok1)
	fmt.Print(visitable)
	fmt.Print(hasexpression)
}


