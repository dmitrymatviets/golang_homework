package domain

type Budka interface {
	CheckPassenger(p *Passenger) (bool, string)
}

type BabkaBudka struct {
}

func baseCheckPassenger(p *Passenger) (bool, string) {
	country := p.GetDestinationCountry()
	if country == nil {
		return false, "Неизвестная страна"
	}
	return country.CheckPassenger(p)
}

func (budka BabkaBudka) CheckPassenger(p *Passenger) (bool, string) {

	if p.HasPet() {
		return false, "Сходи к детке с животным!"
	}

	return baseCheckPassenger(p)
}

type DetkaBudka struct {
}

func (budka DetkaBudka) CheckPassenger(p *Passenger) (bool, string) {
	return baseCheckPassenger(p)
}
