package main

import "fmt"

func main()  {
	// plus compute
	var plusCompute *plusOperatorFactory
	fmt.Printf("Plus compute result: %d\n", compute(plusCompute, 2, 3))

	var minusCompute *minusOperatorFactory
	fmt.Printf("Minus compute result: %d\n", compute(minusCompute, 5, 2))
}

type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

type OperatorFactory interface {
	Create() Operator
}

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

func compute(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}

/*------------------- plus operator -------------------*/
type plusOperatorFactory struct {}
func (*plusOperatorFactory) Create() Operator {
	return &PlusOperator{OperatorBase: &OperatorBase{}}
}

type PlusOperator struct {
	*OperatorBase
}

func (po *PlusOperator) Result() int {
	return po.a + po.b
}

/*------------------- minus operator -------------------*/
type minusOperatorFactory struct {}
func (*minusOperatorFactory) Create() Operator{
	return &MinusOperator{OperatorBase: &OperatorBase{}}
}

type MinusOperator struct {
	*OperatorBase
}

func (mo *MinusOperator) Result() int {
	return mo.a - mo.b
}
