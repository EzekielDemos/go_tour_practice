/*
Exercise: Images
From: https://tour.golang.org/methods/25
*/

package main

import (
	"image"
	"image/color"
	"golang.org/x/tour/pic"
	"math"
)

type Image struct{
	w int
	h int
}

func (i Image) Bounds() image.Rectangle{
	return image.Rect(0, 0, i.w, i.h)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
	// v := uint8(x^y)
	v := uint8(float64(x) * math.Log(float64(y)))
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{300, 300}
	pic.ShowImage(m)
}
