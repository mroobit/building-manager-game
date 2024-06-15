package main

var (
	requests []*Request
)

type Request struct {
	Title             string
	Description       string
	Location          string
	Tenant            *Tenant
	Urgent            bool
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
	r.Resolved = true
	r.ResolutionQuality = quality
	r.Closed = true
}
