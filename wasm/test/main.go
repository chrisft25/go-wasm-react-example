package main

import (
	"image/jpeg"
	"log"
	"os"
	inserter "pokeball-inserter/lib"
)

func main() {

	image1, err := os.ReadFile("pikachu.jpg")

	if err != nil {
		log.Fatal("Error imagen 1", err)
	}

	imgBase64 := inserter.ToBase64(image1)

	image2, err := os.ReadFile("pokeball.png")

	if err != nil {
		log.Fatal("Error imagen 2", err)
	}

	img2Base64 := inserter.ToBase64(image2)

	result, err := inserter.FromBase64(inserter.DrawImages(imgBase64, img2Base64))

	if err != nil {
		log.Fatal("Error imagen 3", err)
	}

	img3, err := os.Create("result.jpg")
	if err != nil {
		log.Fatalf("failed to create: %s", err)
	}
	jpeg.Encode(img3, result, &jpeg.Options{Quality: jpeg.DefaultQuality})
	defer img3.Close()

}
