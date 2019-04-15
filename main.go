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

	// bounds := img.Bounds()
	// newIMG := image.NewRGBA(bounds)

	// fmt.Println(math.Cos((90*135.089)/180)*(float64(3-1.5)) - math.Sin((90*135.089)/180)*(float64(3-1.5)) + float64(1.5))
	// // Loop in all reference and new image pixels
	// for y := 0; y < bounds.Max.Y; y++ {
	// 	for x := 0; x < bounds.Max.X; x++ {
	// 		actualPixel := img.At(x, y)
	// 		var anglen float64 = 270

	// 		angle := (anglen * 135.089) / 180
	// 		newIMG.Set(int(math.Cos(angle)*(float64(x-bounds.Max.X/2))-math.Sin(angle)*(float64(y-bounds.Max.Y/2))+float64(bounds.Max.X/2)),
	// 			int(math.Sin(angle)*(float64(x-bounds.Max.X/2))+math.Cos(angle)*(float64(y-bounds.Max.Y/2))+float64(bounds.Max.Y/2)), actualPixel)

	// 	}
	// }

	newIMG := tools.Rotate(img, 90)

	// fmt.Println((float64(300) * math.Cos(0.1)) - (float64(300) * math.Sin(0.1)))
	// fmt.Println((float64(300) * math.Sin(0.1)) + (float64(300) * math.Cos(0.1)))
	// flippedImg := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{10.10, 1600}})

	outFile, err := os.Create("changed.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	jpeg.Encode(outFile, newIMG, nil)

}
