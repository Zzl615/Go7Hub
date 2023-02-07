/**
 * @Author: Noaghzil
 * @Date:   2023-02-07 08:08:41
 * @Last Modified by:   Noaghzil
 * @Last Modified time: 2023-02-07 08:19:46
 */
package factory_method

import (
	"fmt"
	"testing"
)

func compute(factory OperatorFactory, a int, b int) int {
	var operator Operator = factory.Create()
	operator.SetA(1)
	operator.SetB(2)
	return operator.Result()
}

func TestMinusOperator(t *testing.T) {
	var factory OperatorFactory = MinusOperatorFactory{}
	var result int = compute(factory, 1, 2)
	if compute(factory, 1, 2) != -1 {
		t.Fatal("error with MinusOperatorFactory")
	}
	fmt.Println("Result of TestMinusOperator: ", result)
}

// func TestPlusOperator() {

// 	fmt.Println("TestPlusOperator: ")

// }
