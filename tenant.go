package main

import (
	"bufio"
	"embed"
	"log"
	"math/rand/v2"
)

var (
	initialSatisfaction = 7 // scale of 1-10
	defaultLeaseLength  = 12
)

type Tenant struct {
	// TODO: add max rent a tenant is willing to pay before they'll move out
	Name            string
	Satisfaction    int
	Unit            string
	Rent            int
	MaxRent         int
	MonthsRemaining int // this decrements regularly
	WillRenew       bool
}

func NewTenant(name string, unit string, rent int, maxRent int, leaseLength int) *Tenant {
	tenant := &Tenant{
		Name:            name,
		Satisfaction:    initialSatisfaction,
		Unit:            unit,
		Rent:            rent,
		MaxRent:         maxRent,
		MonthsRemaining: leaseLength,
		WillRenew:       true,
	}
	return tenant
}

func (g *Game) initializeTenantPool(fs embed.FS) {
	tenants := []*Tenant{}

	file, err := fs.Open("data/tenants.txt")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rent := g.Building.BaseRent + 10*(rand.IntN(12))
		leaseLength := rand.IntN(7) + 3
		maxRent := rent + 10*(rand.IntN(20)+5)
		name := scanner.Text()
		t := NewTenant(name, "", rent, maxRent, leaseLength)
		tenants = append(tenants, t)
	}

	g.TenantPool = tenants
}

func (g *Game) initializeTenants() {
	tenants := []*Tenant{
		&Tenant{Unit: "1A"},
		&Tenant{Unit: "1B"},
		&Tenant{Unit: "1C"},
		&Tenant{Unit: "1D"},
		&Tenant{Unit: "2A"},
		&Tenant{Unit: "2B"},
		&Tenant{Unit: "2C"},
		&Tenant{Unit: "2D"},
		&Tenant{Unit: "3A"},
		&Tenant{Unit: "3B"},
	}

	for i, t := range tenants {
		index := rand.IntN(len(g.TenantPool))
		for g.TenantPool[index].Unit != "" {
			index = rand.IntN(len(g.TenantPool))
		}
		g.TenantPool[index].Unit = t.Unit
		tenants[i] = g.TenantPool[index]
	}
	g.Building.Tenants = tenants
}

// TODO: func (t *Tenant)Renew {}
// TODO: func (t *Tenant)UpdateMonthsRemaining(months int) {}
// This method's parameters can be positive or negative: positive for new lease, negative for time passing, early moveout, eviction
// TODO: func (t *Tenant).MoveOut(b *Building) & MoveInto

func (t *Tenant) Impact(i int) {
	t.Satisfaction += i
	switch {
	case t.Satisfaction < 0:
		t.Satisfaction = 0
	case t.Satisfaction <= 3:
		t.WillRenew = false
	case t.Satisfaction > 10:
		t.Satisfaction = 10
	case t.Satisfaction >= 7:
		t.WillRenew = true
	}
}

func (t *Tenant) Needed() bool {
	return t.Name == ""
}

func (t *Tenant) MoveOut() {
	t.Name = ""
	// increase rent for listing price
}

func (t *Tenant) MoveIn() {
	// TODO: pull tenant from tenant pool
	// if tenant.MaxRent < unit rent, assign to the empty unit
}

// TODO loadTenantPool from json
