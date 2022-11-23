package main

import "fmt"

type Prettier interface {
	Pretty() string
}

func prettyPrint[T Prettier](items []T) {
	for _, item := range items {
		fmt.Println(item.Pretty())
	}
}

type animal int

func (a animal) Pretty() string {
	return fmt.Sprintf("Pretty %d", a)
}

func main() {
	m := make([]animal, 0, 3)

	m = append(m, animal(1))
	m = append(m, animal(2))
	m = append(m, animal(3))

	prettyPrint[animal](m)

}
