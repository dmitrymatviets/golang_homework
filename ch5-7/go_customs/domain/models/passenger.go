package domain

type Passenger struct {
	IsFemale bool

	internationalPassport *InternationalPassport
	internalPassport      *InternalPassport
	ticket                *Ticket

	pet                  *Pet
	petPassport          *PetPassport
	petOwnershipDocument *PetOwnershipDocument
	petSafetyDocument    *PetSafetyDocument
}

func (p *Passenger) GetDestinationCountry() *Country {
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

func (p *Passenger) PickUpDocuments(internationalPassport *InternationalPassport, internalPassport *InternalPassport, ticket *Ticket) *Passenger {
	p.internationalPassport = internationalPassport
	p.internalPassport = internalPassport
	p.ticket = ticket
	return p
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

type InternationalPassport struct {
	Number string
	Name   string
}

type InternalPassport struct {
	Number    string
	IsMarried bool
}

type Ticket struct {
	Number             string
	PassengerName      string
	DestinationCountry *Country
}

type Pet struct {
	weightKg int
	chipId   int
}

type PetPassport struct {
	number int
	chipId int
}

type PetOwnershipDocument struct {
	number int
	chipId int
}

type PetSafetyDocument struct {
	number int
	chipId int
}
