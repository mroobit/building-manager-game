package main

import (
	"embed"
	"encoding/json"
	"log"

	"github.com/google/uuid"
)

type Request struct {
	Title             string `json:"title"`
	Description       string `json:"description"`
	ID                uuid.UUID
	Location          string `json:"location"`
	Tenant            *Tenant
	Urgent            bool       `json:"urgent"`
	Cooldown          int        `json:"cooldown"`
	LastOpened        int        // only applies in RequestPool, decrement with day increment, check against cooldown
	Solutions         []Solution `json:"solutions"`
	Attempts          []int      // solutions that have been attempted
	DaysOpen          int        // this increments regularly
	Closed            bool       // requests can be closed without resolving
	Resolved          bool       // was the problem actually fixed
	ResolutionQuality int
}

type Solution struct {
	Action    string `json:"action"`
	Outcome   string `json:"outcome"` // story followup // TODO: make this a []string with possible outcomes?
	Cost      int    `json:"cost"`
	Efficacy  int    `json:"efficacy"`
	Time      int    `json:"time"`
	Impact    int    `json:"impact"` // impact on tenant satisfaction
	Attempted bool
}

type wrapper struct {
	Requests []*Request `json:"requests"`
}

func NewRequest(title string, description string, location string, tenant *Tenant, urgent bool) *Request {
	r := &Request{
		Title:       title,
		Description: description,
		Location:    location,
		Tenant:      tenant,
		Urgent:      urgent,
		DaysOpen:    0,
		Closed:      false,
		Resolved:    false,
	}
	return r
}

func (r *Request) Resolve(solutionIndex int) (cost, time int) {
	r.Solutions[solutionIndex].Attempted = true
	solution := r.Solutions[solutionIndex]
	r.Attempts = append(r.Attempts, solutionIndex)
	r.ResolutionQuality = solution.Efficacy
	if r.ResolutionQuality >= 7 {
		r.Resolved = true
	}
	r.Tenant.Impact(solution.Impact)
	r.Closed = true
	return solution.Cost, solution.Time
}

func (g *Game) initializeRequestPool(fs embed.FS) {
	var rawRequestPool []*Request
	requestData, err := fs.ReadFile("data/requests.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	err = json.Unmarshal(requestData, &rawRequestPool)
	if err != nil {
		log.Fatal("Error when unmarshalling: ", err)
	}

	g.RequestPool = rawRequestPool
	for _, r := range g.RequestPool {
		r.LastOpened = r.Cooldown
	}
}

func (r *Request) AvailableSolutionIndices() []int {
	indices := []int{}
	for i, s := range r.Solutions {
		if !s.Attempted {
			indices = append(indices, i)
		}
	}
	return indices
}

func (r *Request) AvailableSolutionsCount() int {
	count := 0
	for _, s := range r.Solutions {
		if !s.Attempted {
			count += 1
		}
	}
	return count
}

// TODO unmarshal JSON into array of all possible requests
// TODO second JSON for consequences/escalations?
