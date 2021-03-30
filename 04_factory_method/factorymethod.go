package factorymethod

//Operator
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

//OperatorFactory
type OperatorFactory interface {
	Create() Operator
}

//OperatorBase
type OperatorBase struct {
	a, b int
}

//SetA
func (o *OperatorBase) SetA(a int) {
	o.a = a
}

//SetB
func (o *OperatorBase) SetB(b int) {
	o.b = b
}

//PlusOperatorFactory
type PlusOperatorFactory struct{}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

//PlusOperator Operator
type PlusOperator struct {
	*OperatorBase
}

//Result
func (o PlusOperator) Result() int {
	return o.a + o.b
}

//MinusOperatorFactory
type MinusOperatorFactory struct{}

func (MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}

//MinusOperator Operator
type MinusOperator struct {
	*OperatorBase
}

//Result
func (o MinusOperator) Result() int {
	return o.a - o.b
}
