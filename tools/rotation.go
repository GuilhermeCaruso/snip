package tools

import (
	"fmt"
	"image"
	"math"
)

// Rotate Method responsible for free image rotation
func Rotate(img image.Image, angle float64) image.Image {
	fmt.Println(angle)
	if angle == 90 {
		return Rotate90(img)
	} else if angle == 180 {
		return Rotate180(img)
	} else if angle == 270 {
		return Rotate270(img)
	} else {
		bounds := img.Bounds()
		newIMG := image.NewRGBA(bounds)

		calculatedAngle := calculateAngle(angle)
		for y := 0; y < bounds.Max.Y; y++ {
			for x := 0; x < bounds.Max.X; x++ {
				calculate(bounds, x, y, calculatedAngle, newIMG, img)
			}
		}
		return newIMG

	}
}

// Rotate90 is responsible for rotation image by 90 degree
func Rotate90(img image.Image) image.Image {
	bounds := img.Bounds()
	newIMG := image.NewRGBA(image.Rectangle{
		image.Point{
			X: 0,
			Y: 0,
		},
		image.Point{
			X: bounds.Max.Y,
			Y: bounds.Max.X,
		},
	})

	angle := calculateAngle(90)
	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			calculate90(bounds, x, y, angle, newIMG, img)
		}
	}
	return newIMG
}

// Rotate180 is responsible for rotation image by 180 degree
func Rotate180(img image.Image) image.Image {
	bounds := img.Bounds()
	newIMG := image.NewRGBA(bounds)

	angle := calculateAngle(180)
	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			calculate(bounds, x, y, angle, newIMG, img)
		}
	}
	return newIMG
}

// Rotate270 is responsible for rotation image by 270 degree
func Rotate270(img image.Image) image.Image {
	bounds := img.Bounds()
	newIMG := image.NewRGBA(image.Rectangle{
		image.Point{
			X: 0,
			Y: 0,
		},
		image.Point{
			X: bounds.Max.Y,
			Y: bounds.Max.X,
		},
	})

	angle := calculateAngle(270)
	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			calculate90(bounds, x, y, angle, newIMG, img)
		}
	}
	return newIMG
}

// Method responsible for distributing pixels in new positions
func calculate(bounds image.Rectangle, x, y int, calculatedAngle float64, newIMG *image.RGBA, img image.Image) {
	actualPixel := img.At(x, y)
	newIMG.Set(
		int(math.Cos(calculatedAngle)*(float64(x-bounds.Max.X/2))-
			math.Sin(calculatedAngle)*(float64(y-bounds.Max.Y/2))+
			float64(bounds.Max.X/2)),
		int(math.Sin(calculatedAngle)*(float64(x-bounds.Max.X/2))+
			math.Cos(calculatedAngle)*(float64(y-bounds.Max.Y/2))+
			float64(bounds.Max.Y/2)),
		actualPixel)
}

func calculate90(bounds image.Rectangle, x, y int, calculatedAngle float64, newIMG *image.RGBA, img image.Image) {
	actualPixel := img.At(x, y)
	newIMG.Set(
		int(math.Cos(calculatedAngle)*(float64(x-bounds.Max.X/2))-
			math.Sin(calculatedAngle)*(float64(y-bounds.Max.Y/2))+
			float64(bounds.Max.X/2)-
			float64(bounds.Max.X/2-bounds.Max.Y/2)),
		int(math.Sin(calculatedAngle)*(float64(x-bounds.Max.X/2))+
			math.Cos(calculatedAngle)*(float64(y-bounds.Max.Y/2))+
			float64(bounds.Max.Y/2)+
			float64(bounds.Max.X/2-bounds.Max.Y/2)),
		actualPixel)
}

func calculateAngle(angle float64) float64 {
	return (angle * 135.089) / 180
}
