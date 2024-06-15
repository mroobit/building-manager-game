package main

import (
	"fmt"
	"strconv"
)

var (
	building            *Building
	inspectionCycle     = 12 // frequency of inspections in months
	monthlyBuildingCost = 1000
)

type Building struct {
	Tenants     []*Tenant
	Requests    []*Request
	Maintenance int // monthly cost
	Inspection  int // months until next inspection
}

func initializeBuilding() {
	r := make([]*Request, 0, 30)

	building = &Building{
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

func (b *Building) GenerateRequest() {
	// get request details from requests json/similar
	// select a tenant to assign
	// r := NewRequest( // use grabbed details for this
}

// TODO: func (b *Building)Vacancies {} -- reports which units are vacant
