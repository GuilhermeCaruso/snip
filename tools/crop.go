package tools

import "image"

/*
FieldMask is the struct to define a mask used on
crop method
*/
type FieldMask struct {
	X      int
	Y      int
	Width  int
	Height int
}

/*
Crop is responsible for cutting the selected region of an image.
*/
func Crop(img image.Image, mask FieldMask) image.Image {

	fieldIMG := image.NewRGBA(image.Rectangle{
		Min: image.Point{
			X: mask.X,
			Y: mask.Y,
		},
		Max: image.Point{
			X: mask.X + mask.Width,
			Y: mask.Y + mask.Height,
		},
	})

	for x := mask.X; x < mask.X+mask.Width; x++ {
		for y := mask.Y; y < mask.Y+mask.Height; y++ {
			fieldIMG.Set(x, y, img.At(x, y))
		}
	}

	return fieldIMG
}
