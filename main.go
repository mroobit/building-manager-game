package main

import (
	"embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/tinne26/etxt"
)

var (
	hover        = ""
	cursor       [2]int
	rentIncrease = 3 // percentage increase at renewal
	debugIndex   int

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
	TenantPool  []*Tenant
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
		g.initializeClock()
		g.initializeBuilding()
		g.initializeTenantPool(FileSystem)
		g.initializeTenants()
		g.initializeRequestPool(FileSystem)
		g.ConfigureTextRenderer()
		loadLetter(FileSystem)
		g.CreateProblems()
		g.State = "title"
	} else if g.State == "story" {
		auntJosLetter.Play()
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.State = "meta"
			g.Page = "login"
			auntJosLetter.Pause()
		}
		if button["continue"].Hover(cursor) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			g.State = "meta"
			g.Page = "login"
			auntJosLetter.Pause()
		}
	} else if g.State == "meta" {
		loop0.Play()

		if g.Page == "login" {
			switch {
			case button["login-play"].Hover(cursor):
				hover = "login-play"
			case button["how-to-play"].Hover(cursor):
				hover = "how-to-play"
			case button["settings"].Hover(cursor):
				hover = "settings"
			case button["about"].Hover(cursor):
				hover = "about"
			default:
				hover = ""
			}
		}
		if g.Page == "settings" || g.Page == "about" || g.Page == "how-to-play" || g.Page == "ending" {
			switch {
			case button["upper-x"].Hover(cursor):
				hover = "upper-x"
			case button["back"].Hover(cursor):
				hover = "back"
			}
		}
		if g.Page == "settings" {
			if button["volume"].Hover(cursor) {
				hover = "volume"
			}
		}
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && hover != "" {
			switch hover {
			case "login-play":
				g.Page = "overview"
				g.State = "play"
				loop0.Pause()
			case "how-to-play":
				g.Page = "how-to-play"
			case "settings":
				g.Page = "settings"
			case "volume":
				switch {
				case cursor[0] < button["volume"].UpperLeft[0]+button["volume"].Width/8:
					musicVolume = 0
				case cursor[0] < button["volume"].UpperLeft[0]+3*button["volume"].Width/8:
					musicVolume = 0.25
				case cursor[0] < button["volume"].UpperLeft[0]+5*button["volume"].Width/8:
					musicVolume = 0.5
				case cursor[0] < button["volume"].UpperLeft[0]+7*button["volume"].Width/8:
					musicVolume = 0.75
				default:
					musicVolume = 1
				}
				loop0.SetVolume(musicVolume)
				loop1.SetVolume(musicVolume)
				loop2.SetVolume(musicVolume)
			case "about":
				g.Page = "about"
			case "upper-x":
				g.Page = "login"
			case "back":
				g.Page = "login"
			}
		}
	} else if g.State == "play" {
		loop1.Play()
		loop2.Play()
		g.Clock.Tick += 1
		g.AdvanceDayByTicks()
		g.CreateProblems()
		g.CheckEndOfMonth()

		if (g.Building.Money == 0 && g.Building.CreditBalance > 3000 && g.UpcomingPayments() < 0) || g.Building.Vacancies() == len(g.Building.Tenants) || (g.Building.Money > 100000) {
			loop1.Pause()
			loop2.Pause()
			g.Page = "ending"
			g.State = "meta"
		}

		switch {
		case button["overview"].Hover(cursor):
			hover = "overview"
		case button["request-list"].Hover(cursor):
			hover = "request-list"
		case button["financial-overview"].Hover(cursor):
			hover = "financial-overview"
		case button["tenants"].Hover(cursor):
			hover = "tenants"
		default:
			hover = ""
		}

		if g.Page == "request-list" {
			if button["request-details"].Hover(cursor) {
				hover = "request-details"
			}
		}
		if g.Page == "request-details" {
			if button["try-to-resolve"].Hover(cursor) {
				hover = "try-to-resolve"
			} else if button["close-request"].Hover(cursor) {
				hover = "close-request"
			}
		}
		if g.Page == "try-to-resolve" && button["solutions"].Hover(cursor) {
			hover = "solutions"
		}
		if g.Page == "resolution-outcome" {
			if button["continue"].Hover(cursor) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
				g.Page = "request-list"
			}
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
				i := (cursor[1] - button["solutions"].UpperLeft[1] + 20) / 70
				trueIndices := g.Building.ActiveRequest.AvailableSolutionIndices()
				if cursor[0] >= button["solutions"].UpperLeft[0] && cursor[0] <= button["solutions"].LowerRight[0] && i < len(trueIndices) {
					trueIndex := trueIndices[i]
					cost, time := g.Building.ActiveRequest.Resolve(trueIndex)
					if g.Building.ActiveRequest.ResolutionQuality >= 7 {
						g.Building.RequestsAddressed += 1
					}
					g.Building.CreditBalance += cost
					g.Page = "resolution-outcome"
					g.AdvanceDay(time)
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
		if button["back"].Hover(cursor) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			g.Page = "request-list"
			g.State = "play"
		}

	} else if g.State == "title" {
		if button["start"].Hover(cursor) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			g.State = "story"
		}
		if button["skip"].Hover(cursor) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			g.State = "meta"
			g.Page = "login"
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		// TODO: dialogue to confirm player wants to exit game
		return ebiten.Termination
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.State == "load" {
		op := &ebiten.DrawImageOptions{}
		screen.DrawImage(theBuilding, op)
	} else if g.State == "title" {
		op := &ebiten.DrawImageOptions{}
		screen.DrawImage(theBuilding, op)

		g.Text.SetTarget(screen)

		g.SetTextProfile(textProfile["title-button"])
		vector.DrawFilledRect(
			screen,
			float32(button["start"].UpperLeft[0]),
			float32(button["start"].UpperLeft[1]),
			float32(button["start"].Width),
			float32(button["start"].Height),
			whiteTranslucent,
			false,
		)
		g.Text.Draw(
			"Start",
			button["start"].Width/2+button["start"].UpperLeft[0],
			button["start"].Height/2+button["start"].UpperLeft[1],
		)

		g.SetTextProfile(textProfile["title-button"])
		vector.DrawFilledRect(
			screen,
			float32(button["skip"].UpperLeft[0]),
			float32(button["skip"].UpperLeft[1]),
			float32(button["skip"].Width),
			float32(button["skip"].Height),
			whiteTranslucent,
			false,
		)
		g.Text.Draw(
			"Skip Story",
			button["skip"].Width/2+button["skip"].UpperLeft[0],
			button["skip"].Height/2+button["skip"].UpperLeft[1],
		)
	} else if g.State == "story" {
		g.IntroStory(screen)
	} else if g.State == "meta" {
		g.DrawMeta(screen)
	} else if g.State == "play" {
		g.DrawPortal(screen)
		g.DrawPortalPage(screen)
	} else if g.State == "monthReport" {
		g.DrawMonthEndReport(screen)
	}
}
