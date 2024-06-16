package main

import (
	"embed"
	"image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	titleBackground *ebiten.Image
)

func loadAssets() {
	// TODO
	loadImages()

	// loadSounds()
	// loadFonts()
}

func loadImages() {
	log.Printf("Loading images...")
	titleBackground = loadImage(FileSystem, "images/title-background.png")
}

// This is taken directly from https://github.com/mroobit/untitled-sidescroller/blob/main/helper.go#L120
func loadImage(fs embed.FS, path string) *ebiten.Image {
	rawFile, err := fs.Open(path)
	if err != nil {
		log.Fatalf("Error opening file %s: %v\n", path, err)
	}
	defer rawFile.Close()

	img, err := png.Decode(rawFile)
	if err != nil {
		log.Fatalf("Error decoding file %s: %v\n", path, err)
	}
	loadedImage := ebiten.NewImageFromImage(img)
	return loadedImage
}
