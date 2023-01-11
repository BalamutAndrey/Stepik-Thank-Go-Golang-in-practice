package main

import (
	"fmt"
)

// element - интерфейс элемента последовательности
// (пустой, потому что элемент может быть любым).
type element interface{}

// iterator - интерфейс, который умеет
// поэлементно перебирать последовательность
type iterator interface {
	next() bool
	val() element
}

// iterate обходит последовательность
// и печатает каждый элемент
func iterate(it iterator) {
	for it.next() {
		curr := it.val()
		fmt.Println(curr)
	}
}
