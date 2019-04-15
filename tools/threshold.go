package tools

import (
	"image"
	"image/color"
)

/*

Threshold is responsible for define a new image using
color based on limiar passing on the function call

*/
func Threshold(img image.Image, limiar uint8) image.Image {
	// Get the bounds of the reference image
	bounds := img.Bounds()
	/*
		Creates a new blank image with the bounds has retrived from the
		reference image

		Ex:

			referenceImage:
			Size: {width: 2000px, height : 2000dpx}
			Bounds: (0,0)-(2000, 2000)

			newImage
			Bouds: (0,0)-(2000,2000)

	*/
	newIMG := image.NewRGBA(bounds)

	// Loop in all reference and new image pixels
	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			// Get a pixel of reference image
			actualPixel := img.At(x, y)
			// Get rgb coloro of the pixel
			r, g, b, _ := actualPixel.RGBA()
			// Get luminosity by a specific calculation
			luminosity :=
				0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			// Set the color of a new pixel using luminosity
			pixel := color.Gray{uint8(luminosity / 256)}

			// With past threshold set the replacement color
			if pixel.Y > limiar {
				newIMG.Set(x, y, color.Gray{255})
			} else {
				newIMG.Set(x, y, color.Gray{0})
			}
		}
	}

	return newIMG
}
