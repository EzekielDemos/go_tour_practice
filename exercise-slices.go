/*
Exercise: Slices
From: https://tour.golang.org/moretypes/18
*/

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) (a [][]uint8) {
    a = make([][]uint8, dy)

    for i := range a {
        a[i] = make([]uint8, dx)
    }

    for i := range a {
        for j := range a[i] {
			a[i][j] = uint8(i^j)
        }
    }

    return
}

func main() {
    pic.Show(Pic)
}
