package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (g *Game) DrawMeta(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, 0, 0, float32(g.Width), float32(g.Height), white, false)

	switch g.Page {
	case "login":
		// TODO animate username getting filled in with "manager"
		// include audio for typing on both of these animations!!
		// animate dots filling in password
		// animate "loading" the portal
		// Background
		vector.DrawFilledRect(screen, 0, 0, float32(g.Width), float32(g.Height), white, false)
		vector.DrawFilledRect(screen, 0, 0, float32(g.Width), float32(g.Height), portalPurpleSecondary, false)

		// Dropshadow, then Central White Rectangle
		vector.DrawFilledRect(screen, 398, 161, 484.0, 661.0, portalPurpleSecondary, false)
		vector.DrawFilledRect(screen, 400, 160, 480.0, 660.0, white, false)
		g.SetTextProfile(textProfile["portal-page-title"])
		g.Text.Draw("Building\nManagement\nPortal", 640, 430)

		// Login Boxes, Button, and Link
		vector.DrawFilledRect(screen, 460, 530, 360.0, 40.0, portalTertiary, false)
		g.SetTextProfile(textProfile["login-label"])
		g.Text.Draw("Email", 460, 524)
		g.SetTextProfile(textProfile["login"])
		g.Text.Draw("manager@ebitenbldg.com", 470, 540)

		vector.DrawFilledRect(screen, 460, 610, 360.0, 40.0, portalTertiary, false)
		g.SetTextProfile(textProfile["login-label"])
		g.Text.Draw("Password", 460, 604)
		g.SetTextProfile(textProfile["login-password"])
		g.Text.Draw("●●●●●●●●●●●●●●●●●●", 470, 620)

		vector.DrawFilledRect(
			screen,
			float32(button["login-play"].UpperLeft[0]),
			float32(button["login-play"].UpperLeft[1]),
			float32(button["login-play"].Width),
			float32(button["login-play"].Height),
			portalPurple,
			false,
		)
		g.SetTextProfile(textProfile["login-play"])
		g.Text.Draw("Play Game", 640, 690)

		g.SetTextProfile(textProfile["login-text-link"])
		g.Text.Draw(
			"Need to learn how to play?",
			button["how-to-play"].Width/2+button["how-to-play"].UpperLeft[0],
			button["how-to-play"].Height/2+button["how-to-play"].UpperLeft[1],
		)

		// Below-box buttons
		g.SetTextProfile(textProfile["login-lower-button"])
		vector.DrawFilledRect(
			screen,
			float32(button["settings"].UpperLeft[0]),
			float32(button["settings"].UpperLeft[1]),
			float32(button["settings"].Width),
			float32(button["settings"].Height),
			black,
			false,
		)
		g.Text.Draw(
			"Settings",
			button["settings"].Width/2+button["settings"].UpperLeft[0],
			button["settings"].Height/2+button["settings"].UpperLeft[1],
		)
		vector.DrawFilledRect(
			screen,
			float32(button["about"].UpperLeft[0]),
			float32(button["about"].UpperLeft[1]),
			float32(button["about"].Width),
			float32(button["about"].Height),
			black,
			false,
		)
		g.Text.Draw(
			"About",
			button["about"].Width/2+button["about"].UpperLeft[0],
			button["about"].Height/2+button["about"].UpperLeft[1],
		)
	case "how-to-play":
		// TODO
		g.DrawPortal(screen)
		vector.DrawFilledRect(screen, 0, 0, float32(g.Width), float32(g.Height), whiteScreen, false)
		vector.DrawFilledRect(
			screen,
			float32(button["upper-x"].UpperLeft[0]),
			float32(button["upper-x"].UpperLeft[1]),
			float32(button["upper-x"].Width),
			float32(button["upper-x"].Height),
			white,
			false,
		)
		g.SetTextProfile(textProfile["portal-page-title"])
		g.Text.Draw("x", g.Width-40, 40)
	case "about":
		// TODO
		g.SetTextProfile(textProfile["portal-page-title"])
		g.Text.Draw(
			wrapText("This game was created by Shannon Dybvig for the Ebitengine Game Jam 2024 (theme: Building).\n \nDesign, writing, art, audio, and programming are all by Shannon Dybvig.\n \nEbitengine is a 2D game engine created by Hajime Hoshi.", 30),
			g.Width/2,
			g.Height/2,
		)
		g.Text.Draw("x", g.Width-40, 40)

	case "settings":
		// TODO
		g.SetTextProfile(textProfile["portal-page-title"])
		g.Text.Draw("This is the Settings Page", 400, 400)
		g.Text.Draw("x", g.Width-40, 40)
	case "ending":
		// TODO
	}
}

func (g *Game) DrawAbout(screen *ebiten.Image) {
	// DrawText:
	// - credits for making game
	// - Ebitengine
	// - Fonts
	// - Acknowledge friends
}

func (g *Game) DrawSettings(screen *ebiten.Image) {
	// Settings for sound only
	// - everything
	// - music ????
	// - voice
	// - effects
}

func (g *Game) DrawHowToPlay(screen *ebiten.Image) {
	// DrawPortal? Then overlay stuff? Hover behaviors?
	// Maybe in the body of the portal explain, suggest user hover for explanations
}

func (g *Game) DrawLogin(screen *ebiten.Image) {
	// copy from play
	// hook up to non-play game mode
}

func (g *Game) DrawEnding(screen *ebiten.Image) {
	// draw ending of game based on game/building end-stats
}
