package main

import (
	"embed"
	"encoding/json"
	"log"
)

type Request struct {
	Title       string
	Description string
	Location    string
	Tenant      *Tenant
	Urgent      bool
	//	Options           []*Solutions
	DaysOpen          int  // this increments regularly
	Closed            bool // requests can be closed without resolving
	Resolved          bool // was the problem actually fixed
	ResolutionQuality int
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

func (r *Request) Resolve(quality int) {
	// TODO: implement Options []*Solution in Request struct
	// then add logic for assigning resolution quality to request when resolving based on chosen option
	r.Resolved = true
	r.ResolutionQuality = quality
	r.Closed = true
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
