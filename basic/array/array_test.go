/**
 * @Author: Noaghzil
 * @Date:   2023-02-13 10:50:37
 * @Last Modified by:   Noaghzil
 * @Last Modified time: 2023-02-13 16:12:33
 */
package array

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	array := [5]int{1: 1, 3: 4}
	for i := 0; i < 5; i++ {
		fmt.Printf("one,索引:%d,值:%d\n", i, array[i])
	}
	for i, v := range array {
		fmt.Printf("two,索引:%d,值:%d\n", i, v)
	}
	for k := range array {
		fmt.Printf("two,元素:%d\n", k)
	}
}
