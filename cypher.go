package cypher_go_dsl


func NewNode(primaryLabel string, additionalLabel ...string) Node  {
	var labels = make([]string, 0)
	labels = append(labels, primaryLabel)
	labels = append(labels, additionalLabel...)
	return Node{
		labels: labels,
	}
}
func Matchs(element ...PatternElement ) (OngoingReadingWithoutWhere) {
	return NewDefaultBuilder().Match(element...)
}

func Sort(expression Expression) *SortItem {
	return &SortItem{
		expression: expression,
		direction:  SortDirection{UNDEFINED},
	}
}
