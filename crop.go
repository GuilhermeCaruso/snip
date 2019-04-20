package snip

import (
	"image"
)

/*
FieldMask is the struct to define a mask used on
crop method
*/
type FieldMask struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

/*
Crop is responsible for cutting the selected region of an image.
*/
func Crop(img image.Image, field Field, c *chan ExtractedIMG) {
	defer wg.Done()

	mask := field.Mask

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
	*c <- ExtractedIMG{
		Image: fieldIMG,
		Field: field,
	}

}
