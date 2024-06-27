package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"

	"github.com/google/uuid"
)

var (
	inspectionCycle = 12 // frequency of inspections in months
)

type Building struct {
	Manager           string
	Money             int
	CreditBalance     int
	Reputation        int
	BaseRent          int
	RentIncrease      float32
	Tenants           []*Tenant
	Requests          []*Request
	RequestMap        map[uuid.UUID]*Request
	ActiveRequest     *Request
	RequestsAddressed int
	FixedCosts        int // monthly cost, will increase over time
	Inspection        int // months until next inspection
}

func (g *Game) initializeBuilding() {
	r := make([]*Request, 0, 30)
	m := make(map[uuid.UUID]*Request)
	t := make([]*Tenant, 10, 10)

	initialMoney := 10 * (rand.IntN(50) + 30)
	initialReputation := rand.IntN(4) + 4
	initialFixedCosts := 10 * (rand.IntN(50) + 100)
	baseRent := 10 * (rand.IntN(40) + 40)
	rentIncrease := float32(0.03)

	g.Building = &Building{
		Money:        initialMoney,
		Reputation:   initialReputation,
		BaseRent:     baseRent,
		RentIncrease: rentIncrease,
		Tenants:      t,
		Requests:     r,
		RequestMap:   m,
		FixedCosts:   initialFixedCosts,
		Inspection:   10,
	}
}

func (b *Building) ReceiveRequest(r *Request) {
	r.ID = uuid.New()
	b.RequestMap[r.ID] = r
	if r.Urgent {
		b.Requests = append([]*Request{r}, b.Requests...)
	} else {
		b.Requests = append(b.Requests, r)
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

// OpenIndices returns a slice of the indices of open requests
func (b *Building) OpenIndices() []int {
	indices := []int{}
	for i, r := range b.Requests {
		if !r.Closed {
			indices = append(indices, i)
		}
	}
	return indices
}

func (b *Building) Vacancies() int {
	count := 0
	for _, t := range b.Tenants {
		fmt.Println("Tenant - name, unit: " + t.Name + ", " + t.Unit)
		if t.Name == "" {
			fmt.Println("Name not empty!")
			count += 1
		}
	}
	fmt.Println("Count: " + strconv.Itoa(count))
	return count
}

// VacanciesList returns a slice of the indices of empty units
func (b *Building) VacanciesList() []int {
	indices := []int{}
	for i, t := range b.Tenants {
		if t.Name != "" {
			indices = append(indices, i)
		}
	}
	return indices
}

func (b *Building) DecrementLeases() {
	for i, t := range b.Tenants {
		t.MonthsRemaining--
		if t.MonthsRemaining == 0 {
			b.Renew(i)
		}
	}
}

func (b *Building) Renew(t int) {
	newRent := int(float32(b.Tenants[t].Rent) * (1.00 + b.RentIncrease))
	if b.Tenants[t].Satisfaction < 3 || b.Tenants[t].MaxRent < newRent {
		b.Tenants[t] = &Tenant{
			Unit: b.Tenants[t].Unit,
		}
	} else if b.Tenants[t].Satisfaction < 6 {
		// function of satisfaction, maxRent willing to pay
		if newRent-b.Tenants[t].Rent <= (b.Tenants[t].MaxRent-b.Tenants[t].Rent)*b.Tenants[t].Satisfaction/10 {
			b.Tenants[t].MonthsRemaining = 12
		} else {
			b.Tenants[t] = &Tenant{
				Unit: b.Tenants[t].Unit,
			}
		}
	} else {
		b.Tenants[t].MonthsRemaining = 12
	}
}
