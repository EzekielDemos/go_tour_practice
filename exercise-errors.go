/*
Exercise: Errors
From: https://tour.golang.org/methods/20
*/

package main

import (
	"fmt"
)

// TODO: ...

func Sqrt(x float64) (float64, error) {
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
