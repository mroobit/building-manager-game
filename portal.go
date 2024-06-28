package main

import (
	"image/color"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	portalSidebarWidth  = 350.0
	portalSidebarHeight = 845.0
	portalHeaderHeight  = 55.0
	portalFooterHeight  = 80.0
	portalWindowWidth   = 930.0
	portalWindowHeight  = portalSidebarHeight
)

var (
	portalPurple          = color.RGBA{170, 130, 200, 255}
	portalPurpleSecondary = color.RGBA{70, 30, 100, 95}
	portalTertiary        = color.RGBA{200, 200, 200, 255}
	white                 = color.RGBA{255, 255, 255, 255}
	whiteScreen           = color.RGBA{75, 75, 75, 95}
	black                 = color.RGBA{30, 30, 50, 235}
	transparentPurple     = color.RGBA{40, 0, 60, 30}
	rowColor              = []color.Color{black, portalPurpleSecondary}

	alertGreen  = color.RGBA{20, 200, 20, 205}
	alertYellow = color.RGBA{255, 190, 75, 205}
	alertRed    = color.RGBA{250, 105, 90, 205}
	diffRed     = color.RGBA{250, 130, 130, 255}
)

func (g *Game) DrawPortal(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, 0, 0, float32(g.Width), float32(g.Height), white, false)

	vector.DrawFilledRect(screen, 0, 0, float32(g.Width), portalHeaderHeight, portalPurple, false)
	vector.DrawFilledRect(screen, 0, portalHeaderHeight, portalSidebarWidth, portalSidebarHeight, portalPurpleSecondary, false)
	vector.DrawFilledRect(screen, 0, portalHeaderHeight+portalSidebarHeight, float32(1280.0), portalFooterHeight, portalPurple, false)

	g.SetTextProfile(textProfile["portal-button"])
	g.Text.SetTarget(screen)
	g.Text.Draw("Overview", 30, button["overview"].UpperLeft[1])
	g.Text.Draw("Requests", 30, button["request-list"].UpperLeft[1])

	// TODO: improve alert circle logic
	// TODO: improve alert circle appearance
	var alertColor color.Color

	numR := g.Building.OpenRequestCount()
	switch {
	case numR <= 3:
		alertColor = alertGreen
	case numR <= 6:
		alertColor = alertYellow
	default:
		alertColor = alertRed
	}

	vector.DrawFilledCircle(screen, 290.0, 180.0, 22.0, alertColor, false)
	g.SetTextProfile(textProfile["alert-button"])
	g.Text.Draw(strconv.Itoa(g.Building.OpenRequestCount()), 290, 180)

	g.SetTextProfile(textProfile["portal-button"])
	g.Text.Draw("Tenants", 30, button["tenants"].UpperLeft[1])
	g.Text.Draw("Financial Picture", 30, button["financial-overview"].UpperLeft[1])

	g.SetTextProfile(textProfile["portal-header-footer"])
	g.Text.Draw("2406 Ebiten Ln", 1100, 30)

	moneyLeftX := 42
	moneyRightX := 310
	moneyY := 562
	vector.DrawFilledRect(screen, 30, 550, 290.0, 150.0, white, false)
	g.SetTextProfile(textProfile["portal-money-left"])
	g.Text.Draw("Bank", moneyLeftX, moneyY)
	g.SetTextProfile(textProfile["portal-money-right"])
	g.Text.Draw("$"+strconv.Itoa(g.Building.Money), moneyRightX, moneyY)
	moneyY += 40
	g.SetTextProfile(textProfile["portal-money-left"])
	g.Text.Draw("Credit Card", moneyLeftX, moneyY)
	g.SetTextProfile(textProfile["portal-money-right-red"])
	g.Text.Draw("$"+strconv.Itoa(g.Building.CreditBalance), moneyRightX, moneyY)
	moneyY += 60
	g.SetTextProfile(textProfile["portal-money-left"])
	g.Text.Draw("Upcoming", moneyLeftX, moneyY)
	g.SetTextProfile(textProfile["portal-money-right-green"])
	sign := "+"
	netChange := g.UpcomingPayments()
	if netChange < 0 {
		sign = "-"
		netChange *= -1
		g.SetTextProfile(textProfile["portal-money-right-red"])
	}
	g.Text.Draw(sign+"$"+strconv.Itoa(netChange), moneyRightX, moneyY)

	vector.DrawFilledRect(screen, 30, 720, 290.0, 150.0, white, false)

	g.SetTextProfile(textProfile["portal-calendar-label"])
	g.Text.Draw("Days Left in Month", 175, 760)
	g.SetTextProfile(textProfile["portal-calendar"])
	g.Text.Draw(strconv.Itoa(30-(g.Clock.Days/3)), 175, 820)
}

