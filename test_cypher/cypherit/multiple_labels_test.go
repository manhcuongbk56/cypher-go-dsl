package cypherit

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestMatchWithMultipleLabels(t *testing.T) {
	var builder cypher.BuildableStatement
	node := cypher.NewNodeWithLabels("a", "b", "c").NamedByString("n")
	//
	builder = cypher.
		MatchElements(node).
		ReturningByNamed(node)

	Assert(t, builder, "MATCH (n:`a`:`b`:`c`) RETURN n")
}

func TestCreateWithMultipleLabels(t *testing.T) {
	var builder cypher.BuildableStatement
	node := cypher.NewNodeWithLabels("a", "b", "c").NamedByString("n")
	//
	builder = cypher.
		CypherCreate(node)
	Assert(t, builder, "CREATE (n:`a`:`b`:`c`)")
}
