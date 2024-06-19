package main

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (g *Game) DrawPortal(screen *ebiten.Image) {
	/*
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(0.4, 0.4)
		screen.DrawImage(background["play"], op)
	*/

	vector.DrawFilledRect(screen, 0, 0, 1280.0, 980.0, color.RGBA{255, 255, 255, 255}, false)

	vector.DrawFilledRect(screen, 0, 0, 1280.0, 80.0, color.RGBA{170, 130, 200, 255}, false)
	vector.DrawFilledRect(screen, 0, 80, 360.0, 820.0, color.RGBA{70, 10, 100, 125}, false)
	vector.DrawFilledRect(screen, 0, 900, 1280.0, 80.0, color.RGBA{170, 130, 200, 255}, false)

	g.SetTextProfile(textProfile["portal-button"])
	g.Text.SetTarget(screen)
	g.Text.Draw("Overview", 30, 100)
	g.Text.Draw("Maintenace Requests", 30, 160)
	g.Text.Draw("Financials", 30, 220)

	g.SetTextProfile(textProfile["portal-header-footer"])
	g.Text.Draw("2406 Ebiten Ln", 1100, 40)
}

/*
func (g *Game) DrawMaintenance(screen *ebiten.Image) {
...or
func (g *Game) DrawLayout(screen *ebiten.Image, page string) { // page is string declaring which page is active, eg Requests
}

*/

func (g *Game) DrawRequestList(screen *ebiten.Image) {
	g.Text.SetTarget(screen)
	g.SetTextProfile(textProfile["portal-page-title"])

	g.Text.Draw("Open Maintenance Requests", 815, 130)

	g.SetTextProfile(textProfile["portal-header-footer"])

	textX := 400
	y := 180

	issueCol := textX
	receivedCol := textX + 280
	locationCol := textX + 380
	resolvedCol := textX + 480

	vector.DrawFilledRect(screen, 390, 160, 850.0, 40.0, color.RGBA{170, 130, 200, 255}, false)
	g.Text.Draw("Issue", issueCol, y)
	g.Text.Draw("Received", receivedCol, y)
	g.Text.Draw("Location", locationCol, y)
	g.Text.Draw("Resolved?", resolvedCol, y)

	for _, r := range g.Building.Requests {
		y += 40
		received := ""
		if r.DaysOpen == 0 {
			received = "Today"
		} else if r.DaysOpen == 1 {
			received = "Yesterday"
		} else {
			received = strconv.Itoa(r.DaysOpen) + " days ago"
		}
		g.Text.Draw(r.Title, issueCol, y)
		g.Text.Draw(received, receivedCol, y) // TODO: increment r.DaysOpen when g.Clock.Days does
		g.Text.Draw(r.Location, locationCol, y)
		g.Text.Draw(strconv.FormatBool(r.Resolved), resolvedCol, y)
	}

}
