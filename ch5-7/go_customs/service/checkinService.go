package service

import (
	"crypto/md5"
	"encoding/hex"
	"golang_homework/ch5-7/go_customs/model"
)

type CheckinService struct {
	Registry map[string]*model.CheckinServiceRequestModel
}

func (service *CheckinService) GetTicket(data *model.CheckinServiceRequestModel) string {
	if data == nil {
		return ""
	}

	if data.Ticket.DestinationCountry == nil {
		return ""
	}
	country := data.Ticket.DestinationCountry

	if country.Policy == nil {
		return ""
	}

	passengerAvatar := &model.Passenger{IsFemale: data.IsFemale}

	passengerAvatar.PickUpDocuments(
		data.InternationalPassport,
		data.InternalPassport,
		data.Ticket)

	passengerAvatar.PickUpPet(data.Pet, nil, nil, nil)

	ok, _ := country.Policy.CheckPassenger(passengerAvatar)

	if ok {
		sum := md5.Sum([]byte(data.Ticket.Number))
		hash := hex.EncodeToString(sum[:])
		service.Registry[hash] = data
		return hash
	}

	return ""
}

func (service *CheckinService) CheckTicket(model *model.CheckinServiceRequestModel, hash string) bool {
	return model.Equals(service.Registry[hash])
}
