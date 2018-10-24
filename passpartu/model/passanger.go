package model

type Passenger interface {
	GetDocuments() []interface{}
	GetTicket() *Ticket
	GetPass() *Pass
}

type ETicketHolder struct {
	Pass *Pass
}

func (me *ETicketHolder) GetPass() *Pass {
	return me.Pass
}

type American struct {
	ETicketHolder
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
	ETicketHolder
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
