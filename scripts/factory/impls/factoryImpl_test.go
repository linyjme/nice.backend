package impls

import "fmt"
import "testing"

func compute(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}

func TestOperator(t *testing.T) {
	var factory OperatorFactory
	factory = PlusOperatorFactory{}
	result1 := compute(factory, 1, 2)
	fmt.Println(result1)

	factory = MinusOperatorFactory{}
	result2 := compute(factory, 4, 2)
	fmt.Println(result2)
}
