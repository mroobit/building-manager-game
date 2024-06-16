package main

import (
	"embed"
	"fmt"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	play         = false
	tprint       = true
	rentIncrease = 3 // percentage increase at renewal

	//go:embed images
	FileSystem embed.FS
)

func main() {
	gameWidth, gameHeight := 1280, 960

	ebiten.SetWindowSize(gameWidth, gameHeight)
	ebiten.SetWindowTitle("Ebitengine Game Jam '24")

	loadAssets()
	initializeTenants(tenants)
	initializeBuilding()

	game := &Game{
		Width:  gameWidth,
		Height: gameHeight,
		Player: &Player{
			Money:      1000,
			Reputation: 7,
		},
		Complex: building,
		Tenants: tenants,
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

	fmt.Println("vim-go")
}

type Game struct {
	Width   int
	Height  int
	Player  *Player
	Complex *Building
	Tenants []*Tenant
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return g.Width, g.Height
}

func (g *Game) Update() error {
	if play {
		if tprint {
			g.Complex.ListTenants()
			tprint = false
		}
		// TODO
		// logic for interacting with Maintenance Portal
	} else {
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			play = true
		}
		// TODO
		// if keyboard input == Enter, play = true
		//	Later: logic to interact with menu
		// set a count-down to display transition
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if play {
		ebitenutil.DebugPrint(screen, "Game")
		ebitenutil.DebugPrintAt(screen, strconv.Itoa(len(g.Tenants)), 0, 30)
		// TODO
		// logic for displaying Maintenance Portal
	} else {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(0.4, 0.4)
		screen.DrawImage(titleBackground, op)
		ebitenutil.DebugPrint(screen, "Menu - Press Enter to Play")
		// TODO
		// logic to display Building Art
		// logic to display "Play" button
		// Later:	logic to display "Controls" menu button
	}
}

type Player struct {
	// Name       string	// maybe include later the option to enter name
	Money      int
	Reputation int
}