func (g *Game) DrawPortalPage(screen *ebiten.Image) {
	g.Text.SetTarget(screen)

	titleX := 815
	titleY := 125

	crumbX := 370
	crumbY := 80

	labelX := crumbX + 30
	valueX := labelX + 650
	sectionX := labelX

	switch g.Page {
	case "request-list":
		g.SetTextProfile(textProfile["portal-page-title"])
		g.Text.Draw("Open Tenant Requests", titleX, titleY)

		g.SetTextProfile(textProfile["portal-breadcrumb"])
		g.Text.Draw("Home > Tenant Requests", crumbX, crumbY)

		g.DrawRequestList(screen)

	case "request-details":
		g.SetTextProfile(textProfile["portal-page-title"])
		g.Text.Draw("Tenant Request - Details", titleX, titleY)

		g.SetTextProfile(textProfile["portal-breadcrumb"])
		g.Text.Draw("Home > Tenant Requests > Request Details", crumbX, crumbY)

		g.DrawRequestDetails(screen)
		g.DrawResolveClose(screen)

	case "try-to-resolve":
		g.SetTextProfile(textProfile["portal-page-title"])
		g.Text.Draw("Tenant Request - Details", titleX, titleY)

		g.SetTextProfile(textProfile["portal-breadcrumb"])
		g.Text.Draw("Home > Tenant Requests > Request Details", crumbX, crumbY)

		g.DrawRequestDetails(screen)
		g.DrawSolutions(screen)
	case "resolution-outcome":
		g.DrawOutcome(screen)
	case "financial-overview":
		g.SetTextProfile(textProfile["portal-page-title"])
		g.Text.Draw("Financial Overview", titleX, titleY)

		g.SetTextProfile(textProfile["portal-breadcrumb"])
		g.Text.Draw("Home > Financial Overview", crumbX, crumbY)

		y := titleY + 50
		valueX += 150
		sign := "-"

		g.SetTextProfile(textProfile["financial-section"])
		g.Text.Draw("Current Balances", sectionX, y)
		y += 40

		g.SetTextProfile(textProfile["financial-green"])
		g.Text.Draw("Bank Account Balance", labelX, y)
		g.SetTextProfile(textProfile["financial-right-green"])
		g.Text.Draw("$"+strconv.Itoa(g.Building.Money), valueX, y)
		y += 40
		g.SetTextProfile(textProfile["financial-red"])
		g.Text.Draw("Credit Card Balance", labelX, y)
		g.SetTextProfile(textProfile["financial-right-red"])
		g.Text.Draw("$"+strconv.Itoa(g.Building.CreditBalance), valueX, y)
		y += 60
		netBalance := g.Building.Money - g.Building.CreditBalance
		if netBalance >= 0 {
			g.SetTextProfile(textProfile["financial-green"])
		} else {
			g.SetTextProfile(textProfile["financial-red"])
		}
		g.Text.Draw("Net Balance", labelX, y)
		if netBalance >= 0 {
			g.SetTextProfile(textProfile["financial-right-green"])
			sign = ""
		} else {
			g.SetTextProfile(textProfile["financial-right-red"])
			sign = "-"
			netBalance *= -1
		}
		g.Text.Draw(sign+"$"+strconv.Itoa(netBalance), valueX, y)

		y += 60
		g.SetTextProfile(textProfile["financial-section"])
		g.Text.Draw("Upcoming Revenue and Costs", sectionX, y)
		y += 40
		g.SetTextProfile(textProfile["financial-red"])
		g.Text.Draw("Fixed Costs (eg mortgage, insurance)", labelX, y)
		g.SetTextProfile(textProfile["financial-right-red"])
		g.Text.Draw("-$"+strconv.Itoa(g.Building.FixedCosts), valueX, y)

		y += 40
		g.SetTextProfile(textProfile["financial-red"])
		g.Text.Draw("Credit Card Payment", labelX, y)
		g.SetTextProfile(textProfile["financial-right-red"])
		g.Text.Draw("-$"+strconv.Itoa(g.Building.CreditBalance), valueX, y)

		y += 40
		g.SetTextProfile(textProfile["financial-green"])
		g.Text.Draw("Rent Income", labelX, y)
		g.SetTextProfile(textProfile["financial-right-green"])
		g.Text.Draw("+$"+strconv.Itoa(g.UpcomingRent()), valueX, y)

		y += 60

		netChange := g.UpcomingPayments()
		if netChange >= 0 {
			g.SetTextProfile(textProfile["financial-green"])
			sign = ""
		} else {
			g.SetTextProfile(textProfile["financial-red"])
			sign = "-"
			netChange *= -1
		}
		g.Text.Draw("Net Change", labelX, y)
		g.SetTextProfile(textProfile["financial-right-green"])
		g.Text.Draw(sign+"$"+strconv.Itoa(netChange), valueX, y)

	case "tenants":
		g.SetTextProfile(textProfile["portal-page-title"])
		g.Text.Draw("Tenants", titleX, titleY)

		g.SetTextProfile(textProfile["portal-breadcrumb"])
		g.Text.Draw("Home > Tenants", crumbX, crumbY)

		unitX := labelX + 30
		unitRightX := labelX + 330
		unitY := titleY + 60
		rectX := float32(labelX + 10)
		rectY := float32(titleY + 50)
		yIncr := 30
		for i, t := range g.Building.Tenants {
			vector.DrawFilledRect(screen, rectX, rectY, 340.0, 105.0, portalTertiary, false)

			g.SetTextProfile(textProfile["tenant-bold-left"])
			g.Text.Draw("Unit "+t.Unit, unitX, unitY)
			if t.Name == "" {
				g.SetTextProfile(textProfile["tenant-bold-right"])
				g.Text.Draw("(Vacant)", unitRightX, unitY)
				unitY += yIncr * 3
			} else {
				g.SetTextProfile(textProfile["tenant-regular-right"])
				g.Text.Draw(t.Name, unitRightX, unitY)
				unitY += yIncr
				g.SetTextProfile(textProfile["tenant-regular-left"])
				g.Text.Draw("Rent", unitX, unitY)
				g.SetTextProfile(textProfile["tenant-regular-right"])
				g.Text.Draw(strconv.Itoa(t.Rent), unitRightX, unitY)
				unitY += yIncr
				g.SetTextProfile(textProfile["tenant-regular-left"])
				g.Text.Draw("Months Left in Lease", unitX, unitY)
				g.SetTextProfile(textProfile["tenant-regular-right"])
				g.Text.Draw(strconv.Itoa(t.MonthsRemaining), unitRightX, unitY)
				unitY += yIncr
			}
			unitY += yIncr + 10
			rectY += float32(yIncr + 100)
			if i == 4 {
				unitX += 440
				unitRightX += 440
				rectX += 440
				unitY = titleY + 60
				rectY = float32(titleY + 50)
			}
		}
	case "overview":
		fallthrough
	default:
		g.SetTextProfile(textProfile["portal-page-title"])
		g.Text.Draw("Overview", titleX, titleY)

		g.SetTextProfile(textProfile["portal-breadcrumb"])
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
		g.Text.Draw(strconv.Itoa(g.Building.Vacancies()), valueX, y)
		y += 40
		g.Text.Draw("Open Requests", labelX, y)
		g.Text.Draw(strconv.Itoa(len(g.Building.Requests)), valueX, y)
		y += 40
		g.Text.Draw("Months Until Next Inspection", labelX, y)
		g.Text.Draw(strconv.Itoa(g.Building.Inspection), valueX, y)
		/*
			y += 40
			g.Text.Draw("Next Spraying: "+strconv.Itoa(g.Building., labelX, y)
			g.Text.Draw(strconv.Itoa(len(g.Building.Requests)), valueX, y)
		*/
	}
}

