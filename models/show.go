package models

import (
	"github.com/UniversityRadioYork/myradio-go"
	"github.com/UniversityRadioYork/ury-ical/structs"
)

// ShowModel is the model for the Show controller.
type ShowModel struct {
	Model
}

// NewShowModel returns a new ShowModel on the MyRadio session s.
func NewShowModel(s *myradio.Session, c *structs.Config) *ShowModel {
	return &ShowModel{Model{session: s, config: c}}
}

// Get gets the data required for the Show controller from MyRadio.
//
// Otherwise, it returns undefined data and the error causing failure.
func (m *ShowModel) Get(id int) (show myradio.ShowMeta, events []CalendarEvent, err error) {
	s, err := m.session.GetShow(id)
	if err != nil {
		return
	}
	show = *s
	seasons, err := m.session.GetSeasons(id)
	if err != nil {
		return
	}

	var timeslots []myradio.Timeslot
	for _, season := range seasons {
		ts, err := m.session.GetTimeslotsForSeason(season.SeasonID)
		if err != nil {
			break
		}
		timeslots = append(timeslots, ts...)
	}

	events, err = timeslotsToCalendarEvents(timeslots, m.config)

	return
}
