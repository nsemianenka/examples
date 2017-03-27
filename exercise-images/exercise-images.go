package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	w, h int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	left, right := image.Point{0, 0}, image.Point{img.w, img.h}
	return image.Rectangle{left, right}
}

func (img Image) At(x, y int) color.Color {
	v := uint8((x ^ y) / 2)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{500, 500}
	pic.ShowImage(m)
}
