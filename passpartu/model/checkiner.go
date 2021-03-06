package model

import (
	"fmt"
	"time"
)

type PassRegistry interface {
	GetTicketId(pass *Pass) string
	SetPass(pass *Pass, ticketId string)
}

type InMemoryPassRegistry struct {
	PassRegistry map[string]string
}

func (registry *InMemoryPassRegistry) GetTicketId(pass *Pass) string {
	return registry.PassRegistry[pass.ID]
}

func (registry *InMemoryPassRegistry) SetPass(pass *Pass, ticketId string) {
	registry.PassRegistry[pass.ID] = ticketId
}

type CheckIner interface {
	CheckIn(passenger Passenger) (*Pass, error)
}

func checkin(passenger Passenger) (*Pass, error) {

	var internationalPassport InternationalPassport
	var ok bool
	ticket := passenger.GetTicket()

	for _, document := range passenger.GetDocuments() {
		internationalPassport, ok = document.(InternationalPassport)
		if ok {
			break
		}
	}

	if internationalPassport == nil {
		return nil, fmt.Errorf("international passport not found")
	}

	if internationalPassport.GetName() == "" || internationalPassport.GetCredentials() == "" {
		return nil, fmt.Errorf("international passport is invalid")
	}

	if (ticket.Destination.Country == "USA" || ticket.Destination.Country == "UAE") && internationalPassport.GetSex() == SexFemale {
		var basePassport BasePassport
		for _, document := range passenger.GetDocuments() {
			basePassport, ok = document.(BasePassport)
			if ok {
				break
			}
		}
		if basePassport == nil {
			return nil, fmt.Errorf("base passport is invalid")
		}

		if basePassport.GetName() == "" || basePassport.GetCredentials() == "" {
			return nil, fmt.Errorf("base passport is invalid")
		}

		if !basePassport.GetMarried() {
			return nil, fmt.Errorf("bb go home")
		}
	}

	for i, pet := range ticket.Pets {
		if pet.Passport == nil {
			return nil, fmt.Errorf("no Passport for animal #%v", i)
		}
		if pet.OwnershipCertificate == nil {
			return nil, fmt.Errorf("no OwnershipCertificate for animal #%v", i)
		}
		if pet.Weight > 40 && pet.SafetyCertificate == nil {
			return nil, fmt.Errorf("no SafetyCertificate for heavy animal #%v", i)
		}
	}

	return &Pass{
		Passenger: passenger,
		ID:        fmt.Sprintf("Some Unique Identifier %v", time.Now()),
	}, nil

}

func CheckInPassenger(passenger Passenger) (*Pass, error) {

	if passenger.GetTicket().Pets != nil || len(passenger.GetTicket().Pets) > 0 {
		return nil, fmt.Errorf("animal checkin not allowed")
	}

	return checkin(passenger)
}

func CheckInPassangerWithAnimal(passenger Passenger) (*Pass, error) {
	return checkin(passenger)
}

type Babka struct{}

func (b *Babka) CheckIn(passenger Passenger) (*Pass, error) {
	return CheckInPassenger(passenger)
}

type Detka struct{}

func (d *Detka) CheckIn(passenger Passenger) (*Pass, error) {
	return CheckInPassangerWithAnimal(passenger)
}

type SkynetCheckiner struct {
	PassRegistry PassRegistry
}

func (d *SkynetCheckiner) CheckIn(passenger Passenger) (*Pass, error) {
	pass := passenger.GetPass()

	if pass == nil {
		return nil, fmt.Errorf("No electonic pass")
	}

	if d.PassRegistry.GetTicketId(pass) != passenger.GetTicket().ID {
		return nil, fmt.Errorf("Fake electonic pass")
	}

	return pass, nil
}

type GosuslugiApi struct {
	PassRegistry PassRegistry
}

func (gosUslugi *GosuslugiApi) RegisterPassenger(passenger Passenger) (*Pass, error) {
	if passenger.GetTicket().Pets != nil || len(passenger.GetTicket().Pets) > 0 {
		return nil, fmt.Errorf("animal checkin not allowed")
	}

	pass, err := checkin(passenger)

	if pass != nil {
		gosUslugi.PassRegistry.SetPass(pass, passenger.GetTicket().ID)
		passenger.SetPass(pass)
	}

	return pass, err
}
