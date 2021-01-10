package cypherit

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestConditionChainingAnd(t *testing.T) {
	statement, err := cypher.
		Match(userNode).
		WhereConditionContainer(userNode.Property("name").IsEqualTo(cypher.LiteralOf("Test")).
			And(userNode.Property("age").IsEqualTo(cypher.LiteralOf(21)).Get())).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) WHERE (u.name = 'Test' AND u.age = 21) RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestConditionChainingOr(t *testing.T) {
	statement, err := cypher.
		Match(userNode).
		WhereConditionContainer(userNode.Property("name").IsEqualTo(cypher.LiteralOf("Test")).
			Or(userNode.Property("age").IsEqualTo(cypher.LiteralOf(21)).Get())).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) WHERE (u.name = 'Test' OR u.age = 21) RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestNestedCondition1(t *testing.T) {
	statement, err := cypher.
		Match(userNode).
		WhereConditionContainer(cypher.ConditionsIsTrue().Or(cypher.ConditionsIsFalse()).
			And(cypher.ConditionsIsTrue())).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) WHERE ((true OR false) AND true) RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestNestedCondition2(t *testing.T) {
	statement, err := cypher.
		Match(userNode).
		WhereConditionContainer(cypher.ConditionsIsTrue().Or(cypher.ConditionsIsFalse()).
			And(cypher.ConditionsIsTrue())).
		Or(cypher.ConditionsIsFalse()).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) WHERE (((true OR false) AND true) OR false) RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestNestedCondition3(t *testing.T) {
	statement, err := cypher.
		Match(userNode).
		WhereConditionContainer(cypher.ConditionsIsTrue().Or(cypher.ConditionsIsFalse()).
			And(cypher.ConditionsIsTrue())).
		Or(cypher.ConditionsIsFalse()).
		And(cypher.ConditionsIsFalse()).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) WHERE ((((true OR false) AND true) OR false) AND false) RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestNestedCondition4(t *testing.T) {
	statement, err := cypher.
		Match(userNode).
		WhereConditionContainer(cypher.ConditionsIsTrue().Or(cypher.ConditionsIsFalse()).
			And(cypher.ConditionsIsTrue())).
		Or(cypher.ConditionsIsFalse().And(cypher.ConditionsIsTrue()).Get()).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) WHERE ((true OR false) AND true OR (false AND true)) RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestNestedCondition5(t *testing.T) {
	statement, err := cypher.
		Match(userNode).
		WhereConditionContainer(cypher.ConditionsIsTrue().Or(cypher.ConditionsIsFalse()).
			And(cypher.ConditionsIsTrue())).
		Or(cypher.ConditionsIsFalse().And(cypher.ConditionsIsTrue()).Get()).
		And(cypher.ConditionsIsTrue()).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) WHERE ((true OR false) AND true OR (false AND true) AND true) RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestNestedCondition6(t *testing.T) {
	statement, err := cypher.
		Match(userNode).
		WhereConditionContainer(cypher.ConditionsIsTrue().Or(cypher.ConditionsIsFalse()).
			And(cypher.ConditionsIsTrue())).
		Or(cypher.ConditionsIsFalse().Or(cypher.ConditionsIsTrue()).Get()).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) WHERE ((true OR false) AND true OR (false OR true)) RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestNestedCondition7(t *testing.T) {
	statement, err := cypher.
		Match(userNode).
		WhereConditionContainer(cypher.ConditionsIsTrue().Or(cypher.ConditionsIsFalse()).
			And(cypher.ConditionsIsTrue()).Or(cypher.ConditionsIsFalse().Or(cypher.ConditionsIsTrue()).Get())).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) WHERE ((true OR false) AND true OR (false OR true)) RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestConditionChainingXor(t *testing.T) {
	statement, err := cypher.
		Match(userNode).
		WhereConditionContainer(userNode.Property("name").IsEqualTo(cypher.LiteralOf("Test")).
			Xor(userNode.Property("age").IsEqualTo(cypher.LiteralOf(21)).Get())).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) WHERE (u.name = 'Test' XOR u.age = 21) RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestMultipleEmptyConditionsMustCollapse(t *testing.T) {
	statement, err := cypher.
		Match(userNode).
		WhereConditionContainer(cypher.ConditionsNoCondition().Or(cypher.ConditionsNoCondition())).
		And(cypher.ConditionsNoCondition().And(cypher.ConditionsNoCondition()).Or(cypher.ConditionsNoCondition()).Get()).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestMultipleEmptyConditionsMustCollapse1(t *testing.T) {
	statement, err := cypher.
		Match(userNode).
		WhereConditionContainer(cypher.ConditionsNoCondition().Or(cypher.ConditionsNoCondition())).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestMultipleEmptyConditionsMustCollapse2(t *testing.T) {
	statement, err := cypher.
		Match(userNode).
		WhereConditionContainer(cypher.ConditionsNoCondition().And(cypher.ConditionsNoCondition()).Or(cypher.ConditionsNoCondition())).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestMultipleEmptyConditionsMustCollapse3(t *testing.T) {
	statement, err := cypher.
		Match(userNode).
		WhereConditionContainer(cypher.ConditionsNoCondition().And(userNode.Property("a").IsEqualTo(cypher.CypherLiteralTrue()).Get()).Or(cypher.ConditionsNoCondition())).
		And(cypher.ConditionsNoCondition().And(cypher.ConditionsNoCondition()).Or(cypher.ConditionsNoCondition()).Get()).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) WHERE u.a = true RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestMultipleEmptyConditionsMustCollapse4(t *testing.T) {
	no := cypher.ConditionsNoCondition()
	have := userNode.Property("a").IsEqualTo(cypher.CypherLiteralTrue()).Get()
	statement, err := cypher.
		Match(userNode).
		WhereConditionContainer(no.Or(no).Or(have)).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) WHERE u.a = true RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestMultipleEmptyConditionsMustCollapse5(t *testing.T) {
	no := cypher.ConditionsNoCondition()
	have := userNode.Property("a").IsEqualTo(cypher.CypherLiteralTrue()).Get()
	statement, err := cypher.
		Match(userNode).
		WhereConditionContainer(no.And(have).Or(no)).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) WHERE u.a = true RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestMultipleEmptyConditionsMustCollapse6(t *testing.T) {
	no := cypher.ConditionsNoCondition()
	have := userNode.Property("a").IsEqualTo(cypher.CypherLiteralTrue()).Get()
	statement, err := cypher.
		Match(userNode).
		WhereConditionContainer(no.Or(no)).
		And(have).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) WHERE u.a = true RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestMultipleEmptyConditionsMustCollapse7(t *testing.T) {
	no := cypher.ConditionsNoCondition()
	a := userNode.Property("a").IsEqualTo(cypher.CypherLiteralTrue()).Get()
	b := userNode.Property("b").IsEqualTo(cypher.CypherLiteralFalse()).Get()
	statement, err := cypher.
		Match(userNode).
		WhereConditionContainer(no.And(a).Or(no)).
		And(no.And(b).Or(no).Get()).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) WHERE (u.a = true AND u.b = false) RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestMultipleEmptyConditionsMustCollapse8(t *testing.T) {
	no := cypher.ConditionsNoCondition()
	a := userNode.Property("a").IsEqualTo(cypher.CypherLiteralTrue()).Get()
	b := userNode.Property("b").IsEqualTo(cypher.CypherLiteralFalse()).Get()
	statement, err := cypher.
		Match(userNode).
		WhereConditionContainer(no.Or(no).Or(a).And(b)).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) WHERE (u.a = true AND u.b = false) RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestMultipleEmptyConditionsMustCollapse9(t *testing.T) {
	no := cypher.ConditionsNoCondition()
	a := userNode.Property("a").IsEqualTo(cypher.CypherLiteralTrue()).Get()
	b := userNode.Property("b").IsEqualTo(cypher.CypherLiteralFalse()).Get()
	statement, err := cypher.
		Match(userNode).
		WhereConditionContainer(no.Or(a).Or(no).And(b).Or(no)).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) WHERE (u.a = true AND u.b = false) RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestChainingOnWhere(t *testing.T) {
	test := cypher.LiteralOf("Test")
	foobar := cypher.LiteralOf("foobar")
	statement, err := cypher.
		Match(userNode).
		WhereConditionContainer(userNode.Property("name").IsEqualTo(test)).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) WHERE u.name = 'Test' RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
	//
	statement, err = cypher.
		Match(userNode).
		WhereConditionContainer(userNode.Property("name").IsEqualTo(test)).
		And(userNode.Property("name").IsEqualTo(test).Get()).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ = cypher.NewRenderer().Render(statement)
	expect = "MATCH (u:`User`) WHERE (u.name = 'Test' AND u.name = 'Test') RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
	//
	statement, err = cypher.
		Match(userNode).
		WhereConditionContainer(userNode.Property("name").IsEqualTo(test)).
		And(userNode.Property("name").IsEqualTo(test).Get()).
		And(userNode.Property("name").IsEqualTo(test).Get()).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ = cypher.NewRenderer().Render(statement)
	expect = "MATCH (u:`User`) WHERE (u.name = 'Test' AND u.name = 'Test' AND u.name = 'Test') RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
	//
	statement, err = cypher.
		Match(userNode).
		WhereConditionContainer(userNode.Property("name").IsEqualTo(test)).
		Or(userNode.Property("name").IsEqualTo(test).Get()).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ = cypher.NewRenderer().Render(statement)
	expect = "MATCH (u:`User`) WHERE (u.name = 'Test' OR u.name = 'Test') RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
	//
	statement, err = cypher.
		Match(userNode).
		WhereConditionContainer(userNode.Property("name").IsEqualTo(test)).
		Or(userNode.Property("name").IsEqualTo(test).Get()).
		Or(userNode.Property("name").IsEqualTo(test).Get()).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ = cypher.NewRenderer().Render(statement)
	expect = "MATCH (u:`User`) WHERE (u.name = 'Test' OR u.name = 'Test' OR u.name = 'Test') RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
	//
	statement, err = cypher.
		Match(userNode).
		WhereConditionContainer(userNode.Property("name").IsEqualTo(test)).
		And(userNode.Property("name").IsEqualTo(test).Get()).
		Or(userNode.Property("name").IsEqualTo(foobar).Get()).
		And(userNode.Property("name").IsEqualTo(test).Get()).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ = cypher.NewRenderer().Render(statement)
	expect = "MATCH (u:`User`) WHERE (((u.name = 'Test' AND u.name = 'Test') OR u.name = 'foobar') AND u.name = 'Test') RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
	//
	statement, err = cypher.
		Match(userNode).
		WhereConditionContainer(userNode.Property("name").IsEqualTo(test)).
		Or(userNode.Property("name").IsEqualTo(foobar).Get()).
		And(userNode.Property("name").IsEqualTo(test).Get()).
		And(userNode.Property("name").IsEqualTo(test).Get()).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ = cypher.NewRenderer().Render(statement)
	expect = "MATCH (u:`User`) WHERE ((u.name = 'Test' OR u.name = 'foobar') AND u.name = 'Test' AND u.name = 'Test') RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
	//
	statement, err = cypher.
		Match(userNode).
		WhereConditionContainer(userNode.Property("name").IsEqualTo(test)).
		Or(userNode.Property("name").IsEqualTo(foobar).Get()).
		And(userNode.Property("name").IsEqualTo(test).Get()).
		Or(userNode.Property("name").IsEqualTo(foobar).Get()).
		And(userNode.Property("name").IsEqualTo(test).Get()).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ = cypher.NewRenderer().Render(statement)
	expect = "MATCH (u:`User`) WHERE ((((u.name = 'Test' OR u.name = 'foobar') AND u.name = 'Test') OR u.name = 'foobar') AND u.name = 'Test') RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
	//
	statement, err = cypher.
		Match(userNode).
		WhereConditionContainer(userNode.Property("name").IsNotNull()).
		And(userNode.Property("name").IsEqualTo(test).Get()).
		Or(userNode.Property("age").IsEqualTo(cypher.LiteralOf(21)).Get()).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ = cypher.NewRenderer().Render(statement)
	expect = "MATCH (u:`User`) WHERE ((u.name IS NOT NULL AND u.name = 'Test') OR u.age = 21) RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestChainingOnCondition(t *testing.T) {
	test := cypher.LiteralOf("Test")
	foobar := cypher.LiteralOf("foobar")
	bazbar := cypher.LiteralOf("bazbar")
	statement, err := cypher.
		Match(userNode).
		WhereConditionContainer(userNode.Property("name").IsEqualTo(test).
			Or(userNode.Property("name").IsEqualTo(foobar).Get()).
			Or(userNode.Property("name").IsEqualTo(foobar).Get())).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) WHERE (u.name = 'Test' OR u.name = 'foobar' OR u.name = 'foobar') RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
	//
	statement, err = cypher.
		Match(userNode).
		WhereConditionContainer(userNode.Property("name").IsEqualTo(test).
			And(userNode.Property("name").IsEqualTo(bazbar).Get()).
			Or(userNode.Property("name").IsEqualTo(foobar).Get()).
			Or(userNode.Property("name").IsEqualTo(foobar).Get())).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ = cypher.NewRenderer().Render(statement)
	expect = "MATCH (u:`User`) WHERE ((u.name = 'Test' AND u.name = 'bazbar') OR u.name = 'foobar' OR u.name = 'foobar') RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
	//
	statement, err = cypher.
		Match(userNode).
		WhereConditionContainer(userNode.Property("name").IsEqualTo(test).
			And(userNode.Property("name").IsEqualTo(bazbar).
				And(userNode.Property("name").IsEqualTo(foobar).Get()).Get())).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ = cypher.NewRenderer().Render(statement)
	expect = "MATCH (u:`User`) WHERE (u.name = 'Test' AND u.name = 'bazbar' AND u.name = 'foobar') RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
	//
	statement, err = cypher.
		Match(userNode).
		WhereConditionContainer(userNode.Property("name").IsEqualTo(test).
			And(userNode.Property("name").IsEqualTo(bazbar).Get()).
			Or(userNode.Property("name").IsEqualTo(foobar).Get()).
			Or(userNode.Property("name").IsEqualTo(foobar).Get())).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ = cypher.NewRenderer().Render(statement)
	expect = "MATCH (u:`User`) WHERE ((u.name = 'Test' AND u.name = 'bazbar') OR u.name = 'foobar' OR u.name = 'foobar') RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
	//
	statement, err = cypher.
		Match(userNode).
		WhereConditionContainer(userNode.Property("name").IsEqualTo(test).
			And(userNode.Property("name").IsEqualTo(bazbar).Get()).
			Or(userNode.Property("name").IsEqualTo(foobar).Get()).
			Or(userNode.Property("name").IsEqualTo(foobar).Get()).
			And(userNode.Property("name").IsEqualTo(bazbar).Get())).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ = cypher.NewRenderer().Render(statement)
	expect = "MATCH (u:`User`) WHERE (((u.name = 'Test' AND u.name = 'bazbar') OR u.name = 'foobar' OR u.name = 'foobar') AND u.name = 'bazbar') RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestChainingCombined(t *testing.T) {
	test := cypher.LiteralOf("Test")
	foobar := cypher.LiteralOf("foobar")
	bazbar := cypher.LiteralOf("bazbar")
	statement, err := cypher.
		Match(userNode).
		WhereConditionContainer(userNode.Property("name").IsEqualTo(test).
			And(userNode.Property("name").IsEqualTo(bazbar).Get()).
			Or(userNode.Property("name").IsEqualTo(foobar).Get()).
			Or(userNode.Property("name").IsEqualTo(foobar).Get())).
		And(userNode.Property("name").IsEqualTo(bazbar).
			And(userNode.Property("name").IsEqualTo(foobar).Get()).
			Or(userNode.Property("name").IsEqualTo(test).Get()).
			Not().Get()).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) WHERE (((u.name = 'Test' AND u.name = 'bazbar') OR u.name = 'foobar' OR u.name = 'foobar') AND NOT (((u.name = 'bazbar' AND u.name = 'foobar') OR u.name = 'Test'))) RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestNegatedCondition(t *testing.T) {
	statement, err := cypher.
		Match(userNode).
		WhereConditionContainer(userNode.Property("name").IsNotNull().Not()).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) WHERE NOT (u.name IS NOT NULL) RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestNoConditionShouldNotBeRendered(t *testing.T) {
	statement, err := cypher.
		Match(userNode).
		Where(cypher.ConditionsNoCondition()).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (u:`User`) RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
	//
	statement, err = cypher.
		Match(userNode).
		WhereConditionContainer(userNode.Property("name").IsEqualTo(cypher.LiteralOf("test"))).
		And(cypher.ConditionsNoCondition().Or(cypher.ConditionsNoCondition()).Get()).
		ReturningByNamed(userNode).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ = cypher.NewRenderer().Render(statement)
	expect = "MATCH (u:`User`) WHERE u.name = 'test' RETURN u"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}
