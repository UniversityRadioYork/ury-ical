package controllers

import (
	"net/http"

	"github.com/UniversityRadioYork/myradio-go"
	"github.com/UniversityRadioYork/ury-ical/models"
	"github.com/UniversityRadioYork/ury-ical/structs"
	"github.com/UniversityRadioYork/ury-ical/utils/ical"
)

// ControllerInterface is the interface to which controllers adhere.
type ControllerInterface interface {
	Get()     //method = GET processing
	Post()    //method = POST processing
	Delete()  //method = DELETE processing
	Put()     //method = PUT handling
	Head()    //method = HEAD processing
	Patch()   //method = PATCH treatment
	Options() //method = OPTIONS processing
}

// Controller is the base type of controllers in the ury-ical architecture.
type Controller struct {
	session *myradio.Session
	config  *structs.Config
}

// Get handles a HTTP GET request r, writing to w.
//
// Unless overridden, controllers refuse this method.
func (c *Controller) Get(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method Not Allowed", 405)
}

// Post handles a HTTP POST request r, writing to w.
//
// Unless overridden, controllers refuse this method.
func (c *Controller) Post(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method Not Allowed", 405)
}

// Delete handles a HTTP DELETE request r, writing to w.
//
// Unless overridden, controllers refuse this method.
func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method Not Allowed", 405)
}

// Put handles a HTTP PUT request r, writing to w.
//
// Unless overridden, controllers refuse this method.
func (c *Controller) Put(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method Not Allowed", 405)
}

// Head handles a HTTP HEAD request r, writing to w.
//
// Unless overridden, controllers refuse this method.
func (c *Controller) Head(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method Not Allowed", 405)
}

// Patch handles a HTTP PATCH request r, writing to w.
//
// Unless overridden, controllers refuse this method.
func (c *Controller) Patch(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method Not Allowed", 405)
}

// Options handles a HTTP OPTIONS request r, writing to w.
//
// Unless overridden, controllers refuse this method.
func (c *Controller) Options(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method Not Allowed", 405)
}

// Propfind handles a HTTP PROPFIND request r, writing to w.
//
// Unless overridden, controllers refuse this method.
func (c *Controller) Propfind(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method Not Allowed", 405)
}

func (c *Controller) renderICAL(cal ical.VCalendar, events []models.CalendarEvent, w http.ResponseWriter) {
	for _, event := range events {

		cal.AddComponent(ical.VComponent(ical.VEvent{
			UID:         event.ID,
			SUMMARY:     event.Title,
			DESCRIPTION: event.Description,
			DTSTART:     event.StartTime,
			DTEND:       event.EndTime,
			DTSTAMP:     event.CreatedTime,
			LOCATION:    "University Radio York",
			TZID:        "Europe/London",
			AllDay:      false,
		}))

	}

	w.Header().Set("Content-Type", "text/calendar; charset=utf-8")
	w.Header().Set("Content-Disposition", "inline; filename=ury.ics")
	cal.Encode(w)

}
