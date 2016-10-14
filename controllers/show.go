package controllers

import (
	"bytes"
	"fmt"
	"github.com/UniversityRadioYork/myradio-go"
	"github.com/UniversityRadioYork/ury-ical/models"
	"github.com/UniversityRadioYork/ury-ical/structs"
	"github.com/gorilla/mux"
	"github.com/jaytaylor/html2text"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

// ShowController is the controller for the index page.
type ShowController struct {
	Controller
}

// NewShowController returns a new ShowController with the MyRadio session s
// and configuration context c.
func NewShowController(s *myradio.Session, c *structs.Config) *ShowController {
	return &ShowController{Controller{session: s, config: c}}
}

// Get handles the HTTP GET request r for the index page, writing to w.
func (ic *ShowController) Get(w http.ResponseWriter, r *http.Request) {

	im := models.NewShowModel(ic.session)

	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	show, timeslots, err := im.Get(id)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	cal := ic.config.Calendar

	cal.NAME = fmt.Sprintf("%s - %s", cal.NAME, show.Title)
	cal.X_WR_CALNAME = cal.NAME

	t := template.New("calendar template")
	t.Funcs(template.FuncMap{
		"html2text": html2text.FromString,
		"trim":      strings.TrimSpace,
	})
	t, _ = t.Parse(ic.config.CalendarDescription)

	var desc bytes.Buffer

	data := structs.CalendarTemplateData{
		Show:   show,
		Config: *ic.config,
	}

	err = t.Execute(&desc, data)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	cal.DESCRIPTION = desc.String()
	cal.X_WR_CALDESC = cal.DESCRIPTION

	ic.renderICAL(cal, timeslots, w)

}
