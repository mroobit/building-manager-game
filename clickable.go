package main

import "github.com/hajimehoshi/ebiten/v2"

var (
	titleClickable  map[string]*Clickable
	portalClickable map[string]*Clickable
)

type Clickable struct {
	UpperLeft  [2]int
	LowerRight [2]int
	HoverText  string
	ClickText  string
}

func NewClickable(ul [2]int, lr [2]int, hover string, click string) *Clickable {
	c := &Clickable{
		UpperLeft:  ul,
		LowerRight: lr,
		HoverText:  hover,
		ClickText:  click,
	}
	return c
}

func (c *Clickable) Hover(cursor [2]int) bool {
	hovering := (cursor[0] >= c.UpperLeft[0] && cursor[0] <= c.LowerRight[0]) &&
		(cursor[1] >= c.UpperLeft[1] && cursor[1] <= c.LowerRight[1])
	return hovering
}

func initializeClickables() {
	// TODO: load from json/csv/whatever later
	titleClickable = make(map[string]*Clickable)
	cPlay := NewClickable([2]int{100, 300}, [2]int{400, 360}, "play", "Start Game")
	titleClickable["play"] = cPlay
	cControls := NewClickable([2]int{100, 400}, [2]int{400, 465}, "controls", "Display Control Info")
	titleClickable["controls"] = cControls

	portalClickable = make(map[string]*Clickable)
	cOverview := NewClickable([2]int{0, 100}, [2]int{360, 140}, "overview", "Display Portal Overview")
	portalClickable["overview"] = cOverview
	cRequestList := NewClickable([2]int{0, 160}, [2]int{360, 200}, "request-list", "Display Open Tenant Requests")
	portalClickable["request-list"] = cRequestList
	cTenant := NewClickable([2]int{0, 220}, [2]int{360, 260}, "financial-overview", "Display Financial Overview")
	portalClickable["tenants"] = cTenant
	cFinancial := NewClickable([2]int{0, 280}, [2]int{360, 320}, "financial-overview", "Display Financial Overview")
	portalClickable["financial-overview"] = cFinancial

	cDetails := NewClickable([2]int{390, 200}, [2]int{1240, 845}, "request-details", "Display Individual Tenant Request")
	portalClickable["request-details"] = cDetails
	cResolve := NewClickable([2]int{530, 400}, [2]int{800, 470}, "try-to-resolve", "Display Possible Solutions")
	portalClickable["try-to-resolve"] = cResolve
	cClose := NewClickable([2]int{830, 400}, [2]int{1000, 470}, "close-request", "Close Active Request")
	portalClickable["close-request"] = cClose
}

func (c *Clickable) DrawHoverEffect(screen ebiten.Image) {
	// TODO: use clickable dimensions to draw hover effect
}
