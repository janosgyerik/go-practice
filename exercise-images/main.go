package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	width, height int
	fun func(int, int) color.Color
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (img Image) At(x, y int) color.Color {
	return img.fun(x, y)
}

func grayscale(x, y int) color.Color {
	v := uint8(x + y)
	return color.RGBA{v, v, v, 255}
}

func main() {
	m := Image{100, 100, grayscale}
	pic.ShowImage(m)
}
