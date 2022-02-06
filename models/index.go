package models

import (
	"github.com/UniversityRadioYork/myradio-go"
	"github.com/UniversityRadioYork/ury-ical/structs"
)

// IndexModel is the model for the Index controller.
type IndexModel struct {
	Model
}

// NewIndexModel returns a new IndexModel on the MyRadio session s.
func NewIndexModel(s *myradio.Session, c *structs.Config) *IndexModel {
	return &IndexModel{Model{session: s, config: c}}
}

// Get gets the data required for the Index controller from MyRadio.
//
// Otherwise, it returns undefined data and the error causing failure.
func (m *IndexModel) Get() ([]CalendarEvent, error) {
	seasons, err := m.session.GetAllSeasonsInLatestTerm()
	if err != nil {
		return nil, err
	}
	var timeslots []myradio.Timeslot
	for _, season := range seasons {
		ts, err := m.session.GetTimeslotsForSeason(season.SeasonID)
		if err != nil {
			return nil, err
		}
		timeslots = append(timeslots, ts...)
	}

	events, err := timeslotsToCalendarEvents(timeslots, m.config)

	if err != nil {
		return nil, err
	}

	return events, nil
}
