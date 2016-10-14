package structs

import (
	"github.com/UniversityRadioYork/myradio-go"
)

type TimeslotTemplateData struct {
	Timeslot myradio.Timeslot
	Config   Config
}

type CalendarTemplateData struct {
	Show   myradio.ShowMeta
	User   myradio.Member
	Config Config
}
