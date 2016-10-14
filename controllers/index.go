package controllers

import (
	"bytes"
	"github.com/UniversityRadioYork/myradio-go"
	"github.com/UniversityRadioYork/ury-ical/models"
	"github.com/UniversityRadioYork/ury-ical/structs"
	"github.com/jaytaylor/html2text"
	"net/http"
	"strings"
	"text/template"
	"log"
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
		log.Print(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	cal := ic.config.Calendar

	t := template.New("calendar template")
	t.Funcs(template.FuncMap{
		"html2text": html2text.FromString,
		"trim":      strings.TrimSpace,
	})
	t, _ = t.Parse(ic.config.CalendarDescription)

	var desc bytes.Buffer

	data := structs.CalendarTemplateData{
		Config: *ic.config,
	}

	err = t.Execute(&desc, data)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	cal.DESCRIPTION = desc.String()
	cal.X_WR_CALDESC = desc.String()

	ic.renderICAL(cal, timeslots, w)

}
