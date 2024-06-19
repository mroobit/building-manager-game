package main

import (
	"fmt"
	"strconv"
)

var (
	inspectionCycle     = 12 // frequency of inspections in months
	monthlyBuildingCost = 1000
)

type Building struct {
	Manager       string
	Money         int
	CreditBalance int
	Reputation    int
	Tenants       []*Tenant
	Requests      []*Request
	FixedCosts    int // monthly cost, will increase over time
	Inspection    int // months until next inspection
}

func (g *Game) initializeBuilding() {
	r := make([]*Request, 0, 30)

	g.Building = &Building{
		Money:      1000,
		Reputation: 7,
		Tenants:    tenants,
		Requests:   r,
		FixedCosts: monthlyBuildingCost,
		Inspection: 10,
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

func (b *Building) ReceiveRequest(r *Request) {
	// TODO: Also trigger notification if player is not on request list screen
	if r.Urgent {
		b.Requests = append([]*Request{r}, b.Requests...)
	} else {
		b.Requests = append(b.Requests, r)
	}
}

func (b *Building) ReopenRequests() {
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

func (b *Building) OpenRequestCount() int {
	count := 0
	for _, r := range b.Requests {
		if !r.Closed {
			count += 1
		}
	}

	return count
}

// TODO: func (b *Building)Vacancies {} -- reports which units are vacant
