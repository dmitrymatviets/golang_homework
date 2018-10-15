package model

import (
	"strings"
)

type ICountryCustomsPolicy interface {
	CheckPassenger(p *Passenger) (bool, string)
}

type RegularCountryCustomsPolicy struct {
}

func commonCheckPassenger(p *Passenger) (bool, string) {
	internationalPassport, _, ticket := p.ShowDocuments()
	if internationalPassport == nil {
		return false, "Отсутствует загран-паспорт"
	}

	if len(internationalPassport.Number) != 9 || internationalPassport.Name == "" {
		return false, "Недействительный загран-паспорт"
	}

	if ticket == nil {
		return false, "Отсутствует билет"
	}

	if len(ticket.Number) < 10 || len(strings.Split(ticket.Number, "-")) != 2 || ticket.DestinationCountry == nil {
		return false, "Недействительный билет"
	}

	if ticket.PassengerName != internationalPassport.Name {
		return false, "Не совпадают имена на билете и паспорте"
	}

	if p.HasPet() {
		pet, petPassport, petOwnershipDocument, petSafetyDocument := p.ShowPet()

		if petOwnershipDocument == nil {
			return false, "Отсутствует документ о праве собственности над животным"
		}
		if petPassport == nil {
			return false, "Отсутствует паспорт животного"
		}
		if len(petOwnershipDocument.Number) < 5 {
			return false, "Недействительный документ о праве собственности над животным"
		}

		if len(petPassport.Number) < 5 {
			return false, "Недействительный паспорт животного"
		}

		if pet.ChipId != petPassport.ChipId {
			return false, "Не совпадает чип животного с паспортом"
		}

		if pet.ChipId != petOwnershipDocument.ChipId {
			return false, "Не совпадает чип животного с документом о праве собственности над животным"
		}

		if pet.WeightKg >= 40 {
			if petSafetyDocument == nil {
				return false, "Нет документа о безопасности животного с весом > 40кг"
			}

			if len(petSafetyDocument.Number) < 5 {
				return false, "Недействительный документ о безопасности животного"
			}

			if pet.ChipId != petSafetyDocument.ChipId {
				return false, "Не совпадает чип животного с документом о безопасности животного"
			}
		}
	}

	return true, ""
}

func (policy RegularCountryCustomsPolicy) CheckPassenger(p *Passenger) (bool, string) {
	return commonCheckPassenger(p)
}

type OnlyMarriedWomenCustomsPolicy struct {
}

func (policy OnlyMarriedWomenCustomsPolicy) CheckPassenger(p *Passenger) (bool, string) {
	ok, msg := commonCheckPassenger(p)
	if !ok {
		return ok, msg
	}
	_, internalPassport, _ := p.ShowDocuments()
	if internalPassport == nil {
		return false, "Остутствует внутренний паспорт"
	}
	if len(internalPassport.Number) != 7 {
		return false, "Недействительный внутренний паспорт"
	}
	if p.IsFemale && !internalPassport.IsMarried {
		return false, "Запрещен въезд незамужним женщинам"
	}
	return true, ""
}

type Country struct {
	Name   string
	Code   string
	Policy ICountryCustomsPolicy `json:"-"`
}
