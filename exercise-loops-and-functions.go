/*
Exercise: Loops and Functions
From: https://tour.golang.org/flowcontrol/8
*/

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2*z)
	}

	return z
}

func Sqrt1(x float64) float64 {
	z := 1.0
	t := 0.0
	i := 0
	for {
		i++
        z, t = z - (z*z - x) / (2*z), z
        if math.Abs(t-z) < 1e-6 {
            break
        }
    }
	fmt.Printf("*** Sqrt1(%v) calc times: %v, value= %v\n", x, i, z)
	return z
}

func main() {
	fmt.Println("math.Sqrt(2)=", math.Sqrt(2))
	fmt.Println("Sqrt(2)=", Sqrt(2))
	fmt.Println("Sqrt1(2)=", Sqrt1(2))
}
