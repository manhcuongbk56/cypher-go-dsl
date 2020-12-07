package cypher_go_dsl

func Concat(op1 Expression, op2 Expression) Operation {
	return OperationCreate(op1, CONCAT, op2)
}

func Add(op1 Expression, op2 Expression) Operation {
	return OperationCreate(op1, ADDITION, op2)
}

func Subtract(op1 Expression, op2 Expression) Operation {
	return OperationCreate(op1, SUBTRACTION, op2)
}

func multiply(op1 Expression, op2 Expression) Operation {
	return OperationCreate(op1, MULTIPLICATION, op2)
}

func divide(op1 Expression, op2 Expression) Operation {
	return OperationCreate(op1, DIVISION, op2)
}

func remainder(op1 Expression, op2 Expression) Operation {
	return OperationCreate(op1, MODULO_DIVISION, op2)
}

func pow(op1 Expression, op2 Expression) Operation {
	return OperationCreate(op1, EXPONENTIATION, op2)
}

func set(target Expression, value Expression) Operation {
	return OperationCreate(target, SET, value)
}

func set1(target Node, label ...string) Operation {
	return OperationCreate2(target, SET, label...)
}

func remove(target Node, label ...string) Operation {
	return OperationCreate2(target, REMOVE_LABEL, label...)
}
