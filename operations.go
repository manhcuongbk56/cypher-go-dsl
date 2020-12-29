package cypher

func OperationConcat(op1 Expression, op2 Expression) Operation {
	return OperationCreate(op1, CONCAT, op2)
}

func OperationAdd(op1 Expression, op2 Expression) Operation {
	return OperationCreate(op1, ADDITION, op2)
}

func OperationSubtract(op1 Expression, op2 Expression) Operation {
	return OperationCreate(op1, SUBTRACTION, op2)
}

func OperationMultiply(op1 Expression, op2 Expression) Operation {
	return OperationCreate(op1, MULTIPLICATION, op2)
}

func OperationDivide(op1 Expression, op2 Expression) Operation {
	return OperationCreate(op1, DIVISION, op2)
}

func OperationRemainder(op1 Expression, op2 Expression) Operation {
	return OperationCreate(op1, MODULO_DIVISION, op2)
}

func OperationPow(op1 Expression, op2 Expression) Operation {
	return OperationCreate(op1, EXPONENTIATION, op2)
}

func OperationSet(target Expression, value Expression) Operation {
	return OperationCreate(target, SET, value)
}

func OperationSetLabel(target Node, label ...string) Operation {
	return OperationCreate2(target, SET_LABEL, label...)
}

func OperationRemove(target Node, label ...string) Operation {
	return OperationCreate2(target, REMOVE_LABEL, label...)
}
