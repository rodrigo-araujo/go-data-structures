package main

import (
	"fmt"
)

type Hashset struct {
	items map[any]struct{}
}

func (hashset *Hashset) Add(items ...any) {
	for _, item := range items {
		hashset.items[item] = struct{}{}
	}
}

func (hashset *Hashset) Contains(item any) bool {
	_, found := hashset.items[item]
	if found {
		return true
	}
	return false
}

func (hashset *Hashset) Remove(item any) {
	delete(hashset.items, item)
}

func (hashset *Hashset) RemoveMany(itens ...any) {
	for _, item := range hashset.items {
		delete(hashset.items, item)
	}
}

func (hashset *Hashset) Size() int {
	return len(hashset.items)
}

func (hashset *Hashset) Empty() bool {
	return len(hashset.items) == 0
}

func (hashset *Hashset) Values() []any {
	values := make([]any, hashset.Size())
	index := 0
	for value := range hashset.items {
		values[index] = value
		index++
	}

	return values
}

func (hashset *Hashset) Clear() {
	hashset.items = make(map[any]struct{})
}

func (hashset *Hashset) Merge(input *Hashset) {
	for _, item := range input.Values() {
		hashset.items[item] = struct{}{}
	}
}

func New(values ...any) *Hashset {
	hashset := &Hashset{items: make(map[any]struct{})}
	if len(values) > 0 {
		hashset.Add(values...)
	}

	return hashset
}

func main() {
	hs := New(1, 4, "abc")
	fmt.Println(hs)

	hs.Add("hello")
	hs.Add(true)
	hs.Add(true)

	fmt.Println(hs.Values())
	fmt.Println("contains 1: ", hs.Contains(1))

	hs.Remove(1)

	fmt.Println(hs.Values())
	fmt.Println("contains 1: ", hs.Contains(1))
	fmt.Println(hs.Values())

	anotherHs := New(false, 11, 13, "world")
	hs.Merge(anotherHs)
	fmt.Println(hs.Values())
}
