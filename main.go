package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/GuilhermeCaruso/snip/tools"
)

func main() {
	file, err := os.Open("jpg.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buffer := make([]byte, 512)

	_, err = file.Read(buffer)

	if err != nil {
	}
	tools.Decode(file, "oi")
	contentType := http.DetectContentType(buffer)
	fmt.Println(contentType)

	// img, err := jpeg.Decode(file)
	// if err != nil {
	// 	log.Fatal(os.Stderr, "%s: %v\n", "flower.jpg", err)
	// }

	// newIMG := tools.Rotate(img, 2)
	// newIMG = tools.Crop(newIMG, tools.FieldMask{
	// 	X:      318,
	// 	Y:      140,
	// 	Height: 78,
	// 	Width:  482,
	// })
	// newIMG = tools.Threshold(newIMG, 100)
	// // newIMG = tools.Threshold(newIMG, 100)

	// outFile, err := os.Create("changed.jpg")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer outFile.Close()
	// jpeg.Encode(outFile, newIMG, nil)
	// cmd := exec.Command("tesseract", "changed.jpg", "stdout", "--oem 1", "-l por")
	// log.Printf("Running command and waiting for it to finish...")
	// a, _ := cmd.Output()
	// aalue := (string(a))
	// fmt.Println(aalue + " Resultado")
	// err := cmd.Run()
	// log.Printf("Command finished with error: %v", err)
}
