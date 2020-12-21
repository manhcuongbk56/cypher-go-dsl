package cypher

type ExposesRelationship interface {
	RelationshipTo(node Node, types ...string) RelationshipPattern
	RelationshipFrom(node Node, types ...string) RelationshipPattern
	RelationshipBetween(node Node, types ...string) RelationshipPattern
}
