package cypherit

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestDoc3651And(t *testing.T) {
	timothy := cypher.NewNode("Person").NamedByString("timothy").
		WithRawProperties("name", cypher.LiteralOf("Timothy"))
	other := cypher.NewNode("Person").NamedByString("other")
	statement, err := cypher.MatchElements(timothy, other).
		WhereConditionContainer(other.Property("name").In(cypher.ListOf(cypher.LiteralOf("Andy"), cypher.LiteralOf("Peter")))).
		AndPattern(timothy.RelationshipFrom(other)).
		Returning(other.Property("name"), other.Property("age")).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (timothy:`Person` {name: 'Timothy'}), (other:`Person`) WHERE (other.name IN ['Andy', 'Peter'] AND (timothy)<--(other)) RETURN other.name, other.age"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
	//
	statement, err = cypher.MatchElements(timothy, other).
		WhereConditionContainer(other.Property("name").In(cypher.ListOf(cypher.LiteralOf("Andy"), cypher.LiteralOf("Peter"))).
			AndPattern(timothy.RelationshipFrom(other))).
		Returning(other.Property("name"), other.Property("age")).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ = cypher.NewRenderer().Render(statement)
	expect = "MATCH (timothy:`Person` {name: 'Timothy'}), (other:`Person`) WHERE (other.name IN ['Andy', 'Peter'] AND (timothy)<--(other)) RETURN other.name, other.age"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestDoc3651Or(t *testing.T) {
	timothy := cypher.NewNode("Person").NamedByString("timothy").
		WithRawProperties("name", cypher.LiteralOf("Timothy"))
	other := cypher.NewNode("Person").NamedByString("other")
	statement, err := cypher.MatchElements(timothy, other).
		WhereConditionContainer(other.Property("name").In(cypher.ListOf(cypher.LiteralOf("Andy"), cypher.LiteralOf("Peter")))).
		OrPattern(timothy.RelationshipFrom(other)).
		Returning(other.Property("name"), other.Property("age")).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (timothy:`Person` {name: 'Timothy'}), (other:`Person`) WHERE (other.name IN ['Andy', 'Peter'] OR (timothy)<--(other)) RETURN other.name, other.age"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
	//
	statement, err = cypher.MatchElements(timothy, other).
		WhereConditionContainer(other.Property("name").In(cypher.ListOf(cypher.LiteralOf("Andy"), cypher.LiteralOf("Peter"))).
			OrPattern(timothy.RelationshipFrom(other))).
		Returning(other.Property("name"), other.Property("age")).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ = cypher.NewRenderer().Render(statement)
	expect = "MATCH (timothy:`Person` {name: 'Timothy'}), (other:`Person`) WHERE (other.name IN ['Andy', 'Peter'] OR (timothy)<--(other)) RETURN other.name, other.age"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestDoc3651Xor(t *testing.T) {
	timothy := cypher.NewNode("Person").NamedByString("timothy").
		WithRawProperties("name", cypher.LiteralOf("Timothy"))
	other := cypher.NewNode("Person").NamedByString("other")
	statement, err := cypher.MatchElements(timothy, other).
		WhereConditionContainer(other.Property("name").In(cypher.ListOf(cypher.LiteralOf("Andy"), cypher.LiteralOf("Peter"))).
			XorPattern(timothy.RelationshipFrom(other))).
		Returning(other.Property("name"), other.Property("age")).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (timothy:`Person` {name: 'Timothy'}), (other:`Person`) WHERE (other.name IN ['Andy', 'Peter'] XOR (timothy)<--(other)) RETURN other.name, other.age"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestDoc3652(t *testing.T) {
	person := cypher.NewNode("Person").NamedByString("person")
	peter := cypher.NewNode("Person").NamedByString("peter").WithRawProperties("name", cypher.LiteralOf("Peter"))
	statement, err := cypher.MatchElements(person, peter).
		Where(cypher.ConditionsNotByPattern(person.RelationshipTo(peter))).
		Returning(person.Property("name"), person.Property("age")).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (person:`Person`), (peter:`Person` {name: 'Peter'}) WHERE NOT (person)-->(peter) RETURN person.name, person.age"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestDoc3653(t *testing.T) {
	person := cypher.NewNode("Person").NamedByString("n")
	statement, err := cypher.MatchElements(person).
		WherePattern(person.RelationshipBetween(cypher.AnyNode().WithRawProperties("name", cypher.LiteralOf("Timothy")), "KNOWS")).
		Returning(person.Property("name"), person.Property("age")).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (n:`Person`) WHERE (n)-[:`KNOWS`]-( {name: 'Timothy'}) RETURN n.name, n.age"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestGh113(t *testing.T) {
	foo := cypher.NewNode("Foo").NamedByString("foo")
	bar := cypher.NewNode("Bar").NamedByString("bar")
	fooBar := foo.RelationshipTo(bar, "FOOBAR").NamedByString("rel")
	pc := cypher.ListBasedOn(fooBar).
		WherePattern(bar.RelationshipTo(cypher.NewNode("ZZZ").NamedByString("zzz"), "HAS")).
		ReturningByNamed(fooBar, bar)

	statement, err := cypher.MatchElements(foo).
		Returning(foo.GetSymbolicName(), pc).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (foo:`Foo`) RETURN foo, [(foo)-[rel:`FOOBAR`]->(bar:`Bar`) WHERE (bar)-[:`HAS`]->(zzz:`ZZZ`) | [rel, bar]]"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestDoc3654(t *testing.T) {
	person := cypher.NewNode("Person").NamedByString("n")
	pathPattern := person.RelationshipTo(cypher.AnyNode()).NamedByString("r")
	statement, err := cypher.MatchElements(pathPattern).
		WhereConditionContainer(person.Property("name").IsEqualTo(cypher.LiteralOf("Andy"))).
		And(cypher.FunctionType(pathPattern).MatchesPattern("K.*").Get()).
		Returning(cypher.FunctionType(pathPattern), pathPattern.Property("since")).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (n:`Person`)-[r]->() WHERE (n.name = 'Andy' AND type(r) =~ 'K.*') RETURN type(r), r.since"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestAfterWith(t *testing.T) {
	timothy := cypher.NewNode("Person").NamedByString("timothy").
		WithRawProperties("name", cypher.LiteralOf("Timothy"))
	other := cypher.NewNode("Person").NamedByString("other")
	statement, err := cypher.MatchElements(timothy, other).
		WithByNamed(timothy, other).
		Where(other.Property("name").In(cypher.ListOf(cypher.LiteralOf("Andy"), cypher.LiteralOf("Peter"))).Get()).
		AndPattern(timothy.RelationshipFrom(other)).
		Returning(other.Property("name"), other.Property("age")).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (timothy:`Person` {name: 'Timothy'}), (other:`Person`) WITH timothy, other WHERE (other.name IN ['Andy', 'Peter'] AND (timothy)<--(other)) RETURN other.name, other.age"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
	//
	statement, err = cypher.MatchElements(timothy, other).
		WithByNamed(timothy, other).
		Where(other.Property("name").In(cypher.ListOf(cypher.LiteralOf("Andy"), cypher.LiteralOf("Peter"))).AndPattern(timothy.RelationshipFrom(other)).Get()).
		Returning(other.Property("name"), other.Property("age")).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ = cypher.NewRenderer().Render(statement)
	expect = "MATCH (timothy:`Person` {name: 'Timothy'}), (other:`Person`) WITH timothy, other WHERE (other.name IN ['Andy', 'Peter'] AND (timothy)<--(other)) RETURN other.name, other.age"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}

func TestInPatternComprehensions(t *testing.T) {
	a := cypher.NewNode("Person").WithRawProperties("name", cypher.LiteralOf("Keanu Reeves")).NamedByString("a")
	b := cypher.AnyNodeNamed("b")
	statement, err := cypher.MatchElements(a).
		Returning(
			cypher.ListBasedOn(a.RelationshipBetween(b)).
				Where(b.HasLabels("Movie").And(b.Property("released").IsNotNull().Get()).Get()).
				Returning(b.Property("released")).
				As("years").Get()).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ := cypher.NewRenderer().Render(statement)
	expect := "MATCH (a:`Person` {name: 'Keanu Reeves'}) RETURN [(a)--(b) WHERE (b:`Movie` AND b.released IS NOT NULL) | b.released] AS years"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
	//
	statement, err = cypher.MatchElements(a).
		Returning(
			cypher.ListBasedOn(a.RelationshipBetween(b)).
				Where(b.HasLabels("Movie").
					And(b.Property("released").IsNotNull().Get()).
					Or(b.Property("title").IsEqualTo(cypher.LiteralOf("The Matrix")).Get()).
					Or(b.Property("title").IsEqualTo(cypher.LiteralOf("The Matrix 2")).Get()).Get()).
				Returning(b.Property("released")).
				As("years").Get()).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ = cypher.NewRenderer().Render(statement)
	expect = "MATCH (a:`Person` {name: 'Keanu Reeves'}) RETURN [(a)--(b) WHERE ((b:`Movie` AND b.released IS NOT NULL) OR b.title = 'The Matrix' OR b.title = 'The Matrix 2') | b.released] AS years"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
	//
	statement, err = cypher.MatchElements(a).
		Returning(
			cypher.ListBasedOn(a.RelationshipBetween(b)).
				Where(b.HasLabels("Movie")).
				And(b.Property("released").IsNotNull().Get()).
				Or(b.Property("title").IsEqualTo(cypher.LiteralOf("The Matrix")).Get()).
				Or(b.Property("title").IsEqualTo(cypher.LiteralOf("The Matrix 2")).Get()).
				Returning(b.Property("released")).
				As("years").Get()).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ = cypher.NewRenderer().Render(statement)
	expect = "MATCH (a:`Person` {name: 'Keanu Reeves'}) RETURN [(a)--(b) WHERE ((b:`Movie` AND b.released IS NOT NULL) OR b.title = 'The Matrix' OR b.title = 'The Matrix 2') | b.released] AS years"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
	//
	statement, err = cypher.MatchElements(a).
		Returning(
			cypher.ListBasedOn(a.RelationshipBetween(b)).
				Where(b.HasLabels("Movie")).
				Returning(b.Property("released")).
				As("years").Get()).
		Build()
	if err != nil {
		t.Errorf("error when build query\n %s", err)
		return
	}
	query, _ = cypher.NewRenderer().Render(statement)
	expect = "MATCH (a:`Person` {name: 'Keanu Reeves'}) RETURN [(a)--(b) WHERE b:`Movie` | b.released] AS years"
	if query != expect {
		t.Errorf("\n%s is incorrect, expect is \n%s", query, expect)
	}
}
