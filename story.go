package main

import (
	"embed"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	storyText string

	letterBackground = color.RGBA{25, 15, 15, 255}
)

func (g *Game) IntroStory(screen *ebiten.Image) {
	// TODO: display building hand-over "paperwork" declaring Number of Tenants, etc

	vector.DrawFilledRect(screen, 0, 0, 1280.0, 980.0, letterBackground, false)
	vector.DrawFilledRect(screen, 40, 40, 1200.0, 880.0, white, false)

	g.SetTextProfile(textProfile["aunt-jos-letter"])
	g.Text.SetTarget(screen)
	g.Text.Draw(storyText, 90, 55)

	g.DrawContinueButton(screen, "continue")
}

func loadLetter(fs embed.FS) {
	storyBytes, err := fs.ReadFile("data/aunt-jos-letter.txt")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	storyString := string(storyBytes[:])

	storyText = wrapText(storyString, 78)
}
