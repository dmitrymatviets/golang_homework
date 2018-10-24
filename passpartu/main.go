package main

import (
	"fmt"
	"golang_homework/passpartu/model"
	"log"
)

const (
	// "babka"
	GatewayA = "A"
	// "detka"
	GatewayB = "B"
)

func main() {

	b1 := &model.Budka{}
	b2 := &model.Budka{}
	b1.Register(&model.Babka{})
	b2.Register(&model.Detka{})

	eais := model.NewEAISPC()
	eais.RegisterBudka(b1, GatewayA)
	eais.RegisterBudka(b2, GatewayB)

	//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@//

	ticket1 := new(model.Ticket)
	ticket1.Destination = &model.Location{
		Country: "UAE",
	}

	lonelyMan := new(model.American)
	lonelyMan.Passport = &model.AmericanPassport{
		FirstName: "Jonh",
		LastName:  "Doe",
		IsMarried: false,
		Sex:       model.SexMale,
		SSN:       "9822-090233/09-Montana",
	}
	lonelyMan.Ticket = ticket1

	ticket2 := new(model.Ticket)
	ticket2.Destination = &model.Location{
		Country: "USA",
	}
	womanWithDogs := new(model.Russian)
	womanWithDogs.BasePassport = &model.RussianBasePassport{
		FirstName:  "Valeria",
		MiddleName: "Anreevna",
		LastName:   "Alekseeva",
		Sex:        model.SexFemale,
		IsMarried:  false,
		Number:     779802,
		Serie:      "77 01",
	}
	womanWithDogs.InternationalPassport = &model.RussianInternationalPassport{
		FirstName:  "Valeria",
		MiddleName: "Anreevna",
		LastName:   "Alekseeva",
		Sex:        model.SexFemale,
		Serie:      "0901",
		Number:     9761112123434,
	}

	dog := new(model.Animal)
	dog.OwnershipCertificate = &model.OwnershipCertificate{
		Passenger: womanWithDogs,
	}
	dog.SafetyCertificate = &model.SafetyCertificate{}
	dog.Passport = &model.AnimalPassport{}
	dog.Weight = 39

	ticket2.Pets = make([]*model.Animal, 0)
	ticket2.Pets = append(ticket1.Pets, dog)
	womanWithDogs.Ticket = ticket2

	gw1 := eais.Budkas[GatewayA]
	gw2 := eais.Budkas[GatewayB]

	pass, err := gw1.CheckIn(lonelyMan)
	if err != nil {
		log.Println(fmt.Sprintf("Checkin error: %s", err.Error()))

	} else {
		log.Println(fmt.Sprintf("Welcome to board with pass: %s", pass.ID))
	}

	pass, err = gw2.CheckIn(lonelyMan)
	if err != nil {
		log.Println(fmt.Sprintf("Checkin error: %s", err.Error()))

	} else {
		log.Println(fmt.Sprintf("Welcome to board with pass: %s", pass.ID))
	}

	pass, err = gw1.CheckIn(womanWithDogs)
	if err != nil {
		log.Println(fmt.Sprintf("Checkin error: %s", err.Error()))

	} else {
		log.Println(fmt.Sprintf("Welcome to board with pass: %s", pass.ID))
	}

	pass, err = gw2.CheckIn(womanWithDogs)
	if err != nil {
		log.Println(fmt.Sprintf("Checkin error: %s", err.Error()))
	} else {
		log.Println(fmt.Sprintf("Welcome to board with pass: %s", pass.ID))
	}

	womanWithDogs.Ticket.Destination.Country = "RUS"
	pass, err = gw2.CheckIn(womanWithDogs)
	if err != nil {
		log.Println(fmt.Sprintf("Checkin error: %s", err.Error()))
	} else {
		log.Println(fmt.Sprintf("Welcome to board with pass: %s", pass.ID))
	}

}
