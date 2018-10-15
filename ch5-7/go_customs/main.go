package main

import (
	"fmt"
	"golang_homework/ch5-7/go_customs/api"
	"golang_homework/ch5-7/go_customs/model"
	"golang_homework/ch5-7/go_customs/service"
	"net/http"
)

const (
	usa    = "usa"
	uk     = "uk"
	russia = "russia"
)

var (
	countriesRegistry map[string]*model.Country

	babkaBudka *model.BabkaBudka
	detkaBudka *model.DetkaBudka
	autoBudka  *model.AutoBudka

	srv              *http.Server
	checkinService   model.ICheckinService
	gosUslugiService *service.GosuslugiService
)

func init() {
	countriesRegistry = map[string]*model.Country{
		usa:    {Name: "США", Code: usa, Policy: model.OnlyMarriedWomenCustomsPolicy{}},
		uk:     {Name: "Великобритания", Code: uk, Policy: model.OnlyMarriedWomenCustomsPolicy{}},
		russia: {Name: "Россия", Code: russia, Policy: model.RegularCountryCustomsPolicy{}}}

	babkaBudka = &model.BabkaBudka{}
	detkaBudka = &model.DetkaBudka{}
	checkinService = &service.CheckinService{Registry: make(map[string]*model.CheckinServiceRequestModel)}
	autoBudka = &model.AutoBudka{Service: checkinService}
	srv = api.StartHttpServer(checkinService, API_PORT, countriesRegistry)

	gosUslugiService = &service.GosuslugiService{Server: srv, Port: API_PORT}
}

