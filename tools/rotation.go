package tools

import (
	"image"
	"math"
)

// Rotate Method responsible for free image rotation
func Rotate(img image.Image, angle float64) image.Image {
	// Final image set on retunr
	var newIMG *image.RGBA
	// Get img bounds
	bounds := img.Bounds()
	// Calculate sin and cos using angle parameter
	sen, cos := math.Sincos(math.Pi * angle / 180)

	if angle == 90 || angle == 270 ||
		angle == -90 || angle == -270 {
		newIMG = image.NewRGBA(image.Rectangle{
			Min: image.Point{
				X: 0,
				Y: 0,
			},
			Max: image.Point{
				X: bounds.Max.Y,
				Y: bounds.Max.X,
			},
		})
		for y := 0; y < bounds.Max.Y; y++ {
			for x := 0; x < bounds.Max.X; x++ {
				calculateFix(bounds, x, y, sen, cos, newIMG, img)
			}
		}
	} else {
		newIMG = image.NewRGBA(bounds)
		for y := 0; y < bounds.Max.Y; y++ {
			for x := 0; x < bounds.Max.X; x++ {
				calculate(bounds, x, y, sen, cos, newIMG, img)
			}
		}
	}
	return newIMG

}

// Method responsible for distributing pixels in new positions
func calculate(bounds image.Rectangle, x, y int, sen, cos float64, newIMG *image.RGBA, img image.Image) {
	actualPixel := img.At(x, y)
	newIMG.Set(
		int(cos*(delta(x, bounds.Max.X))-
			sen*(delta(y, bounds.Max.Y))+
			float64(bounds.Max.X/2)),
		int(sen*(delta(x, bounds.Max.X))+
			cos*(delta(y, bounds.Max.Y))+
			float64(bounds.Max.Y/2)),
		actualPixel)

}

// Method responsible for distributing pixels in new fix position
func calculateFix(bounds image.Rectangle, x, y int, sen, cos float64, newIMG *image.RGBA, img image.Image) {
	actualPixel := img.At(x, y)
	newIMG.Set(
		int(cos*(delta(x, bounds.Max.X))-
			sen*(delta(y, bounds.Max.Y))+
			float64(bounds.Max.X/2)-
			float64(bounds.Max.X/2-bounds.Max.Y/2)),
		int(sen*(delta(x, bounds.Max.X))+
			cos*(delta(y, bounds.Max.Y))+
			float64(bounds.Max.Y/2)+
			float64(bounds.Max.X/2-bounds.Max.Y/2)),
		actualPixel)
}

// Calculate delta value
func delta(x, y int) float64 {
	return float64(x - y/2)
}
