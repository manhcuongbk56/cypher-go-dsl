package expression

type Operator string

const (
	EQUALITY = "equality"
)

type Comparison struct {
	Condition
	Left Expression
	Operator Operator
	Right Expression
}
