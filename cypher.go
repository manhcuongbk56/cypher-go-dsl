package cypher_go_dsl


func NewNode(primaryLabel string) Node  {
	var labels = make([]NodeLabel, 0)
	labels = append(labels, NodeLabel{primaryLabel})
	return Node{
		labels: labels,
	}
}

func NewNodeWithLabels(primaryLabel string, additionalLabel ...string) Node  {
	var labels = make([]NodeLabel, 0)
	labels = append(labels, NodeLabel{primaryLabel})
	for _, label := range additionalLabel{
		labels = append(labels, NodeLabel{label})
	}
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


func escapeName(name string) string{
	return "`" + name + "`"
}