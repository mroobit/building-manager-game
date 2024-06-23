package main

var (
	tenants             = make([]*Tenant, 10)
	initialSatisfaction = 7 // scale of 1-10
	defaultLeaseLength  = 12
)

type Tenant struct {
	// TODO: add max rent a tenant is willing to pay before they'll move out
	Name            string
	Satisfaction    int
	Unit            string
	Rent            int
	MonthsRemaining int // this decrements regularly
	WillRenew       bool
}

func NewTenant(name string, unit string, rent int, leaseLength int) *Tenant {
	tenant := &Tenant{
		Name:            name,
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
	t[0] = NewTenant("Jason Ellingsworth", "1A", 850, 12)
	t[1] = NewTenant("Victoria Kent", "1B", 800, 5)
	t[2] = NewTenant("Winnie Lopez", "1C", 800, 5)
	t[3] = NewTenant("", "1D", 800, 5)
	t[4] = NewTenant("Grace Holtz", "1E", 800, 5)
	t[5] = NewTenant("Andre Svenson", "1F", 800, 5)
	t[6] = NewTenant("Fiona Phelps", "1G", 800, 5)
	t[7] = NewTenant("Maxine Flaherty", "1H", 800, 5)
	t[8] = NewTenant("Harriet Su", "1I", 800, 5)
	t[9] = NewTenant("LeAnne Smith", "1J", 800, 5)
}

// TODO: func (t *Tenant)Renew {}
// TODO: func (t *Tenant)UpdateMonthsRemaining(months int) {}
// This method's parameters can be positive or negative: positive for new lease, negative for time passing, early moveout, eviction
// TODO: func (t *Tenant).MoveOut(b *Building) & MoveInto

func (t *Tenant) Impact(i int) {
	t.Satisfaction += i
	switch {
	case t.Satisfaction < 0:
		t.Satisfaction = 0
	case t.Satisfaction <= 3:
		t.WillRenew = false
	case t.Satisfaction > 10:
		t.Satisfaction = 10
	case t.Satisfaction >= 7:
		t.WillRenew = true
	}
}

func (t *Tenant) Needed() bool {
	return t.Name == ""
}

func (t *Tenant) MoveOut() {
	t.Name = ""
	// increase rent for listing price
}

func (t *Tenant) MoveIn() {
	// TODO: pull tenant from tenant pool
	// if tenant.MaxRent < unit rent, assign to the empty unit
}

// TODO loadTenantPool from json
