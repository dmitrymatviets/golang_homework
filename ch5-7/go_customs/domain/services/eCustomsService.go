package domain

import (
	"crypto/md5"
	"encoding/hex"
	"go_customs/domain/models"
)

type ECustomsServiceRequestModel struct {
	IsFemale bool

	InternationalPassport *domain.InternationalPassport
	InternalPassport      *domain.InternalPassport
	Ticket                *domain.Ticket
}

type ECustomsService struct {
	registry map[string]*ECustomsServiceRequestModel
}

func (service *ECustomsService) GetTicket(model *ECustomsServiceRequestModel) string {
	//TODO validate
	sum := md5.Sum([]byte("sads"))
	hash := hex.EncodeToString(sum[:])
	service.registry[hash] = model
	return hash
}

func (service *ECustomsService) CheckTicket(model *ECustomsServiceRequestModel, hash string) bool {
	//TODO validate
	return service.registry[hash] != nil && service.registry[hash] == model
}
