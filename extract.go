package snip

import (
	"bytes"
	"image"
	"image/jpeg"
	"regexp"
	"sync"

	"github.com/otiai10/gosseract"
)

type Field struct {
	Key    string    `json:"key"`
	Mask   FieldMask `json:"mask"`
	Regexp []string  `json:"regexp"`
}

type ExtractedIMG struct {
	Image image.Image
	Field Field
}

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

var limiarSlice = []uint8{
	90,
	34,
	98,
	43,
	106,
	50,
	114,
	58,
	122,
	66,
	130,
	74,
	138,
	82,
}

var wg sync.WaitGroup

func Extract(img image.Image, fields []Field) []KeyValue {

	c := make(chan ExtractedIMG, len(fields))
	for _, field := range fields {
		wg.Add(1)
		go Crop(img, field, &c)
	}
	wg.Wait()
	close(c)

	c2 := make(chan KeyValue, len(fields))
	for field := range c {
		wg.Add(1)
		go extractText(field, &c2)
	}
	wg.Wait()
	close(c2)

	var values []KeyValue
	for keys := range c2 {
		values = append(values, keys)
	}

	return values
}

func extractText(field ExtractedIMG, c *chan KeyValue) {
	client := gosseract.NewClient()
	client.SetLanguage("por")

	defer client.Close()
	defer wg.Done()

	var finalValue string
	found := false

	for x := 0; x < len(limiarSlice) && found == false; x++ {
		buf := new(bytes.Buffer)

		jpeg.Encode(buf, Threshold(field.Image, limiarSlice[x]), nil)
		imgByte := buf.Bytes()

		client.SetImageFromBytes(imgByte)
		text, _ := client.Text()

		if len(field.Field.Regexp) == 0 {
			finalValue = text
			found = true
		} else {
			for y := 0; y < len(field.Field.Regexp) && found == false; y++ {
				rgx := regexp.MustCompile(string(field.Field.Regexp[y]))
				result := rgx.FindString(text)
				if result != "" {
					finalValue = result
					found = true

				}
			}
		}
	}

	*c <- KeyValue{
		Key:   field.Field.Key,
		Value: finalValue,
	}
}
