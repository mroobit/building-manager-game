package main

import (
	"embed"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var storyText string

func (g *Game) IntroStory(screen *ebiten.Image) {
	// TODO: display Aunt Jo's story
	// TODO: display building hand-over "paperwork" declaring Number of Tenants, etc
	// TODO: display "Skip" button

	vector.DrawFilledRect(screen, 0, 0, 1280.0, 980.0, color.RGBA{25, 15, 15, 255}, false)
	vector.DrawFilledRect(screen, 40, 40, 1200.0, 880.0, color.RGBA{255, 255, 255, 255}, false)

	g.SetTextProfile(textProfile["aunt-jos-letter"])
	g.Text.SetTarget(screen)
	//	g.Text.Draw("I hate to be the bearer of bad news, but I, your beloved Aunt Josephine, am dead!", 90, 70)
	g.Text.Draw(storyText, 90, 55)
}

func loadLetter(fs embed.FS) {
	storyBytes, err := fs.ReadFile("data/aunt-jos-letter.txt")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	storyString := string(storyBytes[:])

	storyText = wrapText(storyString, 78)
}
