package main

import "math/rand/v2"

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
	t := NewTenant("", "", 0, 0)
	for t.Name == "" {
		t = g.Building.Tenants[rand.IntN(len(g.Building.Tenants))]
	}
	blankRequest := g.RequestPool[rand.IntN(len(g.RequestPool))]
	r := Request{
		Title:       blankRequest.Title,
		Description: blankRequest.Description,
		Location:    blankRequest.Location,
		Tenant:      t,
		Urgent:      blankRequest.Urgent,
		Solutions:   blankRequest.Solutions,
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

/*
func (g *Game) RentRenew() (nonpaying, moving []*Tenant) {
	// iterate tenants
	//	- add rent to player money
	// 		- populate non-payment group (or on-Tenant set "paid" to false)
	// 	- renewals/move-outs
	//		- update unit vacancies

	// for _, t := range g.Complex.Tenants {
	//	(unless obstructive reason) g.Player.Money += t.Rent
	//
	return
}
*/
