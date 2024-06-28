package main

import "github.com/hajimehoshi/ebiten/v2"

const (
	sidebarButtonLeftX     = 0
	sidebarButtonRightX    = int(portalSidebarWidth)
	sidebarHeightIncrement = 60
	requestListXMargin     = 40
)

var (
	button map[string]*Clickable

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
	button = make(map[string]*Clickable)
	cPlay := NewClickable(
		[2]int{100, 300},
		[2]int{400, 360},
		"play",
		"Start Game",
	)
	button["play"] = cPlay
	cControls := NewClickable(
		[2]int{100, 400},
		[2]int{400, 465},
		"controls",
		"Display Control Info",
	)
	button["controls"] = cControls

	// Login Buttons
	cLoginPlay := NewClickable(
		[2]int{460, 670},
		[2]int{820, 710},
		"login-play",
		"Play Game",
	)
	button["login-play"] = cLoginPlay
	cHowToPlay := NewClickable(
		[2]int{460, 715},
		[2]int{820, 785},
		"how-to-play",
		"Learn how to play",
	)
	button["how-to-play"] = cHowToPlay
	cSettings := NewClickable(
		[2]int{480, 850},
		[2]int{620, 900},
		"settings",
		"Adjust settings",
	)
	button["settings"] = cSettings
	cAbout := NewClickable(
		[2]int{660, 850},
		[2]int{800, 900},
		"about",
		"About the game",
	)
	button["about"] = cAbout
	cUpperX := NewClickable(
		[2]int{1220, 20},
		[2]int{1260, 60},
		"upper-x",
		"Return to previous screen",
	)
	button["upper-x"] = cUpperX

	// Portal Buttons
	cOverview := NewClickable(
		[2]int{sidebarButtonLeftX, sidebarButtonLeftY},
		[2]int{sidebarButtonRightX, sidebarButtonRightY},
		"overview",
		"Display Portal Overview",
	)
	button["overview"] = cOverview
	sidebarButtonLeftY += sidebarHeightIncrement
	sidebarButtonRightY += sidebarHeightIncrement
	cRequestList := NewClickable(
		[2]int{sidebarButtonLeftX, sidebarButtonLeftY},
		[2]int{sidebarButtonRightX, sidebarButtonRightY},
		"request-list",
		"Display Open Tenant Requests",
	)
	button["request-list"] = cRequestList
	sidebarButtonLeftY += sidebarHeightIncrement
	sidebarButtonRightY += sidebarHeightIncrement
	cTenant := NewClickable(
		[2]int{sidebarButtonLeftX, sidebarButtonLeftY},
		[2]int{sidebarButtonRightX, sidebarButtonRightY},
		"tenants",
		"Display Tenant Overview",
	)
	button["tenants"] = cTenant
	sidebarButtonLeftY += sidebarHeightIncrement
	sidebarButtonRightY += sidebarHeightIncrement
	cFinancial := NewClickable(
		[2]int{sidebarButtonLeftX, sidebarButtonLeftY},
		[2]int{sidebarButtonRightX, sidebarButtonRightY},
		"financial-overview",
		"Display Financial Overview",
	)
	button["financial-overview"] = cFinancial

	cDetails := NewClickable(
		[2]int{int(portalSidebarWidth) + requestListXMargin, 200},
		[2]int{1280 - requestListXMargin, 845},
		"request-details",
		"Display Individual Tenant Request",
	)
	button["request-details"] = cDetails
	cResolve := NewClickable(
		[2]int{530, 500},
		[2]int{800, 570},
		"try-to-resolve",
		"Display Possible Solutions",
	)
	button["try-to-resolve"] = cResolve
	cClose := NewClickable(
		[2]int{830, 500},
		[2]int{1100, 570},
		"close-request",
		"Close Active Request",
	)
	button["close-request"] = cClose

	cSolutions := NewClickable(
		[2]int{400, 480},
		[2]int{1250, 850},
		"solutions",
		"Resolve Request with Selected Solution",
	)
	button["solutions"] = cSolutions

	// General "continue" and "back" buttons
	cContinue := NewClickable(
		[2]int{460, 795},
		[2]int{820, 865},
		"continue",
		"Continue to next screen",
	)
	button["continue"] = cContinue
	cBack := NewClickable(
		[2]int{460, 715},
		[2]int{820, 800},
		"back",
		"Return to previous screen",
	)
	button["back"] = cBack

}

func (c *Clickable) DrawHoverEffect(screen ebiten.Image) {
	// TODO: use clickable dimensions to draw hover effect
}
