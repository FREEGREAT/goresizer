package main

import (
	"image/jpeg"
	"log"
	"os"

	"github.com/ncruces/zenity"
	"github.com/nfnt/resize"
	"goresizer.com/m/ui"
)

const defaultPath = ``

func main() {

	img, err := zenity.SelectFile(
		zenity.Filename(defaultPath),
		zenity.FileFilters{
			{"Image files", []string{"*.png", "*.gif", "*.ico", "*.jpg", "*.webp"}, true},
		})
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(img)
	if err != nil {
		log.Fatal(err)
	}

	img_decode, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	resizepercent := ui.SelectOption()
	println(img_decode.Bounds().Dx(), " \n", img_decode.Bounds().Dy())

	m := resize.Resize(uint(float64(img_decode.Bounds().Dx())*resizepercent), 0, img_decode, resize.Bicubic)

	name := ui.CreateFileName()

	out, err := os.Create(name + ".jpg")
	if err != nil {
		log.Fatal(err)
	}

	defer out.Close()

	jpeg.Encode(out, m, nil)

}
