package cypherit

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestShouldRenderUnions(t *testing.T) {
	statement1, err := cypher.MatchElements(bikeNode).
		WhereConditionContainer(bikeNode.Property("a").IsEqualTo(cypher.LiteralOf("A"))).
		ReturningByNamed(bikeNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	statement2, err := cypher.MatchElements(bikeNode).
		WhereConditionContainer(bikeNode.Property("b").IsEqualTo(cypher.LiteralOf("B"))).
		ReturningByNamed(bikeNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	statement3, err := cypher.MatchElements(bikeNode).
		WhereConditionContainer(bikeNode.Property("c").IsEqualTo(cypher.LiteralOf("C"))).
		ReturningByNamed(bikeNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	statement4 := cypher.CypherUnion(statement1, statement2, statement3)
	AssertStatement(t, statement4, "MATCH (b:`Bike`) WHERE b.a = 'A' RETURN b UNION MATCH (b) WHERE b.b = 'B' RETURN b UNION"+
		" MATCH (b) WHERE b.c = 'C' RETURN b")
}

func TestShouldRenderAllUnions(t *testing.T) {
	statement1, err := cypher.MatchElements(bikeNode).
		WhereConditionContainer(bikeNode.Property("a").IsEqualTo(cypher.LiteralOf("A"))).
		ReturningByNamed(bikeNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	statement2, err := cypher.MatchElements(bikeNode).
		WhereConditionContainer(bikeNode.Property("b").IsEqualTo(cypher.LiteralOf("B"))).
		ReturningByNamed(bikeNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	statement3 := cypher.CypherUnionAll(statement1, statement2)
	AssertStatement(t, statement3, "MATCH (b:`Bike`) WHERE b.a = 'A' RETURN b UNION ALL MATCH (b) WHERE b.b = 'B' RETURN b")
}

func TestShouldAppendToExistingUnions(t *testing.T) {
	statement1, err := cypher.MatchElements(bikeNode).
		WhereConditionContainer(bikeNode.Property("a").IsEqualTo(cypher.LiteralOf("A"))).
		ReturningByNamed(bikeNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	statement2, err := cypher.MatchElements(bikeNode).
		WhereConditionContainer(bikeNode.Property("b").IsEqualTo(cypher.LiteralOf("B"))).
		ReturningByNamed(bikeNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	statement := cypher.CypherUnionAll(statement1, statement2)
	statement3, err := cypher.MatchElements(bikeNode).
		WhereConditionContainer(bikeNode.Property("c").IsEqualTo(cypher.LiteralOf("C"))).
		ReturningByNamed(bikeNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	statement4, err := cypher.MatchElements(bikeNode).
		WhereConditionContainer(bikeNode.Property("d").IsEqualTo(cypher.LiteralOf("D"))).
		ReturningByNamed(bikeNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	statement = cypher.CypherUnionAll(statement, statement3, statement4)
	AssertStatement(t, statement, "MATCH (b:`Bike`) WHERE b.a = 'A' RETURN b UNION ALL MATCH (b) WHERE b.b = 'B' RETURN b "+
		"UNION ALL MATCH (b) WHERE b.c = 'C' RETURN b UNION ALL MATCH (b) WHERE b.d = 'D' RETURN"+
		" b")
}

func TestShouldNotMix(t *testing.T) {
	statement1, err := cypher.MatchElements(bikeNode).
		WhereConditionContainer(bikeNode.Property("a").IsEqualTo(cypher.LiteralOf("A"))).
		ReturningByNamed(bikeNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	statement2, err := cypher.MatchElements(bikeNode).
		WhereConditionContainer(bikeNode.Property("b").IsEqualTo(cypher.LiteralOf("B"))).
		ReturningByNamed(bikeNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	statement := cypher.CypherUnionAll(statement1, statement2)
	statement3, err := cypher.MatchElements(bikeNode).
		WhereConditionContainer(bikeNode.Property("c").IsEqualTo(cypher.LiteralOf("C"))).
		ReturningByNamed(bikeNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	statement = cypher.CypherUnion(statement, statement3)
	AssertStatementError(t, statement, "cannot mix union and union all")
}
