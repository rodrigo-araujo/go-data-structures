package main

import "fmt"

type Node[T any] struct {
	data T
	next *Node[T]
}

func main() {
	n1 := Node[string]{
		data: "Hello",
		next: nil,
	}

	n2 := Node[string]{
		data: "world",
		next: nil,
	}

	n1.next = &n2

	fmt.Println(n1.data)
	fmt.Println(n1.next.data)
}
