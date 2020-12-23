package cypher

type RelationshipPattern interface {
	ExposesRelationshipChain
	PatternElement
	//User for relationship
	Named(name string) Relationship
	//Use for relationship chain
	NamedC(name string) RelationshipChain
}

type Direction struct {
	symbolLeft  string
	symbolRight string
	notNil      bool
}

func DirectionCreate(symbolLeft string, symbolRight string) Direction {
	return Direction{symbolLeft, symbolRight, true}
}

func LTR() Direction {
	return DirectionCreate("-", "->")
}

func RTL() Direction {
	return DirectionCreate("<-", "-")
}

func UNI() Direction {
	return DirectionCreate("-", "-")
}
