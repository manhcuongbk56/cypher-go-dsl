package cypherit

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestValueAtShouldWork(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.CypherReturning(cypher.CypherValueAt(cypher.FunctionRangeRaw(0, 10), 3))

	Assert(t, builder, "RETURN range(0, 10)[3]")
}
