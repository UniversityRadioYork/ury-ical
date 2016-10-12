package web

import (
	"github.com/UniversityRadioYork/myradio-go"
	"github.com/UniversityRadioYork/ury-ical/controllers"
	"github.com/UniversityRadioYork/ury-ical/structs"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	*negroni.Negroni
}

func NewServer(c *structs.Config) (*Server, error) {

	s := Server{negroni.Classic()}

	session, err := myradio.NewSessionFromKeyFile()

	if err != nil {
		return &s, err
	}

	router := mux.NewRouter().StrictSlash(true)
	router.NotFoundHandler = http.HandlerFunc(http.NotFound)

	getRouter := router.Methods("GET").Subrouter()

	// Routes go in here

	ic := controllers.NewIndexController(session, c)
	getRouter.HandleFunc("/ury.ics", ic.Get)

	// End routes

	s.UseHandler(router)

	return &s, nil

}
