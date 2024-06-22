package main

import (
	"embed"
	"fmt"
	"image/color"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/tinne26/etxt"
)

var (
	hover        = ""
	cursor       [2]int
	rentIncrease = 3 // percentage increase at renewal

	//go:embed data
	//go:embed fonts
	//go:embed images
	FileSystem embed.FS
)

func main() {
	gameWidth, gameHeight := 1280, 960

	ebiten.SetWindowSize(gameWidth, gameHeight)
	ebiten.SetWindowTitle("Building Manager")

	loadAssets()

	game := &Game{
		Width:  gameWidth,
		Height: gameHeight,
		State:  "load",
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	Width       int
	Height      int
	State       string
	Page        string
	Building    *Building
	RequestPool []*Request
	Text        *etxt.Renderer
	Clock       *Clock
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return g.Width, g.Height
}

func (g *Game) Update() error {
	cursor[0], cursor[1] = ebiten.CursorPosition()

	if g.State == "load" {
		initializeClickables()
		initializeTenants(tenants)
		g.initializeClock()
		g.initializeBuilding()
		g.initializeRequestPool(FileSystem)
		g.ConfigureTextRenderer()
		loadLetter(FileSystem)
		if len(background) == 2 {
			g.State = "title"
		}
	} else if g.State == "story" {
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.State = "play"
		}

	} else if g.State == "play" {
		// TODO: Make incrementing of months a function of tasks done(weight) + ticks
		g.Clock.Tick += 1
		g.CheckDaysInMonth()
		// TODO: generate problems based on Tick/Day + some randomness
		g.CreateProblems()

		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.GenerateRequest()
		}

		switch {
		case portalClickable["overview"].Hover(cursor):
			hover = "overview"
		case portalClickable["request-list"].Hover(cursor):
			hover = "request-list"
		case portalClickable["financial-overview"].Hover(cursor):
			hover = "financial-overview"
		default:
			hover = ""
		}
		if g.Page == "request-list" {
			if portalClickable["request-details"].Hover(cursor) {
				hover = "request-details"
			}
		}
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && hover != "" {
			g.Page = hover
			if g.Page == "request-details" {
				i := (cursor[1] - 200) / 40
				if i < len(g.Building.Requests) {
					id := g.Building.Requests[i].ID
					g.Building.ActiveRequest = g.Building.RequestMap[id]
				} else {
					g.Page = "request-list"
				}
			}
		}

		if g.Page == "request-details" {
			// TODO: create option to mark closed without doing anything
			// TODO: logic to select a resolution option to try
			// tmp: hard-coded doing first option
			if inpututil.IsKeyJustPressed(ebiten.KeyS) {
				cost, time := g.Building.ActiveRequest.Resolve(0)
				g.AdvanceDay(time)
				g.Building.CreditBalance += cost
			}
		}

	} else if g.State == "infoControls" {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			g.State = ""
		}
	} else {
		switch {
		case titleClickable["play"].Hover(cursor):
			hover = "play"
		case titleClickable["controls"].Hover(cursor):
			hover = "controls"
		default:
			hover = ""
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyD) {
			fmt.Println(g.RequestPool[0])
		}
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			switch hover {
			case "play":
				g.State = "story"
			case "controls":
				g.State = "infoControls"
			}
		}
		// TODO
		// set a count-down to display transition
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		// TODO: dialogue to confirm player wants to exit game
		return ebiten.Termination
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.State == "load" {
	} else if g.State == "story" {
		g.IntroStory(screen)
	} else if g.State == "play" {
		g.DrawPortal(screen)
		g.DrawPortalPage(screen)
	} else {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(0.4, 0.4)
		screen.DrawImage(background["title"], op)
		if g.State == "infoControls" {
			ebitenutil.DrawRect(
				screen,
				400,
				350,
				480,
				260,
				color.Black,
			)
		}
	}
	/*
		ebitenutil.DebugPrintAt(screen, "Clock.Tick: "+strconv.Itoa(g.Clock.Tick), 20, 100)
		ebitenutil.DebugPrintAt(screen, "Clock.Month: "+strconv.Itoa(g.Clock.Month), 20, 120)
		ebitenutil.DebugPrintAt(screen, "Clock.Days: "+strconv.Itoa(g.Clock.Days), 20, 140)
		ebitenutil.DebugPrintAt(screen, "Cursor X: "+strconv.Itoa(cursor[0]), 30, 45)
	*/
	ebitenutil.DebugPrintAt(screen, "Cursor Y: "+strconv.Itoa(cursor[1]), 30, 65)
}
