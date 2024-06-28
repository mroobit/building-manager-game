package main

import (
	"math/rand/v2"
)

func (g *Game) CreateProblems() {
	// Problems arise as a function of time with some randomness
	variance := 100 * (rand.IntN(7) + 2)
	if g.Clock.Tick%variance == 0 {
		g.GenerateRequest()
	}
}

func (g *Game) GenerateRequest() {
	if g.Building.Vacancies() == len(g.Building.Tenants) {
		return
	}
	t := NewTenant("", "", 0, 0, 0)
	for t.Name == "" {
		t = g.Building.Tenants[rand.IntN(len(g.Building.Tenants))]
	}
	blankRequest := &Request{
		Cooldown:   1,
		LastOpened: 0,
	}
	requestPoolIndex := 0
	c := 0
	for blankRequest.LastOpened < blankRequest.Cooldown {
		requestPoolIndex = rand.IntN(len(g.RequestPool))
		blankRequest = g.RequestPool[requestPoolIndex]
		c++
		if c > 4 {
			return
		}
	}
	g.RequestPool[requestPoolIndex].LastOpened = 0
	debugIndex = requestPoolIndex
	s := []Solution{}
	for _, sol := range blankRequest.Solutions {
		bs := Solution{
			Action:    sol.Action,
			Outcome:   sol.Outcome,
			Cost:      sol.Cost,
			Efficacy:  sol.Efficacy,
			Time:      sol.Time,
			Impact:    sol.Impact,
			Attempted: false,
		}
		s = append(s, bs)
	}
	r := Request{
		Title:       blankRequest.Title,
		Description: blankRequest.Description,
		Location:    blankRequest.Location,
		Tenant:      t,
		Urgent:      blankRequest.Urgent,
		Solutions:   s,
		DaysOpen:    0,
		Closed:      false,
		Resolved:    false,
	}

	if r.Location == "unit" {
		r.Location = t.Unit
	}
	r.Tenant = t
	g.Building.ReceiveRequest(&r)
}
