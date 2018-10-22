package model

type Passenger interface {
	GetDocuments() []interface{}
	GetTicket() *Ticket
}

type American struct {
	Passport *AmericanPassport
	Ticket   *Ticket
}

func (a *American) GetDocuments() []interface{} {
	return []interface{}{a.Passport}
}

func (a *American) GetTicket() *Ticket {
	return a.Ticket
}

type Russian struct {
	BasePassport          *RussianBasePassport
	InternationalPassport *RussianInternationalPassport
	Ticket                *Ticket
}

func (r *Russian) GetDocuments() []interface{} {
	return []interface{}{r.BasePassport, r.InternationalPassport}
}

func (r *Russian) GetTicket() *Ticket {
	return r.Ticket
}
