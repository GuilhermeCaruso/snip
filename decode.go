package snip

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

/*
Decode responsible for decoding and instantiating image based on
a content-type.
*/
func Decode(image io.Reader, contentType string) (imageRef image.Image, err error) {
	switch contentType {
	case "image/jpeg":
		imageRef, err = jpeg.Decode(image)
	case "image/jpg":
		imageRef, err = jpeg.Decode(image)
	case "image/png":
		imageRef, err = png.Decode(image)
	}
	return
}
