package main

import "fmt"

func main() {
	input := []int{11, 9, 7}

	result := mergeSort(input)

	fmt.Println(result)
}

func mergeSort(in []int) []int {
	size := len(in)
	if size == 1 {
		return in
	}

	meio := size / 2
	a := mergeSort(in[0:meio])
	b := mergeSort(in[meio:])

	return merge(a, b)
}

func merge(a, b []int) []int {
	tamanhoFinal := len(a) + len(b)
	result := make([]int, tamanhoFinal)

	i, j := 0, 0
	for k := 0; k < tamanhoFinal; k++ {
		if i >= len(a) {
			result[k] = b[j]
			j++
		} else if j >= len(b) {
			result[k] = a[i]
			i++
		} else if a[i] < b[j] {
			result[k] = a[i]
			i++
		} else {
			result[k] = b[j]
			j++
		}
	}

	return result
}
