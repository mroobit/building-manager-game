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
	vector.DrawFilledRect(screen, 0, 80, 360.0, 820.0, color.RGBA{70, 30, 100, 95}, false)
	vector.DrawFilledRect(screen, 0, 900, 1280.0, 80.0, color.RGBA{170, 130, 200, 255}, false)

	g.SetTextProfile(textProfile["portal-button"])
	g.Text.SetTarget(screen)
	g.Text.Draw("Overview", 30, 100)
	g.Text.Draw("Requests", 30, 160)

	// TODO: improve alert circle logic
	// TODO: improve alert circle appearance
	var alertColor color.Color

	numR := len(g.Building.Requests)
	switch {
	case numR <= 3:
		alertColor = color.RGBA{20, 200, 20, 205}
	case numR <= 6:
		alertColor = color.RGBA{255, 190, 75, 205}
	default:
		alertColor = color.RGBA{250, 105, 90, 205}
	}

	vector.DrawFilledCircle(screen, 290.0, 180.0, 22.0, alertColor, false)
	g.SetTextProfile(textProfile["alert-button"])
	g.Text.Draw(strconv.Itoa(g.Building.OpenRequestCount()), 290, 180)

	g.SetTextProfile(textProfile["portal-button"])
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

func (g *Game) DrawPortalPage(screen *ebiten.Image) {
	g.Text.SetTarget(screen)

	titleX := 815
	titleY := 130

	crumbX := 370
	crumbY := 40

	labelX := crumbX + 30
	valueX := labelX + 650
	signX := valueX - 10
	sectionX := labelX

	switch g.Page {
	case "request-list":
		g.SetTextProfile(textProfile["portal-page-title"])
		g.Text.Draw("Open Maintenance Requests", titleX, titleY)

		g.SetTextProfile(textProfile["portal-header-footer"])
		g.Text.Draw("Home > Maintenance Requests", crumbX, crumbY)

		g.DrawRequestList(screen)

	case "request-details":
		g.SetTextProfile(textProfile["portal-page-title"])
		g.Text.Draw("Maintenance Request - Details", titleX, titleY)

		g.SetTextProfile(textProfile["portal-header-footer"])
		g.Text.Draw("Home > Maintenance Request Details", crumbX, crumbY)

		g.DrawRequestDetails(screen)

	case "financial-overview":
		g.SetTextProfile(textProfile["portal-page-title"])
		g.Text.Draw("Financial Overview", titleX, titleY)

		g.SetTextProfile(textProfile["portal-header-footer"])
		g.Text.Draw("Home > Financial Overview", crumbX, crumbY)

		y := titleY + 50

		g.SetTextProfile(textProfile["financial-section"])
		g.Text.Draw("Current Balances", sectionX, y)
		y += 40

		g.SetTextProfile(textProfile["financial-green"])
		g.Text.Draw("Bank Account Balance", labelX, y)
		g.Text.Draw("$"+strconv.Itoa(g.Building.Money), valueX, y)
		y += 40
		g.SetTextProfile(textProfile["financial-red"])
		g.Text.Draw("Credit Card Balance", labelX, y)
		g.Text.Draw("$"+strconv.Itoa(g.Building.CreditBalance), valueX, y)
		y += 60
		if g.Building.Money >= 0 {
			g.SetTextProfile(textProfile["financial-green"])
		} else {
			g.SetTextProfile(textProfile["financial-red"])
		}
		g.Text.Draw("Net Balance", labelX, y)
		g.Text.Draw("$"+strconv.Itoa(g.Building.Money), valueX, y)

		y += 60
		g.SetTextProfile(textProfile["financial-section"])
		g.Text.Draw("Upcoming Revenue and Costs", sectionX, y)
		y += 40
		g.SetTextProfile(textProfile["financial-red"])
		g.Text.Draw("Fixed Costs (eg mortgage, insurance)", labelX, y)
		g.Text.Draw("-", signX, y)
		g.Text.Draw("$"+strconv.Itoa(g.Building.FixedCosts), valueX, y)

		y += 40
		g.Text.Draw("Credit Card Payment", labelX, y)
		g.Text.Draw("-", signX, y)
		g.Text.Draw("$"+strconv.Itoa(g.Building.CreditBalance), valueX, y)

		y += 40
		g.SetTextProfile(textProfile["financial-green"])
		g.Text.Draw("Rent Income", labelX, y)
		g.Text.Draw("+", signX, y)
		g.Text.Draw("$"+strconv.Itoa(g.UpcomingRent()), valueX, y)

		y += 60
		if g.Building.Money >= 0 {
			g.SetTextProfile(textProfile["financial-green"])
		} else {
			g.SetTextProfile(textProfile["financial-red"])
		}
		g.Text.Draw("Net Change", labelX, y)
		g.Text.Draw("$"+strconv.Itoa(g.UpcomingPayments()), valueX, y)

	case "overview":
		fallthrough
	default:
		g.SetTextProfile(textProfile["portal-page-title"])
		g.Text.Draw("Overview", titleX, titleY)

		g.SetTextProfile(textProfile["portal-header-footer"])
		g.Text.Draw("Home", crumbX, crumbY)

		y := titleY + 50

		g.SetTextProfile(textProfile["portal-header-footer"])
		g.Text.Draw("Reputation", labelX, y)
		g.Text.Draw(strconv.Itoa(g.Building.Reputation), valueX, y)
		y += 40
		/*
			g.Text.Draw("Net Balance", labelX, y)
			g.Text.Draw(strconv.Itoa(g.NetBalance()), valueX, y)
			y += 40
		*/
		g.Text.Draw("Number of Tenants", labelX, y)
		g.Text.Draw(strconv.Itoa(len(g.Building.Tenants)), valueX, y)
		y += 40
		g.Text.Draw("Vacancies", labelX, y)
		g.Text.Draw(strconv.Itoa(10-len(g.Building.Tenants)), valueX, y)
		y += 40
		g.Text.Draw("Open Requests", labelX, y)
		g.Text.Draw(strconv.Itoa(len(g.Building.Requests)), valueX, y)
	}
}

func (g *Game) DrawRequestList(screen *ebiten.Image) {
	g.SetTextProfile(textProfile["portal-header-footer"])

	textX := 410
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

	g.SetTextProfile(textProfile["request-list"])

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

func (g *Game) DrawRequestDetails(screen *ebiten.Image) {
	// hard-code specific request for now, to use selected request later
	request := g.Building.Requests[0]

	labelCol := 410
	valueCol := 460
	y := 180

	g.SetTextProfile(textProfile["portal-header-footer"])
	g.Text.Draw("Issue", labelCol, y)
	g.Text.Draw(request.Title, valueCol, y)

	g.Text.Draw("Description", labelCol, y+40)
	g.Text.Draw(request.Description, valueCol, y+40) // write function to wrap text based on width available, text size

}
