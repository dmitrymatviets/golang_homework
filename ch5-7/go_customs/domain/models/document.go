package domain

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
