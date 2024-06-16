package main

import (
	"embed"
	"image/color"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	load         = true
	hover        = ""
	play         = false
	infoControls = false
	cursor       [2]int
	tprint       = true
	rentIncrease = 3 // percentage increase at renewal

	//go:embed images
	//go:embed requests.json
	FileSystem embed.FS
)

func main() {
	gameWidth, gameHeight := 1280, 960

	ebiten.SetWindowSize(gameWidth, gameHeight)
	ebiten.SetWindowTitle("Ebitengine Game Jam '24")

	loadAssets()

	game := &Game{
		Width:  gameWidth,
		Height: gameHeight,
		Player: &Player{
			Money:      1000,
			Reputation: 7,
		},
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	Width       int
	Height      int
	Player      *Player
	Complex     *Building
	RequestPool []*Request
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return g.Width, g.Height
}

func (g *Game) Update() error {
	cursor[0], cursor[1] = ebiten.CursorPosition()

	if load {
		initializeClickables()
		initializeTenants(tenants)
		g.initializeBuilding()
		g.initializeRequestPool(FileSystem)
		load = false
	} else if play {
		if tprint {
			g.Complex.ListTenants()
			tprint = false
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.Complex.GenerateRequest(g.RequestPool)
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
	if play {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(0.4, 0.4)
		screen.DrawImage(portalBackground, op)
		ebitenutil.DebugPrint(screen, "Game")
		ebitenutil.DebugPrintAt(screen, strconv.Itoa(len(g.Complex.Tenants)), 0, 30)
		y := 0
		for _, r := range g.Complex.Requests {
			ebitenutil.DebugPrintAt(screen, r.Title, 40, y)
			y += 20
		}
	} else {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(0.4, 0.4)
		screen.DrawImage(titleBackground, op)
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
	ebitenutil.DebugPrintAt(screen, "Cursor X: "+strconv.Itoa(cursor[0]), 30, 45)
	ebitenutil.DebugPrintAt(screen, "Cursor Y: "+strconv.Itoa(cursor[1]), 30, 65)
}

type Player struct {
	// Name       string	// maybe include later the option to enter name
	Money      int
	Reputation int
}
