package main

import "github.com/hajimehoshi/ebiten/v2"

const (
	sidebarButtonLeftX     = 0
	sidebarButtonRightX    = int(portalSidebarWidth)
	sidebarHeightIncrement = 60
	requestListXMargin     = 40
)

var (
	titleButton  map[string]*Clickable
	portalButton map[string]*Clickable

	sidebarButtonLeftY  = 100
	sidebarButtonRightY = sidebarButtonLeftY + 40
)

type Clickable struct {
	UpperLeft  [2]int
	LowerRight [2]int
	Width      int
	Height     int
	HoverText  string
	ClickText  string
}

func NewClickable(ul [2]int, lr [2]int, hover string, click string) *Clickable {
	w := lr[0] - ul[0]
	h := lr[1] - ul[1]
	c := &Clickable{
		UpperLeft:  ul,
		LowerRight: lr,
		Width:      w,
		Height:     h,
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
	titleButton = make(map[string]*Clickable)
	cPlay := NewClickable(
		[2]int{100, 300},
		[2]int{400, 360},
		"play",
		"Start Game",
	)
	titleButton["play"] = cPlay
	cControls := NewClickable(
		[2]int{100, 400},
		[2]int{400, 465},
		"controls",
		"Display Control Info",
	)
	titleButton["controls"] = cControls

	portalButton = make(map[string]*Clickable)
	cOverview := NewClickable(
		[2]int{sidebarButtonLeftX, sidebarButtonLeftY},
		[2]int{sidebarButtonRightX, sidebarButtonRightY},
		"overview",
		"Display Portal Overview",
	)
	portalButton["overview"] = cOverview
	sidebarButtonLeftY += sidebarHeightIncrement
	sidebarButtonRightY += sidebarHeightIncrement
	cRequestList := NewClickable(
		[2]int{sidebarButtonLeftX, sidebarButtonLeftY},
		[2]int{sidebarButtonRightX, sidebarButtonRightY},
		"request-list",
		"Display Open Tenant Requests",
	)
	portalButton["request-list"] = cRequestList
	sidebarButtonLeftY += sidebarHeightIncrement
	sidebarButtonRightY += sidebarHeightIncrement
	cTenant := NewClickable(
		[2]int{sidebarButtonLeftX, sidebarButtonLeftY},
		[2]int{sidebarButtonRightX, sidebarButtonRightY},
		"tenants",
		"Display Tenant Overview",
	)
	portalButton["tenants"] = cTenant
	sidebarButtonLeftY += sidebarHeightIncrement
	sidebarButtonRightY += sidebarHeightIncrement
	cFinancial := NewClickable(
		[2]int{sidebarButtonLeftX, sidebarButtonLeftY},
		[2]int{sidebarButtonRightX, sidebarButtonRightY},
		"financial-overview",
		"Display Financial Overview",
	)
	portalButton["financial-overview"] = cFinancial

	cDetails := NewClickable(
		[2]int{int(portalSidebarWidth) + requestListXMargin, 200},
		[2]int{1280 - requestListXMargin, 845},
		"request-details",
		"Display Individual Tenant Request",
	)
	portalButton["request-details"] = cDetails
	cResolve := NewClickable(
		[2]int{530, 400},
		[2]int{800, 470},
		"try-to-resolve",
		"Display Possible Solutions",
	)
	portalButton["try-to-resolve"] = cResolve
	cClose := NewClickable(
		[2]int{830, 400},
		[2]int{1100, 470},
		"close-request",
		"Close Active Request",
	)
	portalButton["close-request"] = cClose

	cSolutions := NewClickable(
		[2]int{400, 430},
		[2]int{1250, 850},
		"solutions",
		"Resolve Request with Selected Solution",
	)
	portalButton["solutions"] = cSolutions

}

func (c *Clickable) DrawHoverEffect(screen ebiten.Image) {
	// TODO: use clickable dimensions to draw hover effect
}
