package model

type IBudka interface {
	CheckPassenger(p *Passenger) (bool, string)
}

type BabkaBudka struct {
}

func baseCheckPassenger(p *Passenger) (bool, string) {
	country := p.GetDestinationCountry()
	if country == nil {
		return false, "Неизвестная страна"
	}
	if country.Policy == nil {
		return false, "Неизвестная политика на въезд страны"
	}
	return country.Policy.CheckPassenger(p)
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

type AutoBudka struct {
	Service ICheckinService
}

func (budka AutoBudka) CheckPassenger(p *Passenger) (bool, string) {

	if budka.Service == nil {
		return false, "Сервис недоступен"
	}

	voucher := p.ShowERegVoucher()

	if voucher == "" {
		return false, "Отсутствует электронный ваучер"
	}

	ok := budka.Service.CheckTicket(p.AsERegModel(), voucher)

	if !ok {
		return false, "Не найдена запись о регистрации для электронного ваучера " + voucher
	}

	return true, voucher
}
