package main

import ("golang.org/x/tour/pic"
		"image"
		"image/color"
	   )

type Image struct{
	width, height int
}

func (i Image) At(x,y int) color.Color {
	alg := uint8((x ^ y) / 2)
	
	return color.RGBA{alg, alg, 128, 84}
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0,0, i.width, i.height)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func main() {
	m := Image{128, 128}
	pic.ShowImage(m)
}
