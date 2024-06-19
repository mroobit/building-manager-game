package main

import "fmt"

type Clock struct {
	Tick      int
	Timer     int
	Days      int
	Month     int
	Recurring map[string][2]int
}

func (g *Game) initializeClock() {
	c := &Clock{
		Recurring: map[string][2]int{
			"Spray for bugs":    [2]int{3, 2},
			"Annual inspection": [2]int{12, 3},
		},
	}
	g.Clock = c
}

// update tenant months left
// update money in/out
// initiate events (prod to schedule spraying, annual inspection, infestation, etc)
// update months played
func (c *Clock) IncrementMonth() {
	c.Month += 1
	for key, value := range c.Recurring {
		update := value
		update[1] += 1
		c.Recurring[key] = update
	}
	c.CheckEvents()
	c.Days = 0
	c.Tick = 0
}

func (c *Clock) CheckDaysInMonth() {
	if c.Days >= 31 || c.Tick > 4000 {
		c.IncrementMonth()
	}
}

func (c *Clock) CheckEvents() {
	for key, value := range c.Recurring {
		if value[0] == value[1] {
			// TODO: incorporate events as requests
			fmt.Println(key + " : this event has been triggered")
			resetValue := value
			resetValue[1] = 0
			c.Recurring[key] = resetValue
		}
	}
}
