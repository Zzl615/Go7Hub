package factory_method


// === 

// Operator 是被封装的实际类接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// OperatorFactory 是工厂接口
type OperatorFactory interface {
	Create() Operator
}



// ===  两数操作赋值 ===

// OperatorBase 是Operator 接口实现的基类，封装公用方法
type OperatorBase struct {
	a, b int
}

// SetA 设置 A
func (o *OperatorBase) SetA(a int) {
	o.a = a
}

// SetB 设置 B
func (o *OperatorBase) SetB(b int) {
	o.b = b
}

// === 加法定义 ===

// PlusOperator Operator 的实际加法实现
type PlusOperator struct {
	*OperatorBase
}

// Result 获取结果
func (o PlusOperator) Result() int {
	return o.a + o.b
}

// PlusOperatorFactory 是 PlusOperator 的工厂类
type PlusOperatorFactory struct{}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

// === 减法定义 ===

// MinusOperator Operator 的实际减法实现
type MinusOperator struct {
	*OperatorBase
}

// Result 获取结果
func (o MinusOperator) Result() int {
	return o.a - o.b
}

type MinusOperatorFactory struct{}

func (MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}




