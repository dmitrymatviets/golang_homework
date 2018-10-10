package main

import (
	"fmt"
	"golang_homework/ch5-7/go_customs/domain"
)

func main() {

	test_babkaBudkaWorks_success()
	test_detkaBudkaWorks_success()

	test_singleWomanToUsa_fail()
	test_singleWomanToRussia_success()
	test_marriedWomanToUsa_success()

	test_fakeTicket_fail()
	test_ticketWithWrongName_fail()

	test_babkaBudkaPetWithDoc_fail()
	test_detkaBudkaPetWithDoc_success()
	test_detkaBudkaPetWithoutPassport_fail()
	test_detkaBudkaPetWithoutOwnershipDoc_fail()
	test_detkaBudkaHeavyPetNoSafetyDoc_fail()
	test_detkaBudkaHeavyPetWithSafetyDoc_success()

	test_autoBudkaWithEReg_success()
	test_autoBudkaWithFakeEReg_fail()
	test_autoBudkaWithPetEReg_fail()

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

func test_babkaBudkaPetWithDoc_fail() {
	_, _, russia, babkaBudka, _ := prepare()

	annaPet := domain.Passenger{IsFemale: true}

	annaPet.PickUpDocuments(
		&domain.InternationalPassport{"123456780", "Anna Doglover"},
		&domain.InternalPassport{Number: "1234560", IsMarried: true},
		&domain.Ticket{"111111-22220", "Anna Doglover", russia})

	annaPet.PickUpPet(
		&domain.Pet{"sobaka", 10, "1234"},
		&domain.PetPassport{"123456", "1234"},
		&domain.PetOwnershipDocument{"123456", "1234"},
		&domain.PetSafetyDocument{"123456", "1234"})

	tryToPassCustoms(&annaPet, babkaBudka, false)
}

func test_detkaBudkaPetWithDoc_success() {
	_, _, russia, _, detkaBudka := prepare()

	annaPet := domain.Passenger{IsFemale: true}
	annaPet.PickUpDocuments(
		&domain.InternationalPassport{"123456780", "Anna Doglover"},
		&domain.InternalPassport{Number: "1234560", IsMarried: true},
		&domain.Ticket{"111111-22220", "Anna Doglover", russia})

	annaPet.PickUpPet(
		&domain.Pet{"sobaka", 10, "1234"},
		&domain.PetPassport{"123456", "1234"},
		&domain.PetOwnershipDocument{"123456", "1234"},
		&domain.PetSafetyDocument{"123456", "1234"})
	tryToPassCustoms(&annaPet, detkaBudka, true)
}

func test_detkaBudkaPetWithoutPassport_fail() {
	_, _, russia, _, detkaBudka := prepare()

	annaPet := domain.Passenger{IsFemale: true}
	annaPet.PickUpDocuments(
		&domain.InternationalPassport{"123456780", "Anna Doglover-NoPassport"},
		&domain.InternalPassport{Number: "1234560", IsMarried: true},
		&domain.Ticket{"111111-22220", "Anna Doglover-NoPassport", russia})

	annaPet.PickUpPet(
		&domain.Pet{"sobaka", 10, "1234"},
		nil,
		&domain.PetOwnershipDocument{"123456", "1234"},
		&domain.PetSafetyDocument{"123456", "1234"})
	tryToPassCustoms(&annaPet, detkaBudka, false)
}

func test_detkaBudkaPetWithoutOwnershipDoc_fail() {
	_, _, russia, _, detkaBudka := prepare()

	annaPet := domain.Passenger{IsFemale: true}
	annaPet.PickUpDocuments(
		&domain.InternationalPassport{"123456780", "Anna Doglover-NoOwnership"},
		&domain.InternalPassport{Number: "1234560", IsMarried: true},
		&domain.Ticket{"111111-22220", "Anna Doglover-NoOwnership", russia})

	annaPet.PickUpPet(
		&domain.Pet{"sobaka", 10, "1234"},
		&domain.PetPassport{"123456", "1234"},
		nil,
		&domain.PetSafetyDocument{"123456", "1234"})

	tryToPassCustoms(&annaPet, detkaBudka, false)
}

func test_detkaBudkaHeavyPetNoSafetyDoc_fail() {
	_, _, russia, _, detkaBudka := prepare()

	annaPet := domain.Passenger{IsFemale: true}
	annaPet.PickUpDocuments(
		&domain.InternationalPassport{"123456780", "Anna HeavyDoglover-NotSafe"},
		&domain.InternalPassport{Number: "1234560", IsMarried: true},
		&domain.Ticket{"111111-22220", "Anna HeavyDoglover-NotSafe", russia})

	annaPet.PickUpPet(
		&domain.Pet{"sobaka", 55, "1234"},
		&domain.PetPassport{"123456", "1234"},
		&domain.PetOwnershipDocument{"123456", "1234"},
		nil)

	tryToPassCustoms(&annaPet, detkaBudka, false)
}

func test_detkaBudkaHeavyPetWithSafetyDoc_success() {
	_, _, russia, _, detkaBudka := prepare()

	annaPet := domain.Passenger{IsFemale: true}
	annaPet.PickUpDocuments(
		&domain.InternationalPassport{"123456780", "Anna HeavyDoglover-Safe"},
		&domain.InternalPassport{Number: "1234560", IsMarried: true},
		&domain.Ticket{"111111-22220", "Anna HeavyDoglover-Safe", russia})

	annaPet.PickUpPet(
		&domain.Pet{"sobaka", 55, "1234"},
		&domain.PetPassport{"123456", "1234"},
		&domain.PetOwnershipDocument{"123456", "1234"},
		&domain.PetSafetyDocument{"123456", "1234"})

	tryToPassCustoms(&annaPet, detkaBudka, true)
}

func test_autoBudkaWithEReg_success() {
	_, _, russia, _, _ := prepare()

	eCustomsService := &domain.ECustomsService{make(map[string]*domain.ECustomsServiceRequestModel)}
	autoBudka := &domain.AutoBudka{eCustomsService}

	annaSingle := domain.Passenger{IsFemale: true}
	annaSingle.PickUpDocuments(
		&domain.InternationalPassport{"123456780", "Anna Elektronnaya"},
		&domain.InternalPassport{Number: "1234560", IsMarried: false},
		&domain.Ticket{"111111-22220", "Anna Elektronnaya", russia})

	annaSingle.PickERegVoucher(eCustomsService.GetTicket(annaSingle.AsERegModel()))

	tryToPassCustoms(&annaSingle, autoBudka, true)
}

func test_autoBudkaWithFakeEReg_fail() {
	_, _, russia, _, _ := prepare()

	eCustomsService := &domain.ECustomsService{make(map[string]*domain.ECustomsServiceRequestModel)}
	autoBudka := &domain.AutoBudka{eCustomsService}

	annaSingle := domain.Passenger{IsFemale: true}
	annaSingle.PickUpDocuments(
		&domain.InternationalPassport{"123456780", "Anna Elektronnaya"},
		&domain.InternalPassport{Number: "1234560", IsMarried: false},
		&domain.Ticket{"111111-22220", "Anna Elektronnaya", russia})

	annaSingle.PickERegVoucher("12345678")

	tryToPassCustoms(&annaSingle, autoBudka, false)
}

func test_autoBudkaWithPetEReg_fail() {
	_, _, russia, _, _ := prepare()

	eCustomsService := &domain.ECustomsService{make(map[string]*domain.ECustomsServiceRequestModel)}
	autoBudka := &domain.AutoBudka{eCustomsService}

	annaSingle := domain.Passenger{IsFemale: true}
	annaSingle.PickUpDocuments(
		&domain.InternationalPassport{"123456780", "Anna Elektronnaya-Pet"},
		&domain.InternalPassport{Number: "1234560", IsMarried: false},
		&domain.Ticket{"111111-22220", "Anna Elektronnaya-Pet", russia})

	annaSingle.PickUpPet(
		&domain.Pet{"sobaka", 55, "1234"},
		&domain.PetPassport{"123456", "1234"},
		&domain.PetOwnershipDocument{"123456", "1234"},
		&domain.PetSafetyDocument{"123456", "1234"})

	annaSingle.PickERegVoucher(eCustomsService.GetTicket(annaSingle.AsERegModel()))

	tryToPassCustoms(&annaSingle, autoBudka, false)
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
