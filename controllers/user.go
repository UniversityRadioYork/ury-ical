package controllers

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"github.com/UniversityRadioYork/myradio-go"
	"github.com/UniversityRadioYork/ury-ical/models"
	"github.com/UniversityRadioYork/ury-ical/structs"
	"github.com/gorilla/mux"
	"github.com/jaytaylor/html2text"
)

// UserController is the controller for the index page.
type UserController struct {
	Controller
}

// NewUserController returns a new UserController with the MyRadio session s
// and configuration context c.
func NewUserController(s *myradio.Session, c *structs.Config) *UserController {
	return &UserController{Controller{session: s, config: c}}
}

// Get handles the HTTP GET request r for the index page, writing to w.
func (ic *UserController) Get(w http.ResponseWriter, r *http.Request) {

	im := models.NewUserModel(ic.session, ic.config)

	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	user, timeslots, err := im.Get(id)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	log.Printf("Found %d timeslots for user %d", len(timeslots), id)

	cal := ic.config.Calendar

	cal.NAME = fmt.Sprintf("%s - %s %s", cal.NAME, user.Fname, user.Sname)
	cal.X_WR_CALNAME = cal.NAME

	t := template.New("calendar template")
	t.Funcs(template.FuncMap{
		"html2text": html2text.FromString,
		"trim":      strings.TrimSpace,
	})
	t, _ = t.Parse(ic.config.CalendarDescription)

	var desc bytes.Buffer

	data := structs.CalendarTemplateData{
		User:   user,
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
