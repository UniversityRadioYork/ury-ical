package models

import (
	"time"

	"github.com/UniversityRadioYork/myradio-go"
	"github.com/UniversityRadioYork/ury-ical/structs"
)

type MembersModel struct {
	Model
}

func NewMembersModel(s *myradio.Session, c *structs.Config) *MembersModel {
	return &MembersModel{Model{session: s, config: c}}
}

func (m *MembersModel) Get() ([]CalendarEvent, error) {
	myrEvents, err := m.session.GetEventsInRange(time.Now().AddDate(0, -2, 0), time.Now().AddDate(0, 3, 0))

	if err != nil {
		return nil, err
	}

	calEvents, err := myRadioEventsToCalendarEvents(myrEvents, m.config)

	return calEvents, nil
}
