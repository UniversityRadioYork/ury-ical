package controllers

import (
	"fmt"
	"github.com/UniversityRadioYork/myradio-go"
	"github.com/UniversityRadioYork/ury-ical/models"
	"github.com/UniversityRadioYork/ury-ical/structs"
	"github.com/UniversityRadioYork/ury-ical/utils/ical"
	"github.com/jaytaylor/html2text"
	"net/http"
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

	for _, timeslot := range timeslots {

		desc, err := html2text.FromString(timeslot.Description)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		cal.AddComponent(ical.VComponent(ical.VEvent{
			UID:     fmt.Sprintf("%d", timeslot.TimeslotID),
			SUMMARY: timeslot.Title,
			DESCRIPTION: fmt.Sprintf(
				`Season %d, Episode %d

%s

Credits: %s

%s%s`,
				timeslot.SeasonNum,
				timeslot.TimeslotNum,
				desc,
				timeslot.CreditsString,
				ic.config.Url,
				timeslot.MicroSiteLink.URL),
			DTSTART:  timeslot.StartTime,
			DTEND:    timeslot.StartTime.Add(timeslot.Duration),
			DTSTAMP:  timeslot.Submitted,
			LOCATION: "University Radio York",
			TZID:     "Europe/London",
			AllDay:   false,
		}))

	}

	w.Header().Set("Content-Type", "text/calendar; charset=utf-8")
	w.Header().Set("Content-Disposition", "inline; filename=ury.ics")
	cal.Encode(w)

}
