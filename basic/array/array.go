/**
 * @Author: Noaghzil
 * @Date:   2023-02-01 08:17:05
 * @Last Modified by:   Noaghzil
 * @Last Modified time: 2023-02-13 10:50:50
 */
package array

import (
	"fmt"
)

func main() {
	// 数组赋值
	array := [5]int{1: 1, 3: 4}
	for i := 0; i < 5; i++ {
		fmt.Printf("one,索引:%d,值:%d\n", i, array[i])
	}
	for i, v := range array {
		fmt.Printf("two,索引:%d,值:%d\n", i, v)
	}
	// 数组指针
	// 只可以给索引1和3赋值，因为只有它们分配了内存，才可以赋值
	array2 := [5]*int{1: new(int), 3: new(int)}
	*array2[1] = 1
	fmt.Printf("array2[1]:%d\n", *array2[1])
	array2[0] = new(int)
	*array2[0] = 2
	fmt.Printf("array2[0]:%d\n", *array2[0])
}
