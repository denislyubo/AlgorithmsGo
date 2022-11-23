package main

import "fmt"

type Bar interface {
	Bar() string
}

type barImpl struct {
}

func (bi barImpl) Bar() string {
	return "Bar"
}

type foo[B Bar] interface {
	Foo(b B)
}

type fooImpl struct {
}

func (fi fooImpl) Foo(b Bar) {
	fmt.Println("foo", b.Bar())
}

func Buzz(foostr foo[Bar]) {
	foostr.Foo(barImpl{})
}

func main() {
	foostr := fooImpl{}

	Buzz(foostr)

}
