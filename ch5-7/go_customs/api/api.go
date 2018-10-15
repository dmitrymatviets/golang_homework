package api

import (
	"encoding/json"
	"fmt"
	"golang_homework/ch5-7/go_customs/model"
	"io"
	"io/ioutil"
	"net/http"
)

func StartHttpServer(service model.ICheckinService, port string, countryRegistry map[string]*model.Country) *http.Server {
	srv := &http.Server{Addr: ":" + port}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		result := &model.CheckinServiceRequestModel{}

		err = json.Unmarshal(b, result)
		if err != nil {
			panic(err)
		}
		result.Ticket.DestinationCountry = countryRegistry[result.Ticket.DestinationCountry.Code]

		ticket := service.GetTicket(result)
		io.WriteString(w, ticket)
	})

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			panic(fmt.Errorf("Httpserver: ListenAndServe() error: %s", err))
		}
	}()

	return srv
}
