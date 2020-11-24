package cypher_go_dsl

type RelationshipPattern interface {
	ExposesRelationship
	PatternElement
}



/**
RelationshipTypes
 */
type RelationshipTypes struct {
	values []string
}

func (r RelationshipTypes) Accept(visitor Visitor) {
	visitor.Enter(r)
	visitor.Leave(r)
}

func (r RelationshipTypes) GetType() VisitableType {
	return RelationshipTypesVisitable
}

/**
Relationship length
 */

type RelationshipLength struct {
	minimum *int8
	maximum *int8
	unbounded bool
}

func (relationshipLength RelationshipLength)Unbounded() RelationshipLength {
	return RelationshipLength{nil, nil, true}
}




type Direction struct {
	symbolLeft string
	symbolRight string
}

func LTR() Direction {
	return Direction{symbolLeft: "-", symbolRight: "->"}
}

func RTL() Direction {
	return Direction{symbolLeft: "<-", symbolRight: "-"}
}

func UNI() Direction {
	return Direction{symbolLeft: "-", symbolRight: "-"}
}

type Details struct {
	direction *Direction
	symbolicName *SymbolicName
	types *RelationshipTypes
	length *RelationshipLength
	properties *Properties
}

func Create(direction Direction, name SymbolicName, types RelationshipTypes) Details {
	return Details{
		direction: &direction,
		symbolicName: &name,
		types: &types,
	}
}

func CreateLTR(left Node, direction Direction, right Node, types ...string) Relationship {
	typeSlice := make([]string, 0)
	typeSlice = append(typeSlice, types...)
	relationshipTypes := RelationshipTypes{values: typeSlice}
	details := Details{
		direction:    &direction,
		symbolicName: nil,
		types:        &relationshipTypes,
	}
	return Relationship{
		left:    &left,
		right:   &right,
		details: &details,
	}
}


func (detail Details) hasContent() bool {
	return detail.types != nil ||
		detail.direction != nil ||
		detail.symbolicName != nil || 
		detail.properties != nil
}

