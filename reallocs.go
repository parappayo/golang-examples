//
//  go run reallocs.go
//
//  Shows the new capacity and memory address of a slice that gets reallocated.
//  The Go runtime copies all data to the new slice on realloc.
//

package main

import "fmt"

func appendMultiple(src []int, n int) []int {
	result := make([]int, len(src))
	var newResult []int
	copy(result, src)
	for i := 0; i < n; i++ {
		newResult = append(result, i)
		if &newResult[0] != &result[0] {
			fmt.Printf("reallocation: capacity %d address %p\n", cap(result), result)
		}
		result = newResult
	}
	return result
}

func main() {
	foo := make([]int, 1)
	foo = appendMultiple(foo, 40)
}
