package models

import (
	"github.com/UniversityRadioYork/myradio-go"
	"log"
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
func (m *UserModel) Get(id int) (user myradio.User, timeslots []myradio.Timeslot, err error) {
	u, err := m.session.GetUser(id)
	if err != nil {
		return
	}
	user = *u
	timeslots, err = m.getTimeslotsForUser(user)
	return
}

func (m *UserModel) getTimeslotsForUser(user myradio.User) (timeslots []myradio.Timeslot, err error) {
	shows, err := m.session.GetUserShowCredits(user.Memberid)
	if err != nil {
		return
	}
	log.Printf("Found %d shows for user %d", len(shows), user.Memberid)
	for _, show := range shows {
		ts, err := m.getTimeslotsForShow(show)
		if err != nil {
			break
		}
		timeslots = append(timeslots, ts...)
	}
	return
}

func (m *UserModel) getTimeslotsForShow(show myradio.ShowMeta) (timeslots []myradio.Timeslot, err error) {
	seasons, err := m.session.GetSeasons(show.ShowID)
	if err != nil {
		return
	}
	log.Printf("Found %d seasons for show %d", len(seasons), show.ShowID)
	for _, season := range seasons {
		ts, err := m.session.GetTimeslotsForSeason(season.SeasonID)
		if err != nil {
			break
		}
		log.Printf("Found %d timeslots for season %d", len(ts), season.SeasonID)
		timeslots = append(timeslots, ts...)
	}
	return
}
