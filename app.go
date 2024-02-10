package main

import (
	"fmt"
)

type piece struct {
	position string
	name     string
}

type position struct {
	piece piece
	name  string
}

type board struct {
	spaces [][]int
}

func newBoard(w int, h int) *board {
	s := make([][]int, w)
	for i := range s {
		s[i] = make([]int, h)
	}
	board := board{spaces: s}
	return &board
}

func main() {

	fmt.Println(newBoard(3, 5))
	// fmt.Println(person{"Bob", 20})

	// fmt.Println(person{name: "Alice", age: 30})

	// fmt.Println(person{name: "Fred"})

	// fmt.Println(&person{name: "Ann", age: 40})

	// fmt.Println(newPerson("Jon"))

	// s := person{name: "Sean", age: 50}
	// fmt.Println(s.name)

	// sp := &s
	// fmt.Println(sp.age)

	// sp.age = 51
	// fmt.Println(sp.age)

	// dog := struct {
	//     name   string
	//     isGood bool
	// }{
	//     "Rex",
	//     true,
	// }
	// fmt.Println(dog)
}
