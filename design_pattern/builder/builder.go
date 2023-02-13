/**
 * @Author: Noaghzil
 * @Date:   2023-02-13 17:11:12
 * @Last Modified by:   Noaghzil
 * @Last Modified time: 2023-02-13 20:47:04
 */
package builder

// Builder 是生成器接口
type Builder[T any] interface {
	Part1()
	Part2()
	Part3()
	GetResult() T
}

type Director[T any] struct {
	builder Builder[T]
}

func (d *Director[T]) Construct() {
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
}

func NewDirector[T any](builder Builder[T]) *Director[T] {
	return &Director[T]{
		builder: builder,
	}
}

type Builder1[T string] struct {
	result T
}

func (b *Builder1[T]) Part1() {
	b.result += "1"
}

func (b *Builder1[T]) Part2() {
	b.result += "2"
}

func (b *Builder1[T]) Part3() {
	b.result += "3"
}

func (b *Builder1[T]) GetResult() T {
	return b.result
}

type Builder2[T int] struct {
	result T
}

func (b *Builder2[T]) Part1() {
	b.result += 1
}

func (b *Builder2[T]) Part2() {
	b.result += 2
}

func (b *Builder2[T]) Part3() {
	b.result += 3
}

func (b *Builder2[T]) GetResult() T {
	return b.result
}
