package main

import (
	"embed"
	"image/color"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/tinne26/etxt"
)

const SampleRate = 44100

var (
	hover        = ""
	cursor       [2]int
	rentIncrease = 3 // percentage increase at renewal

	//go:embed audio
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
	AudioPlayer *audio.Player
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
		g.ConfigureAudio()
		loadLetter(FileSystem)
		if len(background) == 2 {
			g.State = "title"
		}
	} else if g.State == "story" {
		// g.AudioPlayer = auntJosLetter
		// g.AudioPlayer.Play()
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.State = "play"
			g.Page = "login"
			g.AudioPlayer.Pause()
		}
	} else if g.State == "play" {
		// TODO: Make incrementing of months a function of tasks done(weight) + ticks
		g.Clock.Tick += 1
		g.AdvanceDayByTicks()
		// TODO: generate problems based on Tick/Day + some randomness
		g.CreateProblems()

		if g.Page == "login" {
			if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
				g.Page = "overview"
			}
		}

		// TODO remove this, it is just for diagnostic purposes
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.GenerateRequest()
		}

		switch {
		case portalButton["overview"].Hover(cursor):
			hover = "overview"
		case portalButton["request-list"].Hover(cursor):
			hover = "request-list"
		case portalButton["financial-overview"].Hover(cursor):
			hover = "financial-overview"
		case portalButton["tenants"].Hover(cursor):
			hover = "tenants"
		default:
			hover = ""
		}
		if g.Page == "request-list" {
			if portalButton["request-details"].Hover(cursor) {
				hover = "request-details"
			}
		}
		if g.Page == "request-details" {
			if portalButton["try-to-resolve"].Hover(cursor) {
				hover = "try-to-resolve"
			} else if portalButton["close-request"].Hover(cursor) {
				hover = "close-request"
			}
		}
		if g.Page == "try-to-resolve" && portalButton["solutions"].Hover(cursor) {
			hover = "solutions"
		}
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && hover != "" {

			switch hover {
			case "overview":
				g.Page = "overview"
			case "request-list":
				g.Page = "request-list"
			case "financial-overview":
				g.Page = "financial-overview"
			case "tenants":
				g.Page = "tenants"
			case "try-to-resolve":
				g.Page = "try-to-resolve"
			case "close-request":
				g.Building.ActiveRequest.Closed = true
				g.Page = "request-list"
			case "solutions":
				i := (cursor[1] - 400) / 50
				if cursor[0] >= 400 && cursor[0] <= 1250 && i < len(g.Building.ActiveRequest.Solutions) {
					solution := g.Building.ActiveRequest.Solutions[i]
					cost, time := g.Building.ActiveRequest.Resolve(solution)
					if g.Building.ActiveRequest.ResolutionQuality >= 7 {
						g.Building.ActiveRequest.Resolved = true
						g.Building.RequestsAddressed += 1
					}
					g.Building.CreditBalance += cost
					g.AdvanceDay(time)
					// TODO: add a dialgoue about doing the solution and its effect
					// as well as that you've closed the issue
					g.Page = "request-list"
				}
			case "request-details":
				i := (cursor[1] - 200) / 40
				trueIndices := g.Building.OpenIndices()
				if i < len(trueIndices) {
					id := g.Building.Requests[trueIndices[i]].ID
					g.Building.ActiveRequest = g.Building.RequestMap[id]
					g.Page = "request-details"
				}
			}
		}
	} else if g.State == "monthReport" {
		// TODO play a little reaction sound, or some music?
		// add a clickable?
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.IncrementMonth()
			g.State = "play"
		}
	} else if g.State == "infoControls" {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			g.State = ""
		}
	} else {
		switch {
		case titleButton["play"].Hover(cursor):
			hover = "play"
		case titleButton["controls"].Hover(cursor):
			hover = "controls"
		default:
			hover = ""
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
	} else if g.State == "monthReport" {
		g.MonthEndReport(screen)
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
