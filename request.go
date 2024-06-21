package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
)

type Request struct {
	Title             string `json:"title"`
	Description       string `json:"description"`
	Location          string `json:"location"`
	Tenant            *Tenant
	Urgent            bool       `json:"urgent"`
	Solutions         []Solution `json:"solutions"`
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

func (r *Request) Resolve(option int) (cost int, time int) {
	// TODO: implement Options []*Solution in Request struct
	// then add logic for assigning resolution quality to request when resolving based on chosen option
	fmt.Println("SOLUTIONS")
	fmt.Println(r.Solutions)
	r.Resolved = true
	//	r.ResolutionQuality = r.Options[option].Efficacy
	// r.Closed = true
	//return r.Solutions[option].Cost, r.Solutions[option].Time
	return 50, 11
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
