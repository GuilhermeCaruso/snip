package tools

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
)

/*
Decode responsible for decoding and instantiating image based on
a content-type.
*/
func Decode(img io.Reader) (image.Image, error) {
	var imageReference image.Image

	contentType, err := contentType(img)
	if err != nil {
		return nil, err
	}
	switch contentType {
	case "image/jpeg":
		imageReference, err = jpeg.Decode(img)
	case "image/png":
		imageReference, err = png.Decode(img)
	}
	return imageReference, nil
}

//Responsible for get image content-type
func contentType(img io.Reader) (string, error) {

	buffer := make([]byte, 512)
	_, err := img.Read(buffer)

	if err != nil {
		return "", err
	}
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
