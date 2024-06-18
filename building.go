package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
)

var (
	inspectionCycle     = 12 // frequency of inspections in months
	monthlyBuildingCost = 1000
)

type Building struct {
	Tenants     []*Tenant
	Requests    []*Request
	Maintenance int // monthly cost
	Inspection  int // months until next inspection
}

func (g *Game) initializeBuilding() {
	r := make([]*Request, 0, 30)

	g.Building = &Building{
		Tenants:     tenants,
		Requests:    r,
		Maintenance: monthlyBuildingCost,
		Inspection:  10,
	}
}

// TODO: (stretch) make this more useable for display as tenant roster
func (b *Building) ListTenants() {
	for _, t := range b.Tenants {
		fmt.Println("Tenant in Unit " + t.Unit)
		fmt.Println(" - Rent: " + strconv.Itoa(t.Rent))
		fmt.Println(" - Satisfaction: " + strconv.Itoa(t.Satisfaction))
		fmt.Println(" - Months Left: " + strconv.Itoa(t.MonthsRemaining))
		fmt.Println(" - Will Renew: " + strconv.FormatBool(t.WillRenew))
	}
}

func (b *Building) GenerateRequest(pool []*Request) {
	t := b.Tenants[rand.IntN(len(b.Tenants))]
	r := pool[rand.IntN(len(pool))]
	if r.Location == "unit" {
		r.Location = t.Unit
	}
	r.Tenant = t
	b.AddRequest(r)
}

func (b *Building) AddRequest(request *Request) {
	b.Requests = append(b.Requests, request)
}

func (b *Building) ReviveRequests() {
	for _, r := range b.Requests {
		if !r.Resolved {
			r.Closed = false
			r.Urgent = true
			r.Tenant.ReduceSatisfaction()
		}
		// TODO: Reduce quality of resolution on poor-solution requests until must reopen
		// Have some such requests actually get fully-resolved
	}
}

// TODO: func (b *Building)Vacancies {} -- reports which units are vacant
