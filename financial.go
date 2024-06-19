package main

func valueSign(money int) string {
	if money > 0 {
		return "+"
	} else if money < 0 {
		return "-"
	}
	return ""
}

func (g *Game) UpcomingRent() int {
	income := 0
	for _, t := range g.Building.Tenants {
		income += t.Rent
	}
	return income
}

// Upcoming Payments is Net Payments (Rent income, CC payment, Fixed costs)
func (g *Game) UpcomingPayments() int {
	payment := g.UpcomingRent() - g.Building.CreditBalance - g.Building.FixedCosts
	return payment
}

// ProcessPayments add rent income to building money and subtracts fixed costs and CC payment (until money reaches zero)
func (g *Game) ProcessPayments() {
	g.Building.Money += g.CollectRent()
	g.Building.Money -= g.Building.FixedCosts
	if g.Building.Money < g.Building.CreditBalance {
		g.Building.CreditBalance -= g.Building.Money
		g.Building.Money = 0
	} else {
		g.Building.Money -= g.Building.CreditBalance
		g.Building.CreditBalance = 0
	}
}

// TODO: logic to have tenants occasionally unable (eg tenant death) or unwilling (withholding to to lack of repair) rent
func (g *Game) CollectRent() int {
	income := 0
	for _, t := range g.Building.Tenants {
		income += t.Rent
	}
	return income
}
