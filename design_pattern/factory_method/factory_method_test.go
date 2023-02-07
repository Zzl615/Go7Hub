/**
 * @Author: Noaghzil
 * @Date:   2023-02-07 08:08:41
 * @Last Modified by:   Noaghzil
 * @Last Modified time: 2023-02-07 08:16:00
 */
package factory_method

import (
	"fmt"
	"testing"
)

func TestMinusOperator(t *testing.T) {
	fmt.Println("TestMinusOperator: ")

	var factory OperatorFactory = MinusOperatorFactory{}
	var operator Operator = factory.Create()
	operator.SetA(1)
	operator.SetB(2)
	if operator.Result() != -1 {
		t.Fatal("error with MinusOperatorFactory")
	}

}

// func TestPlusOperator() {

// 	fmt.Println("TestPlusOperator: ")

// }
