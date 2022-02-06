package controllers

import (
	"bytes"
	"net/http"
	"strings"
	"text/template"

	"github.com/UniversityRadioYork/myradio-go"
	"github.com/UniversityRadioYork/ury-ical/models"
	"github.com/UniversityRadioYork/ury-ical/structs"
	"github.com/jaytaylor/html2text"
)

type MembersController struct {
	Controller
}

func NewMembersController(s *myradio.Session, c *structs.Config) *MembersController {
	return &MembersController{Controller{session: s, config: c}}
}

func (mc *MembersController) Get(w http.ResponseWriter, r *http.Request) {
	mm := models.NewMembersModel(mc.session, mc.config)

	events, err := mm.Get()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	cal := mc.config.Calendar

	t := template.New("calendar template")
	t.Funcs(template.FuncMap{
		"html2text": html2text.FromString,
		"trim":      strings.TrimSpace,
	})
	t, _ = t.Parse(mc.config.CalendarDescription)

	var desc bytes.Buffer

	data := structs.CalendarTemplateData{
		Config: *mc.config,
	}

	err = t.Execute(&desc, data)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	cal.DESCRIPTION = desc.String()
	cal.X_WR_CALDESC = desc.String()

	mc.renderICAL(cal, events, w)

}
