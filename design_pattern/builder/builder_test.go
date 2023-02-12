/**
 * @Author: Noaghzil
 * @Date:   2023-02-11 22:21:20
 * @Last Modified by:   Noaghzil
 * @Last Modified time: 2023-02-12 09:29:13
 */
package builder

import (
	"fmt"
	"testing"
)

func TestBuilder1(t *testing.T) {
	var builder *Builder1 = &Builder1{}
	var director *Director = NewDirector(builder)
	director.Construct()
	res := builder.GetResult()
	fmt.Println("res: ", res)
	if res != "123" {
		t.Fatalf("Builder1 fail expect 123 acture %s", res)
	}
}

func TestBuilder2(t *testing.T) {
	builder := &Builder2{}
	director := NewDirector(builder)
	director.Construct()
	res := builder.GetResult()
	fmt.Println("res: ", res)
	if res != 6 {
		t.Fatalf("Builder2 fail expect 6 acture %d", res)
	}

}
