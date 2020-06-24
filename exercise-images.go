/*
Exercise: Images
From: https://tour.golang.org/methods/25
*/

package main

import "golang.org/x/tour/pic"


type Image struct{}

// TODO: ...

func main() {
	m := Image{}
	pic.ShowImage(m)
}
