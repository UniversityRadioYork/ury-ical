package models

import (
	"github.com/UniversityRadioYork/myradio-go"
)

// UserModel is the model for the User controller.
type UserModel struct {
	Model
}

// NewUserModel returns a new UserModel on the MyRadio session s.
func NewUserModel(s *myradio.Session) *UserModel {
	return &UserModel{Model{session: s}}
}

// Get gets the data required for the User controller from MyRadio.
//
// Otherwise, it returns undefined data and the error causing failure.
func (m *UserModel) Get(id int) (user myradio.Member, timeslots []myradio.Timeslot, err error) {
	s, err := m.session.GetMember(id)
	if err != nil {
		return
	}
	user = *s
	shows, err := m.session.GetUserShowCredits(id)
	if err != nil {
		return
	}
	for _, show := range shows {
		seasons, err := m.session.GetSeasons(show.ShowID)
		if err != nil {
			break
		}
		for _, season := range seasons {
			ts, err := m.session.GetTimeslotsForSeason(season.SeasonID)
			if err != nil {
				break
			}
			timeslots = append(timeslots, ts...)
		}
	}
	return
}
