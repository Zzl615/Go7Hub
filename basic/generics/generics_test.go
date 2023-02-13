/**
 * @Author: Noaghzil
 * @Date:   2023-02-13 14:52:54
 * @Last Modified by:   Noaghzil
 * @Last Modified time: 2023-02-13 16:07:11
 */
package generics

import (
	"fmt"
	"testing"
)

func TestGenericsMapKeys(t *testing.T) {
	var m = map[int32]string{1: "2", 2: "4", 4: "8"}

	result := MapKeys(m)

	fmt.Println("keys:", result)

	fmt.Println("keys:", result[0])

	if result[0] != 1 {
		t.Fatal("error with TestGenericsMapKeys")
	}
}

func TestGenericsList(t *testing.T) {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	result := lst.GetAll()
	fmt.Println("list:", result)
	if result[0] != 10 {
		t.Fatal("error with TestGenericsList")
	}
}
