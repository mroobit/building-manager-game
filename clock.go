package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
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
	g.Clock.CheckEvents()
	g.Clock.Days = 0
	g.Clock.Tick = 0
}

func (g *Game) AdvanceDay(t int) {
	day := -g.Clock.Days / 3
	g.Clock.Days += t
	day += (g.Clock.Days / 3)
	for _, r := range g.Building.Requests {
		r.DaysOpen += day
		if r.Closed && !r.Resolved {
			if r.ResolutionQuality <= 3 {
				if !strings.HasPrefix(r.Title, "(Reopened)") {
					r.Title = "(Reopened) " + r.Title
					r.Description = "I reopened this request because it was not actually resolved. " + r.Description
				}
				r.Closed = false
			} else if r.ResolutionQuality <= 6 {
				// solution efficacy over time: good or bad?
				change := rand.IntN(5) - 2
				r.ResolutionQuality += change
			} else if r.ResolutionQuality >= 7 {
				r.Resolved = true
			}
		}
	}
	if g.Clock.Days >= 90 {
		g.IncrementMonth()
	}
}

func (g *Game) AdvanceDayByTicks() {
	if g.Building.RequestsAddressed < 3 && g.Clock.Tick > 3000 {
		g.AdvanceDay(3)
		g.Clock.Tick = 0
	}
}

func (g *Game) CheckDaysInMonth() {
	if g.Clock.Days >= 93 {
		g.IncrementMonth()
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
