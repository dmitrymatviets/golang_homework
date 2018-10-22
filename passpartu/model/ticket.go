package model

type Ticket struct {
	ID          string
	Destination *Location
	Pets        []*Animal
}
