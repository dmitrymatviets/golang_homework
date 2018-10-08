package main

import (
	"fmt"
	"golang_homework/ch5-7/go_customs/domain/models"
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

	test_singleWomanToUsa_fail(usa)
	test_singleWomanToRussia_success(russia)
	test_marriedWomanToUsa_success(usa)
	test_fakeTicket_fail(uk)
	test_ticketWithWrongName_fail(uk)
}

func test_singleWomanToUsa_fail(usa *domain.Country) {
	annaSingle := domain.Passenger{IsFemale: true}
	annaSingle.PickUpDocuments(
		&domain.InternationalPassport{"123456780", "Anna Nezamuzhnyaya"},
		&domain.InternalPassport{Number: "1234560", IsMarried: false},
		&domain.Ticket{"111111-22220", "Anna Nezamuzhnyaya", usa})
	tryToPassCustoms(&annaSingle, false)
}

func test_singleWomanToRussia_success(russia *domain.Country) {
	annaSingle := domain.Passenger{IsFemale: true}
	annaSingle.PickUpDocuments(
		&domain.InternationalPassport{"123456780", "Anna Nezamuzhnyaya"},
		&domain.InternalPassport{Number: "1234560", IsMarried: false},
		&domain.Ticket{"111111-22220", "Anna Nezamuzhnyaya", russia})
	tryToPassCustoms(&annaSingle, true)
}

func test_marriedWomanToUsa_success(usa *domain.Country) {
	olgaMarried := domain.Passenger{IsFemale: true}
	olgaMarried.PickUpDocuments(
		&domain.InternationalPassport{"123456781", "Olga Zamuzhnyaya"},
		&domain.InternalPassport{Number: "1234561", IsMarried: true},
		&domain.Ticket{"111111-22221", "Olga Zamuzhnyaya", usa})

	tryToPassCustoms(&olgaMarried, true)
}

func test_fakeTicket_fail(uk *domain.Country) {
	alexanderPetrovFakeTicket := domain.Passenger{IsFemale: false}
	alexanderPetrovFakeTicket.PickUpDocuments(
		&domain.InternationalPassport{"123456782", "Alexander Shuler"},
		&domain.InternalPassport{Number: "1234562", IsMarried: false},
		&domain.Ticket{"111111-2", "Alexander Shuler", uk})

	tryToPassCustoms(&alexanderPetrovFakeTicket, false)
}

func test_ticketWithWrongName_fail(uk *domain.Country) {
	ruslanBoshirovSpy := domain.Passenger{IsFemale: false}
	ruslanBoshirovSpy.PickUpDocuments(
		&domain.InternationalPassport{"123456783", "Ruslan Boshirov"},
		&domain.InternalPassport{Number: "1234563", IsMarried: false},
		&domain.Ticket{"111111-22222", "Anatoly Chepiga", uk})

	tryToPassCustoms(&ruslanBoshirovSpy, false)
}

func tryToPassCustoms(passenger *domain.Passenger, result bool) {
	if passenger == nil {
		panic(fmt.Errorf("Не передан пассажир"))
	}
	country := passenger.GetDestinationCountry()
	if country == nil {
		fmt.Println("Отсутствует билет или страна в билете", passenger.GetName())
	}

	ok, msg := country.CheckPassenger(passenger)
	resultSymbol := "☓"
	if ok == result {
		resultSymbol = "✓"
	}
	fmt.Printf("%v %v -> %v [%v] (%v)\r\n", resultSymbol, passenger.GetName(), country.Name, ok, msg)
}
