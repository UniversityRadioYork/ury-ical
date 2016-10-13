package models

import (
	"github.com/UniversityRadioYork/myradio-go"
)

// ShowModel is the model for the Show controller.
type ShowModel struct {
	Model
}

// NewShowModel returns a new ShowModel on the MyRadio session s.
func NewShowModel(s *myradio.Session) *ShowModel {
	// @TODO: Pass in the config options
	return &ShowModel{Model{session: s}}
}

// Get gets the data required for the Show controller from MyRadio.
//
// Otherwise, it returns undefined data and the error causing failure.
func (m *ShowModel) Get(id int) (show myradio.ShowMeta, timeslots []myradio.Timeslot, err error) {
	s, err := m.session.GetShow(id)
	if err != nil {
		return
	}
	show = *s
	seasons, err := m.session.GetSeasons(id)
	if err != nil {
		return
	}
	for _, season := range seasons {
		ts, err := m.session.GetTimeslotsForSeason(season.SeasonID)
		if err != nil {
			break
		}
		timeslots = append(timeslots, ts...)
	}
	return
}
