package main

import (
	"embed"
	"image/color"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/tinne26/etxt"
)

var (
	load         = true
	hover        = ""
	play         = false
	infoControls = false
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
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	Width       int
	Height      int
	State       *State
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

	if load {
		initializeClickables()
		initializeTenants(tenants)
		g.initializeClock()
		g.initializeBuilding()
		g.initializeRequestPool(FileSystem)
		g.ConfigureTextRenderer()
		if len(background) == 2 {
			load = false
		}
	} else if play {
		// TODO: Make incrementing of months a function of tasks done(weight) + ticks
		g.Clock.Tick += 1
		if g.Clock.Tick%200 == 0 {
			g.Clock.IncrementMonth()
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.Building.GenerateRequest(g.RequestPool)
		}
		// TODO
		// logic for interacting with Maintenance Portal
	} else if infoControls {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			infoControls = false
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
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			switch hover {
			case "play":
				play = true
			case "controls":
				infoControls = true
			}
		}
		// TODO
		// set a count-down to display transition
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if load {
	} else if play {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(0.4, 0.4)
		screen.DrawImage(background["play"], op)

		vector.DrawFilledRect(screen, 0, 0, 1280.0, 120.0, color.RGBA{70, 10, 100, 125}, false)
		vector.DrawFilledRect(screen, 0, 120, 300.0, 860.0, color.RGBA{70, 10, 100, 125}, false)
		x := 800
		y := 180
		for _, r := range g.Building.Requests {
			y += 40
			g.Text.SetTarget(screen)
			g.Text.Draw(r.Title, x, y)
		}
	} else {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(0.4, 0.4)
		screen.DrawImage(background["title"], op)
		ebitenutil.DebugPrint(screen, "Menu - Press Enter to Play")
		if infoControls == true {
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
	ebitenutil.DebugPrintAt(screen, "Clock.Tick: "+strconv.Itoa(g.Clock.Tick), 20, 100)
	ebitenutil.DebugPrintAt(screen, "Clock.Month: "+strconv.Itoa(g.Clock.Month), 20, 120)
	ebitenutil.DebugPrintAt(screen, "Cursor X: "+strconv.Itoa(cursor[0]), 30, 45)
	ebitenutil.DebugPrintAt(screen, "Cursor Y: "+strconv.Itoa(cursor[1]), 30, 65)
}
