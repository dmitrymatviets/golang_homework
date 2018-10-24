package model

import (
	"fmt"
	"time"
)

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
