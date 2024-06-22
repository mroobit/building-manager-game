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
	Solutions         []Solution `json:"solutions"`
	Attempts          []string   // solutions that have been attempted
	DaysOpen          int        // this increments regularly
	Closed            bool       // requests can be closed without resolving
	Resolved          bool       // was the problem actually fixed
	ResolutionQuality int
}

type Solution struct {
	Action   string `json:"action"`
	Cost     int    `json:"cost"`
	Efficacy int    `json:"efficacy"`
	Time     int    `json:"time"`
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

func (r *Request) Close() {
	r.Closed = true
}

func (r *Request) Resolve(selection int) (cost, time int) {
	r.Attempts = append(r.Attempts, r.Solutions[selection].Action)
	r.Resolved = true
	r.ResolutionQuality = r.Solutions[selection].Efficacy
	// r.Closed = true
	return r.Solutions[selection].Cost, r.Solutions[selection].Time
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
}

// TODO unmarshal JSON into array of all possible requests
// TODO second JSON for consequences/escalations?
