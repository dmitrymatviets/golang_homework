package domain

type IPassenger interface {
}

type Passenger struct {
	IsFemale bool

	internationalPassport *InternationalPassport
	internalPassport      *InternalPassport
	ticket                *Ticket

	eRegVoucher string

	pet                  *Pet
	petPassport          *PetPassport
	petOwnershipDocument *PetOwnershipDocument
	petSafetyDocument    *PetSafetyDocument
}

func (p Passenger) GetDestinationCountry() *Country {
	if p.ticket != nil {
		return p.ticket.DestinationCountry
	}
	return nil
}

func (p *Passenger) GetName() string {
	if p.internationalPassport != nil {
		return p.internationalPassport.Name
	}
	return ""
}

func (p *Passenger) HasPet() bool {
	return p.pet != nil
}

func (p *Passenger) PickUpDocuments(internationalPassport *InternationalPassport, internalPassport *InternalPassport, ticket *Ticket) *Passenger {
	p.internationalPassport = internationalPassport
	p.internalPassport = internalPassport
	p.ticket = ticket
	return p
}

func (p *Passenger) PickERegVoucher(eRegVoucher string) *Passenger {
	p.eRegVoucher = eRegVoucher
	return p
}

func (p *Passenger) ShowERegVoucher() string {
	return p.eRegVoucher
}

func (p *Passenger) ChangeTicket(ticket *Ticket) *Passenger {
	p.ticket = ticket
	return p
}

func (p *Passenger) ShowDocuments() (*InternationalPassport, *InternalPassport, *Ticket) {
	return p.internationalPassport, p.internalPassport, p.ticket
}

func (p *Passenger) PickUpPet(pet *Pet, petPassport *PetPassport, petOwnershipDocument *PetOwnershipDocument, petSafetyDocument *PetSafetyDocument) *Passenger {
	p.pet = pet
	p.petPassport = petPassport
	p.petOwnershipDocument = petOwnershipDocument
	p.petSafetyDocument = petSafetyDocument
	return p
}

func (p *Passenger) ShowPet() (*Pet, *PetPassport, *PetOwnershipDocument, *PetSafetyDocument) {
	return p.pet, p.petPassport, p.petOwnershipDocument, p.petSafetyDocument
}

func (p *Passenger) AsERegModel() *ECustomsServiceRequestModel {
	return &ECustomsServiceRequestModel{
		p.IsFemale,
		p.internationalPassport,
		p.internalPassport,
		p.ticket,
		p.pet}
}

type Pet struct {
	Kind     string
	WeightKg int
	ChipId   string
}
