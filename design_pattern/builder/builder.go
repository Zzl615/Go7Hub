/**
 * @Author: Noaghzil
 * @Date:   2023-02-11 22:21:13
 * @Last Modified by:   Noaghzil
 * @Last Modified time: 2023-02-12 09:28:57
 */
package builder

// Builder 是生成器接口
type Builder interface {
	Part1()
	Part2()
	Part3()
	// GetResult() interface{}
}

type DirectorBase interface {
}

type Director struct {
	builder Builder
}

func (d *Director) Construct() {
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
}

// func (d *Director) GetResult[T comparable, V int64 | float64]()  T {
// 	return d.builder.GetResult() T
// }

// 同样的地址返回值，返回值定义类型：interface
func NewDirectorBase(builder Builder) DirectorBase {
	return &Director{
		builder: builder,
	}
}

// 同样的地址返回值，返回值定义类型：struct
func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

type Builder1 struct {
	result string
}

func (b *Builder1) Part1() {
	b.result += "1"
}

// Part2 ...

func (b *Builder1) Part2() {
	b.result += "2"
}

// Part3 ...

func (b *Builder1) Part3() {
	b.result += "3"
}

// GetResult ...

func (b *Builder1) GetResult() string {
	return b.result
}

type Builder2 struct {
	result int
}

func (b *Builder2) Part1() {
	b.result += 1
}

// Part2 ...

func (b *Builder2) Part2() {
	b.result += 2
}

// Part3 ...

func (b *Builder2) Part3() {
	b.result += 3
}

// GetResult ...

func (b *Builder2) GetResult() int {
	return b.result
}
