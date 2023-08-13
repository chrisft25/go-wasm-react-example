package inserter

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"strings"
)

func DrawImages(img1 string, img2 string) string {
	img1Decode, err := FromBase64(img1)

	if err != nil {
		log.Fatal("Error Draw 1", err)
	}

	img2Decode, err := FromBase64(img2)

	if err != nil {
		log.Fatal(err)
	}

	offset := image.Pt(300, 200)

	b := img1Decode.Bounds()
	image3 := image.NewRGBA(b)
	draw.Draw(image3, b, img1Decode, image.Point{}, draw.Src)
	draw.Draw(image3, img2Decode.Bounds().Add(offset), img2Decode, image.Point{}, draw.Over)

	result := new(bytes.Buffer)

	jpeg.Encode(result, image3, &jpeg.Options{Quality: jpeg.DefaultQuality})

	return ToBase64(result.Bytes())
}

func ToBase64(f []byte) string {
	var base64Encoding string

	mimeType := http.DetectContentType(f)

	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	base64Encoding += base64.StdEncoding.EncodeToString(f)

	return base64Encoding
}

func FromBase64(s string) (image.Image, error) {
	mark := ";base64,"
	index := strings.Index(s, mark)

	if index != -1 {
		s = s[index+len(mark):]
	}

	decoded, err := base64.StdEncoding.DecodeString(s)

	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(decoded)
	imagen, _, err := image.Decode(reader)

	if err != nil {
		return png.Decode(reader)
	}

	return imagen, err
}
