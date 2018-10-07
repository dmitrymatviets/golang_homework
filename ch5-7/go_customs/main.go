package main

import (
	"fmt"
	"go_customs/domain/models"
)

func main() {
	/*
		- мужчина путешествует в ОАЭ один обратился в будку с бабкой
		- тот же мужчина обратился в будку с деткой
		- тот же мужчина обратился в будку электронного контроля без соответствующего талона о прохождении
		- тот же мужчина обратился в будку электронного контроля с соответствующим талона о прохождении
		- те же сценарии только путешествует не замужняя дама с собачкой в США
	*/
	usa := domain.MakeCountry("США", domain.OnlyMarriedWomenCustomsPolicy{})
	uk := domain.MakeCountry("Великобритания", domain.OnlyMarriedWomenCustomsPolicy{})
	russia := domain.MakeCountry("Россия", domain.RegularCountryCustomsPolicy{})

	annaSingle := domain.Passenger{IsFemale: true}
	annaSingle.PickUpDocuments(
		&domain.InternationalPassport{"123456780", "Anna Ivanova"},
		&domain.InternalPassport{Number: "1234560", IsMarried: false},
		&domain.Ticket{"111111-22220", "Anna Ivanova", usa})

	tryToPassCustoms(&annaSingle)

	annaSingle.ChangeTicket(&domain.Ticket{"111111-22220", "Anna Ivanova", russia})
	tryToPassCustoms(&annaSingle)

	olgaMarried := domain.Passenger{IsFemale: true}
	olgaMarried.PickUpDocuments(
		&domain.InternationalPassport{"123456781", "Olga Magomedova"},
		&domain.InternalPassport{Number: "1234561", IsMarried: true},
		&domain.Ticket{"111111-22221", "Olga Magomedova", usa})

	tryToPassCustoms(&olgaMarried)

	alexanderPetrovFakeTicket := domain.Passenger{IsFemale: false}
	alexanderPetrovFakeTicket.PickUpDocuments(
		&domain.InternationalPassport{"123456782", "Alexander Petrov"},
		&domain.InternalPassport{Number: "1234562", IsMarried: false},
		&domain.Ticket{"111111-2", "Alexander Petrov", uk})

	tryToPassCustoms(&alexanderPetrovFakeTicket)

	ruslanBoshirovSpy := domain.Passenger{IsFemale: false}
	ruslanBoshirovSpy.PickUpDocuments(
		&domain.InternationalPassport{"123456783", "Ruslan Boshirov"},
		&domain.InternalPassport{Number: "1234563", IsMarried: false},
		&domain.Ticket{"111111-22222", "Anatoly Chepiga", uk})

	tryToPassCustoms(&ruslanBoshirovSpy)

	ruslanBoshirovSpy.ChangeTicket(&domain.Ticket{"111111-22222", "Ruslan Boshirov", uk})
	tryToPassCustoms(&ruslanBoshirovSpy)
}

func tryToPassCustoms(passenger *domain.Passenger) {
	if passenger == nil {
		panic(fmt.Errorf("Не передан пассажир"))
	}
	country := passenger.GetDestinationCountry()
	if country == nil {
		fmt.Println("Отсутствует билет или страна в билете")
	}

	ok, msg := country.CheckPassenger(passenger)
	fmt.Printf("%v -> %v [%v] (%v)\r\n", passenger.GetName(), country.Name, ok, msg)
}
