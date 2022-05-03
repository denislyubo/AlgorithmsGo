package main

import "fmt"

type Bar interface {
	Bar() string
}

type BarImpl struct {
}

func (bi BarImpl) Bar() string {
	return "Bar"
}

type Foo[B Bar] interface {
	Foo(b B)
}

type FooImpl struct {
}

func (fi FooImpl) Foo(b Bar) {
	fmt.Println("Foo", b.Bar())
}

func Buzz(foostr Foo[Bar]) {
	foostr.Foo(BarImpl{})
}

func main() {
	foostr := FooImpl{}

	Buzz(foostr)

}
