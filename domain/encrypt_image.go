package domain

import (
    "bytes"
    "image"
    "image/png"
    "log"
    "os"
)

func EncryptImage(imgFile *os.File) []byte {
	// Decode the image from the file
	img, _, err := image.Decode(imgFile)
	if err != nil {
		log.Fatalf("failed to decode image: %v, hihihi: %v", err, imgFile)
	}

	// Create a buffer to hold the PNG image data
	var buf bytes.Buffer

	// Encode the image to PNG format and write it to the buffer
	err = png.Encode(&buf, img)
	if err != nil {
		log.Fatalf("failed to encode image: %v", err)
	}

	// Return the byte slice from the buffer
	return buf.Bytes()
}
