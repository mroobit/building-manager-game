package main

var (
	tenants             = make([]*Tenant, 10)
	initialSatisfaction = 7 // scale of 1-10
)

type Tenant struct {
	// TODO: give tenants names later, for more flavor
	Satisfaction    int
	Unit            string
	Rent            int
	MonthsRemaining int // this decrements regularly
	WillRenew       bool
}

func NewTenant(unit string, rent int, leaseLength int) *Tenant {
	tenant := &Tenant{
		Satisfaction:    initialSatisfaction,
		Unit:            unit,
		Rent:            rent,
		MonthsRemaining: leaseLength,
		WillRenew:       true,
	}
	return tenant
}

func initializeTenants(t []*Tenant) {
	// TODO: for each, generate dynamically (including satisfaction)
	//	alt: create a tenant-pool json, load pool of possible tenants, select some
	t[0] = NewTenant("1A", 850, 12)
	t[1] = NewTenant("1B", 800, 5)
	t[2] = NewTenant("1C", 800, 5)
	t[3] = NewTenant("1D", 800, 5)
	t[4] = NewTenant("1E", 800, 5)
	t[5] = NewTenant("1F", 800, 5)
	t[6] = NewTenant("1G", 800, 5)
	t[7] = NewTenant("1H", 800, 5)
	t[8] = NewTenant("1I", 800, 5)
	t[9] = NewTenant("1J", 800, 5)
}

func (t *Tenant) ReduceSatisfaction() {
	t.Satisfaction -= 1
	switch {
	case t.Satisfaction < 0:
		t.Satisfaction = 0
	case t.Satisfaction <= 3:
		t.WillRenew = false
	}
}

func (t *Tenant) IncreaseSatisfaction() {
	t.Satisfaction += 1
	switch {
	case t.Satisfaction > 10:
		t.Satisfaction = 10
	case t.Satisfaction >= 7:
		t.WillRenew = true
	}
}

// TODO: func (t *Tenant)Renew {}
// TODO: func (t *Tenant)UpdateMonthsRemaining(months int) {}
// This method's parameters can be positive or negative: positive for new lease, negative for time passing, early moveout, eviction
// TODO: func (t *Tenant).MoveOut(b *Building) & MoveInto
