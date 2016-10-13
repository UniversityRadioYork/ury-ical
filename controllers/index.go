package controllers

import (
	"bytes"
	"fmt"
	"github.com/UniversityRadioYork/myradio-go"
	"github.com/UniversityRadioYork/ury-ical/models"
	"github.com/UniversityRadioYork/ury-ical/structs"
	"github.com/UniversityRadioYork/ury-ical/utils/ical"
	"github.com/jaytaylor/html2text"
	"net/http"
	"strings"
	"text/template"
)

// IndexController is the controller for the index page.
type IndexController struct {
	Controller
}

// NewIndexController returns a new IndexController with the MyRadio session s
// and configuration context c.
func NewIndexController(s *myradio.Session, c *structs.Config) *IndexController {
	return &IndexController{Controller{session: s, config: c}}
}

// Get handles the HTTP GET request r for the index page, writing to w.
func (ic *IndexController) Get(w http.ResponseWriter, r *http.Request) {

	im := models.NewIndexModel(ic.session)

	timeslots, err := im.Get()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	cal := ic.config.Calendar

	t := template.New("description template")
	t.Funcs(template.FuncMap{
		"html2text": html2text.FromString,
		"trim":      strings.TrimSpace,
	})
	t, _ = t.Parse(ic.config.TimeslotDescription)

	for _, timeslot := range timeslots {

		var desc bytes.Buffer

		data := structs.TimeslotTemplateData{
			Timeslot: timeslot,
			Config:   *ic.config,
		}

		err := t.Execute(&desc, data)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		cal.AddComponent(ical.VComponent(ical.VEvent{
			UID:         fmt.Sprintf("%d", timeslot.TimeslotID),
			SUMMARY:     timeslot.Title,
			DESCRIPTION: desc.String(),
			DTSTART:     timeslot.StartTime,
			DTEND:       timeslot.StartTime.Add(timeslot.Duration),
			DTSTAMP:     timeslot.Submitted,
			LOCATION:    "University Radio York",
			TZID:        "Europe/London",
			AllDay:      false,
		}))

	}

	w.Header().Set("Content-Type", "text/calendar; charset=utf-8")
	w.Header().Set("Content-Disposition", "inline; filename=ury.ics")
	cal.Encode(w)

}
