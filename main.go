package main

import (
	"image/jpeg"
	"log"
	"os"

	"github.com/GuilhermeCaruso/snip/tools"
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

	newIMG := tools.Rotate(img, 0)
	newIMG = tools.Threshold(newIMG, 100)

	outFile, err := os.Create("changed.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	jpeg.Encode(outFile, newIMG, nil)

}
