package main

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
	cRequestList := NewClickable([2]int{0, 160}, [2]int{360, 200}, "request-list", "Display Open Maintenance Requests")
	portalClickable["request-list"] = cRequestList
	cFinancial := NewClickable([2]int{0, 220}, [2]int{360, 260}, "financial-overview", "Display Financial Overview")
	portalClickable["financial-overview"] = cFinancial
}