func (g *Game) DrawRequestList(screen *ebiten.Image) {
	g.SetTextProfile(textProfile["portal-header-footer"])

	textX := 410
	y := 180

	issueCol := textX
	receivedCol := textX + 360
	locationCol := textX + 500
	nameCol := textX + 630

	vector.DrawFilledRect(screen, 390, 160, 850.0, 40.0, portalPurple, false)
	g.Text.Draw("Issue", issueCol, y)
	g.Text.Draw("Received", receivedCol, y)
	g.Text.Draw("Location", locationCol, y)
	g.Text.Draw("Reported By", nameCol, y)

	vector.DrawFilledRect(
		screen,
		float32(button["request-details"].UpperLeft[0]),
		float32(button["request-details"].UpperLeft[1]),
		float32(button["request-details"].Width),
		float32(button["request-details"].Height),
		portalTertiary,
		false,
	)
	g.SetTextProfile(textProfile["request-list"])

	//	pagination := 0 // for when there are > 16 requests, allow navigation to additional requests?

	//	sortedRequests := g.Building.RequestsByUrgency()

	for _, r := range g.Building.Requests {
		if !r.Closed {
			y += 40
			received := ""
			if r.DaysOpen == 0 {
				received = "Today"
			} else if r.DaysOpen == 1 {
				received = "Yesterday"
			} else if r.DaysOpen < 7 {
				received = strconv.Itoa(r.DaysOpen) + " days ago"
			} else if r.DaysOpen < 14 {
				received = "Last week"
			} else if r.DaysOpen < 31 {
				n := r.DaysOpen / 7
				received = strconv.Itoa(n) + " weeks ago"
			} else if r.DaysOpen < 60 {
				received = "Last month"
			} else {
				n := r.DaysOpen / 30
				received = strconv.Itoa(n) + " months ago"
			}
			g.Text.Draw(r.Title, issueCol, y)
			g.Text.Draw(received, receivedCol, y) // TODO: increment r.DaysOpen when g.Clock.Days does
			g.Text.Draw(r.Location, locationCol, y)
			g.Text.Draw(r.Tenant.Name, nameCol, y)
		}
	}

}

