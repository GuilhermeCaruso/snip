package main

import (
	"github.com/GuilhermeCaruso/snip/tools"
	//"fmt"

	"image/jpeg"
	"log"
	"os"
)

func main() {
	file, err := os.Open("jpg.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(os.Stderr, "%s: %v\n", "flower.jpg", err)
	}

	usedImg := tools.Threshold(img, 90)

	outFile, err := os.Create("changed.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	jpeg.Encode(outFile, usedImg, nil)

}
