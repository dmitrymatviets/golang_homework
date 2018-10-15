package model

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
	Number string
	ChipId string
}

type PetOwnershipDocument struct {
	Number string
	ChipId string
}

type PetSafetyDocument struct {
	Number string
	ChipId string
}
