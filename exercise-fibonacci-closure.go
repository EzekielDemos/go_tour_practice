/*
From:
https://tour.golang.org/moretypes/26
*/
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var i = 0
	var j = 1
	var result = 0
	return func() int {
		result = i
		i, j = j, i + j
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
