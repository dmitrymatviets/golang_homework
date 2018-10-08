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

	test_babkaBudkaWorks_success()
	test_detkaBudkaWorks_success()
	test_singleWomanToUsa_fail()
	test_singleWomanToRussia_success()
	test_marriedWomanToUsa_success()
	test_fakeTicket_fail()
	test_ticketWithWrongName_fail()
	test_babkaBudkaPet_fail()
	test_detkaBudkaPet_success()
}

func prepare() (*domain.Country, *domain.Country, *domain.Country, *domain.BabkaBudka, *domain.DetkaBudka) {
	usa := domain.MakeCountry("США", domain.OnlyMarriedWomenCustomsPolicy{})
	uk := domain.MakeCountry("Великобритания", domain.OnlyMarriedWomenCustomsPolicy{})
	russia := domain.MakeCountry("Россия", domain.RegularCountryCustomsPolicy{})
	babkaBudka := &domain.BabkaBudka{}
	detkaBudka := &domain.DetkaBudka{}
	return usa, uk, russia, babkaBudka, detkaBudka
}

func test_singleWomanToUsa_fail() {
	usa, _, _, babkaBudka, _ := prepare()

	annaSingle := domain.Passenger{IsFemale: true}
	annaSingle.PickUpDocuments(
		&domain.InternationalPassport{"123456780", "Anna Nezamuzhnyaya"},
		&domain.InternalPassport{Number: "1234560", IsMarried: false},
		&domain.Ticket{"111111-22220", "Anna Nezamuzhnyaya", usa})

	tryToPassCustoms(&annaSingle, babkaBudka, false)
}

func test_singleWomanToRussia_success() {
	_, _, russia, babkaBudka, _ := prepare()

	annaSingle := domain.Passenger{IsFemale: true}
	annaSingle.PickUpDocuments(
		&domain.InternationalPassport{"123456780", "Anna Nezamuzhnyaya"},
		&domain.InternalPassport{Number: "1234560", IsMarried: false},
		&domain.Ticket{"111111-22220", "Anna Nezamuzhnyaya", russia})

	tryToPassCustoms(&annaSingle, babkaBudka, true)
}

func test_marriedWomanToUsa_success() {
	usa, _, _, babkaBudka, _ := prepare()

	olgaMarried := domain.Passenger{IsFemale: true}
	olgaMarried.PickUpDocuments(
		&domain.InternationalPassport{"123456781", "Olga Zamuzhnyaya"},
		&domain.InternalPassport{Number: "1234561", IsMarried: true},
		&domain.Ticket{"111111-22221", "Olga Zamuzhnyaya", usa})

	tryToPassCustoms(&olgaMarried, babkaBudka, true)
}

func test_fakeTicket_fail() {
	_, uk, _, babkaBudka, _ := prepare()

	alexanderPetrovFakeTicket := domain.Passenger{IsFemale: false}
	alexanderPetrovFakeTicket.PickUpDocuments(
		&domain.InternationalPassport{"123456782", "Alexander Shuler"},
		&domain.InternalPassport{Number: "1234562", IsMarried: false},
		&domain.Ticket{"111111-2", "Alexander Shuler", uk})

	tryToPassCustoms(&alexanderPetrovFakeTicket, babkaBudka, false)
}

func test_ticketWithWrongName_fail() {
	_, uk, _, babkaBudka, _ := prepare()

	ruslanBoshirovSpy := domain.Passenger{IsFemale: false}
	ruslanBoshirovSpy.PickUpDocuments(
		&domain.InternationalPassport{"123456783", "Ruslan Boshirov"},
		&domain.InternalPassport{Number: "1234563", IsMarried: false},
		&domain.Ticket{"111111-22222", "Anatoly Chepiga", uk})

	tryToPassCustoms(&ruslanBoshirovSpy, babkaBudka, false)
}

func test_babkaBudkaWorks_success() {
	_, _, russia, babkaBudka, _ := prepare()

	annaSingle := domain.Passenger{IsFemale: true}
	annaSingle.PickUpDocuments(
		&domain.InternationalPassport{"123456780", "Anna Normalnaya"},
		&domain.InternalPassport{Number: "1234560", IsMarried: true},
		&domain.Ticket{"111111-22220", "Anna Normalnaya", russia})

	tryToPassCustoms(&annaSingle, babkaBudka, true)
}

func test_detkaBudkaWorks_success() {
	_, _, russia, _, detkaBudka := prepare()

	annaSingle := domain.Passenger{IsFemale: true}
	annaSingle.PickUpDocuments(
		&domain.InternationalPassport{"123456780", "Anna Normalnaya"},
		&domain.InternalPassport{Number: "1234560", IsMarried: true},
		&domain.Ticket{"111111-22220", "Anna Normalnaya", russia})

	tryToPassCustoms(&annaSingle, detkaBudka, true)
}

func tryToPassCustoms(passenger *domain.Passenger, budka domain.Budka, result bool) {
	if passenger == nil {
		panic(fmt.Errorf("Не передан пассажир"))
	}

	ok, msg := budka.CheckPassenger(passenger)
	resultSymbol := "☓"
	if ok == result {
		resultSymbol = "✓"
	}

	fmt.Printf("[%v] %v: %v -> %v [%v] (%v)\r\n", resultSymbol, fmt.Sprintf("%T", budka), passenger.GetName(), *passenger.GetDestinationCountry(), ok, msg)
}
