package structs

import (
	"github.com/UniversityRadioYork/myradio-go"
)

type TimeslotTemplateData struct {
	Timeslot myradio.Timeslot
	Config   Config
}