func (g *Game) DrawRequestDetails(screen *ebiten.Image) {
	labelCol := 410
	valueCol := labelCol + 140
	y := 180

	g.SetTextProfile(textProfile["request-description"])
	g.Text.Draw("Issue", labelCol, y)
	g.Text.Draw(g.Building.ActiveRequest.Title, valueCol, y)

	y += 40
	g.Text.Draw("Location", labelCol, y)
	g.Text.Draw(g.Building.ActiveRequest.Location, valueCol, y)

	y += 40
	g.Text.Draw("Reporter", labelCol, y)
	g.Text.Draw(g.Building.ActiveRequest.Tenant.Name, valueCol, y)

	y += 40
	g.Text.Draw("Description", labelCol, y)
	g.Text.Draw(wrapText(g.Building.ActiveRequest.Description, 60), valueCol, y)

}

func (g *Game) DrawResolveClose(screen *ebiten.Image) {
	vector.DrawFilledRect(
		screen,
		float32(button["try-to-resolve"].UpperLeft[0]),
		float32(button["try-to-resolve"].UpperLeft[1]),
		float32(button["try-to-resolve"].Width),
		float32(button["try-to-resolve"].Height),
		portalPurple,
		false,
	)
	vector.DrawFilledRect(
		screen,
		float32(button["close-request"].UpperLeft[0]),
		float32(button["close-request"].UpperLeft[1]),
		float32(button["close-request"].Width),
		float32(button["close-request"].Height),
		diffRed,
		false,
	)

	g.SetTextProfile(textProfile["request-resolve-close"])
	g.Text.Draw("Try to Resolve", 665, 435)
	g.Text.Draw("Close Request", 965, 435)
}

func (g *Game) DrawSolutions(screen *ebiten.Image) {
	// TODO: Draw a label over a box of solutions
	// h := float32(g.Building.ActiveRequest.AvailableSolutionsCount()) * 50.0
	x := button["solutions"].UpperLeft[0] + 10
	y := button["solutions"].UpperLeft[1] + 30
	solutionSpacing := 60

	g.SetTextProfile(textProfile["request-solutions"])

	spaceY := 0
	rowI := 0
	for _, s := range g.Building.ActiveRequest.Solutions {
		if !s.Attempted {
			vector.DrawFilledRect(
				screen,
				float32(button["solutions"].UpperLeft[0]),
				float32(button["solutions"].UpperLeft[1]+spaceY),
				float32(button["solutions"].Width),
				60.0,
				rowColor[rowI],
				false,
			)
			g.Text.Draw(s.Action+" ($"+strconv.Itoa(s.Cost)+")", x, y)
			y += solutionSpacing
			spaceY += solutionSpacing
			rowI = (rowI + 1) % 2
		}
	}
}

func (g *Game) DrawOutcome(screen *ebiten.Image) {
	// TODO display story text for the outcome of the solution selected
	vector.DrawFilledRect(
		screen,
		0,
		0,
		float32(g.Width),
		float32(g.Height),
		transparentPurple,
		false,
	)
	vector.DrawFilledRect(
		screen,
		float32(g.Width/7),
		float32(g.Height/5),
		float32(5*g.Width/7),
		float32(3*g.Height/5),
		white,
		false,
	)
	r := g.Building.ActiveRequest
	outcome := r.Solutions[r.Attempts[len(r.Attempts)-1]].Outcome

	g.SetTextProfile(textProfile["portal-calendar"])
	g.Text.Draw(wrapText(outcome, 30), g.Width/2, g.Height/2)
}

func wrapText(s string, length int) string {
	line := ""
	w := ""
	sWords := strings.Split(s, " ")

	for i, word := range sWords {
		if strings.Contains(word, "\n") {
			n := strings.Split(word, "\n")
			w = w + line + " " + n[0] + "\n"
			line = n[1]
		} else if len(line)+len(word) <= length {
			if line == "" {
				line = word
			} else {
				line = line + " " + word
			}
		} else {
			w = w + line + "\n"
			line = word
		}

		if i == len(sWords)-1 {
			w = w + line
		}
	}

	return w
}
