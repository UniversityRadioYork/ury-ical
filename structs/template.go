package structs

import (
	"github.com/UniversityRadioYork/myradio-go"
)

type TimeslotTemplateData struct {
	Timeslot myradio.Timeslot
	Config   Config
}

type CalendarTemplateData struct {
	HasShow bool
	Show    myradio.ShowMeta
	Config  Config
}
