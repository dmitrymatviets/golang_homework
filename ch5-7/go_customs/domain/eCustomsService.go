package domain

import (
	"crypto/md5"
	"encoding/hex"
)

type ECustomsServiceRequestModel struct {
	IsFemale bool

	InternationalPassport *InternationalPassport
	InternalPassport      *InternalPassport
	Ticket                *Ticket
	Pet                   *Pet
}

type ECustomsService struct {
	Registry map[string]*ECustomsServiceRequestModel
}

func (service *ECustomsService) GetTicket(model *ECustomsServiceRequestModel) string {
	if model == nil {
		return ""
	}
	if model.Ticket == nil {
		return ""
	}
	if model.Ticket.DestinationCountry == nil {
		return ""
	}
	country := model.Ticket.DestinationCountry

	if country.Policy == nil {
		return ""
	}

	passengerAvatar := &Passenger{IsFemale: model.IsFemale}

	passengerAvatar.PickUpDocuments(
		model.InternationalPassport,
		model.InternalPassport,
		model.Ticket)

	passengerAvatar.PickUpPet(model.Pet, nil, nil, nil)

	ok, _ := country.Policy.CheckPassenger(passengerAvatar)

	if ok {
		sum := md5.Sum([]byte(model.Ticket.Number))
		hash := hex.EncodeToString(sum[:])
		service.Registry[hash] = model
		return hash
	}

	return ""
}

func (service *ECustomsService) CheckTicket(model *ECustomsServiceRequestModel, hash string) bool {
	return service.Registry[hash] != nil && *service.Registry[hash] == *model
}
