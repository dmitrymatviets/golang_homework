package domain

import (
	"fmt"
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
	ICountryCustomsPolicy
	Name string
}

func MakeCountry(name string, customsPolicy ICountryCustomsPolicy) *Country {
	if customsPolicy == nil {
		panic(fmt.Errorf("Не передан параметр customsPolicy"))
	}

	if name == "" {
		panic(fmt.Errorf("Не передан параметр name"))
	}

	return &Country{customsPolicy, name}
}
