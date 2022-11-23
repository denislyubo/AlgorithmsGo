package main

import "fmt"

type bar interface {
	Bar() string
}

type barImpl struct {
}

func (bi barImpl) Bar() string {
	return "bar"
}

type foo[B bar] interface {
	Foo(b B)
}

type fooImpl struct {
}

func (fi fooImpl) Foo(b bar) {
	fmt.Println("foo", b.Bar())
}

func buzz(fooStr foo[bar]) {
	fooStr.Foo(barImpl{})
}

func main() {
	fooStr := fooImpl{}

	buzz(fooStr)

}
