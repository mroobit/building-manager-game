package main

import (
	"embed"
	"fmt"
	"image/color"
	"image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tinne26/etxt"
)

var (
	fonts            *etxt.FontLibrary
	titleBackground  *ebiten.Image
	portalBackground *ebiten.Image
)

func loadAssets() {
	loadImages()
	loadFonts()

	// TODO
	// loadSounds()
}

func loadImages() {
	log.Printf("Loading images...")
	titleBackground = loadImage(FileSystem, "images/title-background.png")
	portalBackground = loadImage(FileSystem, "images/maintenance-portal.png")
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

func loadFonts() {
	fonts = etxt.NewFontLibrary()
	_, _, err := fonts.ParseEmbedDirFonts("fonts", FileSystem)
	if err != nil {
		log.Fatalf("Error while loading fonts: %s", err.Error())
	}
}

func (g *Game) ConfigureTextRenderer() {
	fmt.Println("Configuring text renderer")
	// TODO - create profiles for different text contexts + method to allow quickly reconfiguring based on context
	renderer := etxt.NewStdRenderer()
	cache := etxt.NewDefaultCache(10 * 1024 * 1024)
	renderer.SetCacheHandler(cache.NewHandler())
	renderer.SetFont(fonts.GetFont("Liberation Sans"))
	renderer.SetAlign(etxt.YCenter, etxt.XCenter)
	renderer.SetSizePx(48)

	renderer.SetColor(color.RGBA{239, 91, 91, 255})

	g.Text = renderer
}
