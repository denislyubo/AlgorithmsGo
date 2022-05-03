package main

import "fmt"

type Prettier interface {
	Pretty() string
}

func PrettyPrint[T Prettier](items []T) {
	for _, item := range items {
		fmt.Println(item.Pretty())
	}
}

type Animal int

func (a Animal) Pretty() string {
	return fmt.Sprintf("Pretty %d", a)
}

func main() {
	m := make([]Animal, 0, 3)

	m = append(m, Animal(1))
	m = append(m, Animal(2))
	m = append(m, Animal(3))

	PrettyPrint[Animal](m)

}
