package ui

import (
	"fmt"
	"log"

	"github.com/ncruces/zenity"
)

const defaultPath = ``

func SelectOption() float64 {
	var choose int
	var option float64

	fmt.Println("Select options:\n1. 75% compress\n2. 50% compress\n3. 25% compress")
	fmt.Scanf("%d", &choose)

	switch choose {
	case 1:
		option = 0.25
	case 2:
		option = 0.5
	case 3:
		option = 0.75
	}
	return option
}

func CreateFileName() string {
	var name string
	fmt.Println("Create name for compressed file")
	fmt.Scanln(&name, ".jpg")

	return name

}

func SelectImgFile() (string, error) {

	img, err := zenity.SelectFile(
		zenity.Filename(defaultPath),
		zenity.FileFilters{
			{"Image files", []string{"*.png", "*.gif", "*.ico", "*.jpg", "*.webp"}, true},
		})
	if err != nil {
		log.Fatal(err)
	}

	return img, err

}
