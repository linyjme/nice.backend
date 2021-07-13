package interfaces

//Operator 是被封装的实际类接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}


//OperatorFactory 是工厂接口
type OperatorFactory interface {
	Create() Operator
}
