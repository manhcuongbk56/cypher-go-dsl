package cypherit

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestShouldRenderParameters(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.
		MatchElements(userNode).
		WhereConditionContainer(userNode.Property("a").IsEqualTo(cypher.CypherParameter("aParameter"))).
		DetachDeleteByNamed(userNode).
		ReturningByNamed(userNode)
	Assert(t, builder, "MATCH (u:`User`) WHERE u.a = $aParameter DETACH DELETE u RETURN u")
}
