package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Clock struct {
	Tick      int
	Timer     int
	Days      int
	Month     int
	Recurring map[string][2]int
}

func (g *Game) initializeClock() {
	c := &Clock{
		Recurring: map[string][2]int{
			"Spray for bugs":    [2]int{3, 2},
			"Annual inspection": [2]int{12, 3},
		},
	}
	g.Clock = c
}

func (g *Game) IncrementMonth() {
	g.Clock.Month += 1
	for key, value := range g.Clock.Recurring {
		update := value
		update[1] += 1
		g.Clock.Recurring[key] = update
	}
	g.ProcessPayments()
	g.Building.DecrementLeases()
	g.Clock.CheckEvents()
	g.Clock.Days = 0
	g.Clock.Tick = 0
	g.Building.RequestsAddressed = 0
}

func (g *Game) MonthEndReport(screen *ebiten.Image) {
	vector.DrawFilledRect(
		screen,
		0,
		0,
		float32(g.Width),
		float32(g.Height),
		white,
		false,
	)

	g.Text.SetTarget(screen)

	g.SetTextProfile(textProfile["portal-page-title"])
	g.Text.Draw("Month End Report", 640, 200)
	g.SetTextProfile(textProfile["request-description"])
	g.Text.Draw("This month you addressed "+strconv.Itoa(g.Building.RequestsAddressed)+" tenant requests!", 300, 230)

	// TODO: either IncrementMonth() will have to come before this and output numbers
	// or it will have to be broken apart so that non-payment of rent can be mentioned

	//	g.Text.Draw("You collected "+ strconv.Itoa(
	//
	// include
	// - number of requests addressed
	// - expenditures
	// - rent collected (from how many tenants of how many occupied units)
	// - current balance
	// - number of vacancies, move-ins, move-outs // not currently tracked
	// - any emergent events (inspections, etc) // not currently tracked
	// - building reputation increase/decrease, if any // delta not currently tracked
}

func (g *Game) AdvanceDay(t int) {
	day := -g.Clock.Days / 2
	g.Clock.Days += t
	day += (g.Clock.Days / 2)
	for _, r := range g.Building.Requests {
		r.DaysOpen += day
		if r.Closed && !r.Resolved {
			if r.AvailableSolutionsCount() == 0 {
				r.Resolved = true
			} else if r.ResolutionQuality <= 3 {
				if !strings.HasPrefix(r.Title, "(Reopened)") {
					r.Title = "(Reopened) " + r.Title
					r.Description = "I reopened this request because it was not actually resolved. " + r.Description
				}
				r.Closed = false
				r.Tenant.Satisfaction -= 1
			} else if r.ResolutionQuality <= 6 {
				// solution efficacy over time: good or bad?
				change := rand.IntN(5) - 2
				r.ResolutionQuality += change
			} else if r.ResolutionQuality >= 7 {
				r.Resolved = true
				r.Tenant.Satisfaction += 1
				g.Building.RequestsAddressed += 1
			}
		}
	}
	// increment LastOpened against the cooldown period
	for _, r := range g.RequestPool {
		if r.LastOpened < r.Cooldown {
			r.LastOpened += t / 3
		}
	}
	if g.Clock.Days >= 60 {
		g.State = "monthReport"
	}
}

func (g *Game) AdvanceDayByTicks() {
	if (g.Building.RequestsAddressed/(g.Clock.Days+1)) < 3 && g.Clock.Tick > 3000 {
		g.AdvanceDay(3)
		g.Clock.Tick = 0
	}
}

func (c *Clock) CheckEvents() {
	for key, value := range c.Recurring {
		if value[0] == value[1] {
			// TODO: incorporate events as requests
			fmt.Println(key + " : this event has been triggered")
			resetValue := value
			resetValue[1] = 0
			c.Recurring[key] = resetValue
		}
	}
}
