package main

import (
	"image/jpeg"
	"log"
	"os"

	"github.com/nfnt/resize"

	storage "goresizer.com/m/minio"
	"goresizer.com/m/ui"
)

const defaultPath = ``
const bName = "pic-storage"

func main() {

	img := ui.SelectImgFile()

	file, err := os.Open(img)
	if err != nil {
		log.Fatal(err)
	}

	storage.UploadImgFile(img)

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
