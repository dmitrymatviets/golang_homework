package model

type ICheckinService interface {
	CheckTicket(data *CheckinServiceRequestModel, hash string) bool
	GetTicket(data *CheckinServiceRequestModel) string
}

type CheckinServiceRequestModel struct {
	IsFemale bool

	InternationalPassport *InternationalPassport
	InternalPassport      *InternalPassport
	Ticket                *Ticket
	Pet                   *Pet
}

func (self *CheckinServiceRequestModel) Equals(other *CheckinServiceRequestModel) bool {
	if self == nil {
		return false
	}

	if other == nil {
		return false
	}

	if *self.Ticket != *other.Ticket || *self.InternationalPassport != *other.InternationalPassport {
		return false
	}

	return true
}