func main() {
	defer srv.Shutdown(nil)

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

func test_singleWomanToUsa_fail() {

	annaSingle := model.Passenger{IsFemale: true}
	annaSingle.PickUpDocuments(
		&model.InternationalPassport{Number: "123456780", Name: "Anna Nezamuzhnyaya"},
		&model.InternalPassport{Number: "1234560", IsMarried: false},
		&model.Ticket{Number: "111111-22220", PassengerName: "Anna Nezamuzhnyaya", DestinationCountry: countriesRegistry[usa]})

	tryToPassCustoms(&annaSingle, babkaBudka, false)
}

func test_singleWomanToRussia_success() {

	annaSingle := model.Passenger{IsFemale: true}
	annaSingle.PickUpDocuments(
		&model.InternationalPassport{Number: "123456780", Name: "Anna Nezamuzhnyaya"},
		&model.InternalPassport{Number: "1234560", IsMarried: false},
		&model.Ticket{Number: "111111-22220", PassengerName: "Anna Nezamuzhnyaya", DestinationCountry: countriesRegistry[russia]})

	tryToPassCustoms(&annaSingle, babkaBudka, true)
}

func test_marriedWomanToUsa_success() {

	olgaMarried := model.Passenger{IsFemale: true}
	olgaMarried.PickUpDocuments(
		&model.InternationalPassport{Number: "123456781", Name: "Olga Zamuzhnyaya"},
		&model.InternalPassport{Number: "1234561", IsMarried: true},
		&model.Ticket{Number: "111111-22221", PassengerName: "Olga Zamuzhnyaya", DestinationCountry: countriesRegistry[usa]})

	tryToPassCustoms(&olgaMarried, babkaBudka, true)
}

func test_fakeTicket_fail() {

	alexanderPetrovFakeTicket := model.Passenger{IsFemale: false}
	alexanderPetrovFakeTicket.PickUpDocuments(
		&model.InternationalPassport{Number: "123456782", Name: "Alexander Shuler"},
		&model.InternalPassport{Number: "1234562", IsMarried: false},
		&model.Ticket{Number: "111111-2", PassengerName: "Alexander Shuler", DestinationCountry: countriesRegistry[uk]})

	tryToPassCustoms(&alexanderPetrovFakeTicket, babkaBudka, false)
}

func test_ticketWithWrongName_fail() {

	ruslanBoshirovSpy := model.Passenger{IsFemale: false}
	ruslanBoshirovSpy.PickUpDocuments(
		&model.InternationalPassport{Number: "123456783", Name: "Ruslan Boshirov"},
		&model.InternalPassport{Number: "1234563", IsMarried: false},
		&model.Ticket{Number: "111111-22222", PassengerName: "Anatoly Chepiga", DestinationCountry: countriesRegistry[uk]})

	tryToPassCustoms(&ruslanBoshirovSpy, babkaBudka, false)
}

func test_babkaBudkaWorks_success() {

	annaSingle := model.Passenger{IsFemale: true}
	annaSingle.PickUpDocuments(
		&model.InternationalPassport{Number: "123456780", Name: "Anna Normalnaya"},
		&model.InternalPassport{Number: "1234560", IsMarried: true},
		&model.Ticket{Number: "111111-22220", PassengerName: "Anna Normalnaya", DestinationCountry: countriesRegistry[russia]})

	tryToPassCustoms(&annaSingle, babkaBudka, true)
}

func test_detkaBudkaWorks_success() {

	annaSingle := model.Passenger{IsFemale: true}
	annaSingle.PickUpDocuments(
		&model.InternationalPassport{Number: "123456780", Name: "Anna Normalnaya"},
		&model.InternalPassport{Number: "1234560", IsMarried: true},
		&model.Ticket{Number: "111111-22220", PassengerName: "Anna Normalnaya", DestinationCountry: countriesRegistry[russia]})

	tryToPassCustoms(&annaSingle, detkaBudka, true)
}

func test_babkaBudkaPetWithDoc_fail() {

	annaPet := model.Passenger{IsFemale: true}

	annaPet.PickUpDocuments(
		&model.InternationalPassport{Number: "123456780", Name: "Anna Doglover"},
		&model.InternalPassport{Number: "1234560", IsMarried: true},
		&model.Ticket{Number: "111111-22220", PassengerName: "Anna Doglover", DestinationCountry: countriesRegistry[russia]})

	annaPet.PickUpPet(
		&model.Pet{Kind: "sobaka", WeightKg: 10, ChipId: "1234"},
		&model.PetPassport{Number: "123456", ChipId: "1234"},
		&model.PetOwnershipDocument{Number: "123456", ChipId: "1234"},
		&model.PetSafetyDocument{Number: "123456", ChipId: "1234"})

	tryToPassCustoms(&annaPet, babkaBudka, false)
}

func test_detkaBudkaPetWithDoc_success() {

	annaPet := model.Passenger{IsFemale: true}
	annaPet.PickUpDocuments(
		&model.InternationalPassport{Number: "123456780", Name: "Anna Doglover"},
		&model.InternalPassport{Number: "1234560", IsMarried: true},
		&model.Ticket{Number: "111111-22220", PassengerName: "Anna Doglover", DestinationCountry: countriesRegistry[russia]})

	annaPet.PickUpPet(
		&model.Pet{Kind: "sobaka", WeightKg: 10, ChipId: "1234"},
		&model.PetPassport{Number: "123456", ChipId: "1234"},
		&model.PetOwnershipDocument{Number: "123456", ChipId: "1234"},
		&model.PetSafetyDocument{Number: "123456", ChipId: "1234"})
	tryToPassCustoms(&annaPet, detkaBudka, true)
}

func test_detkaBudkaPetWithoutPassport_fail() {

	annaPet := model.Passenger{IsFemale: true}
	annaPet.PickUpDocuments(
		&model.InternationalPassport{Number: "123456780", Name: "Anna Doglover-NoPassport"},
		&model.InternalPassport{Number: "1234560", IsMarried: true},
		&model.Ticket{Number: "111111-22220", PassengerName: "Anna Doglover-NoPassport", DestinationCountry: countriesRegistry[russia]})

	annaPet.PickUpPet(
		&model.Pet{Kind: "sobaka", WeightKg: 10, ChipId: "1234"},
		nil,
		&model.PetOwnershipDocument{Number: "123456", ChipId: "1234"},
		&model.PetSafetyDocument{Number: "123456", ChipId: "1234"})
	tryToPassCustoms(&annaPet, detkaBudka, false)
}

func test_detkaBudkaPetWithoutOwnershipDoc_fail() {

	annaPet := model.Passenger{IsFemale: true}
	annaPet.PickUpDocuments(
		&model.InternationalPassport{Number: "123456780", Name: "Anna Doglover-NoOwnership"},
		&model.InternalPassport{Number: "1234560", IsMarried: true},
		&model.Ticket{Number: "111111-22220", PassengerName: "Anna Doglover-NoOwnership", DestinationCountry: countriesRegistry[russia]})

	annaPet.PickUpPet(
		&model.Pet{Kind: "sobaka", WeightKg: 10, ChipId: "1234"},
		&model.PetPassport{Number: "123456", ChipId: "1234"},
		nil,
		&model.PetSafetyDocument{Number: "123456", ChipId: "1234"})

	tryToPassCustoms(&annaPet, detkaBudka, false)
}

func test_detkaBudkaHeavyPetNoSafetyDoc_fail() {

	annaPet := model.Passenger{IsFemale: true}
	annaPet.PickUpDocuments(
		&model.InternationalPassport{Number: "123456780", Name: "Anna HeavyDoglover-NotSafe"},
		&model.InternalPassport{Number: "1234560", IsMarried: true},
		&model.Ticket{Number: "111111-22220", PassengerName: "Anna HeavyDoglover-NotSafe", DestinationCountry: countriesRegistry[russia]})

	annaPet.PickUpPet(
		&model.Pet{Kind: "sobaka", WeightKg: 55, ChipId: "1234"},
		&model.PetPassport{Number: "123456", ChipId: "1234"},
		&model.PetOwnershipDocument{Number: "123456", ChipId: "1234"},
		nil)

	tryToPassCustoms(&annaPet, detkaBudka, false)
}

func test_detkaBudkaHeavyPetWithSafetyDoc_success() {

	annaPet := model.Passenger{IsFemale: true}
	annaPet.PickUpDocuments(
		&model.InternationalPassport{Number: "123456780", Name: "Anna HeavyDoglover-Safe"},
		&model.InternalPassport{Number: "1234560", IsMarried: true},
		&model.Ticket{Number: "111111-22220", PassengerName: "Anna HeavyDoglover-Safe", DestinationCountry: countriesRegistry[russia]})

	annaPet.PickUpPet(
		&model.Pet{Kind: "sobaka", WeightKg: 55, ChipId: "1234"},
		&model.PetPassport{Number: "123456", ChipId: "1234"},
		&model.PetOwnershipDocument{Number: "123456", ChipId: "1234"},
		&model.PetSafetyDocument{Number: "123456", ChipId: "1234"})

	tryToPassCustoms(&annaPet, detkaBudka, true)
}

func test_autoBudkaWithEReg_success() {

	annaSingle := model.Passenger{IsFemale: true}
	annaSingle.PickUpDocuments(
		&model.InternationalPassport{Number: "123456780", Name: "Anna Elektronnaya"},
		&model.InternalPassport{Number: "1234560", IsMarried: false},
		&model.Ticket{Number: "111111-22220", PassengerName: "Anna Elektronnaya", DestinationCountry: countriesRegistry[russia]})

	annaSingle.PickERegVoucher(gosUslugiService.GetTicket(annaSingle.AsERegModel()))

	tryToPassCustoms(&annaSingle, autoBudka, true)
}

func test_autoBudkaWithFakeEReg_fail() {

	annaSingle := model.Passenger{IsFemale: true}
	annaSingle.PickUpDocuments(
		&model.InternationalPassport{Number: "123456780", Name: "Anna Elektronnaya-Fake"},
		&model.InternalPassport{Number: "1234560", IsMarried: false},
		&model.Ticket{Number: "111111-22220", PassengerName: "Anna Elektronnaya-Fake", DestinationCountry: countriesRegistry[russia]})

	annaSingle.PickERegVoucher("12345678")

	tryToPassCustoms(&annaSingle, autoBudka, false)
}

func test_autoBudkaWithPetEReg_fail() {

	annaSingle := model.Passenger{IsFemale: true}
	annaSingle.PickUpDocuments(
		&model.InternationalPassport{Number: "123456780", Name: "Anna Elektronnaya-Pet"},
		&model.InternalPassport{Number: "1234560", IsMarried: false},
		&model.Ticket{Number: "111111-22220", PassengerName: "Anna Elektronnaya-Pet", DestinationCountry: countriesRegistry[russia]})

	annaSingle.PickUpPet(
		&model.Pet{Kind: "sobaka", WeightKg: 55, ChipId: "1234"},
		&model.PetPassport{Number: "123456", ChipId: "1234"},
		&model.PetOwnershipDocument{Number: "123456", ChipId: "1234"},
		&model.PetSafetyDocument{Number: "123456", ChipId: "1234"})

	annaSingle.PickERegVoucher(gosUslugiService.GetTicket(annaSingle.AsERegModel()))

	tryToPassCustoms(&annaSingle, autoBudka, false)
}

func tryToPassCustoms(passenger *model.Passenger, budka model.IBudka, expectedResult bool) {
	if passenger == nil {
		panic(fmt.Errorf("Не передан пассажир"))
	}

	ok, msg := budka.CheckPassenger(passenger)
	resultSymbol := "☓"
	if ok == expectedResult {
		resultSymbol = "✓"
	}

	fmt.Printf("[%v] %v: %v -> %v [%v] (%v)\r\n", resultSymbol, fmt.Sprintf("%T", budka), passenger.GetName(), passenger.GetDestinationCountry().Name, ok, msg)
}
