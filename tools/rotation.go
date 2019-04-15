package tools

import (
	"image"
	"math"
)

// Rotate Method responsible for free image rotation
func Rotate(img image.Image, angle float64) image.Image {
	var newIMG *image.RGBA

	bounds := img.Bounds()

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
		int(cos*(float64(x-bounds.Max.X/2))-
			sen*(float64(y-bounds.Max.Y/2))+
			float64(bounds.Max.X/2)),
		int(sen*(float64(x-bounds.Max.X/2))+
			cos*(float64(y-bounds.Max.Y/2))+
			float64(bounds.Max.Y/2)),
		actualPixel)

}

func calculateFix(bounds image.Rectangle, x, y int, sen, cos float64, newIMG *image.RGBA, img image.Image) {
	actualPixel := img.At(x, y)
	newIMG.Set(
		int(cos*(float64(x-bounds.Max.X/2))-
			sen*(float64(y-bounds.Max.Y/2))+
			float64(bounds.Max.X/2)-
			float64(bounds.Max.X/2-bounds.Max.Y/2)),
		int(sen*(float64(x-bounds.Max.X/2))+
			cos*(float64(y-bounds.Max.Y/2))+
			float64(bounds.Max.Y/2)+
			float64(bounds.Max.X/2-bounds.Max.Y/2)),
		actualPixel)
}
