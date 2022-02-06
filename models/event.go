package models

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
	"time"

	"github.com/UniversityRadioYork/myradio-go"
	"github.com/UniversityRadioYork/ury-ical/structs"
	"github.com/jaytaylor/html2text"
)

type CalendarEvent struct {
	ID          string
	Title       string
	Description string
	StartTime   time.Time
	EndTime     time.Time
	CreatedTime time.Time
}

func timeslotsToCalendarEvents(timeslots []myradio.Timeslot, config *structs.Config) ([]CalendarEvent, error) {
	var calEvents []CalendarEvent

	t := template.New("description template")
	t.Funcs(template.FuncMap{
		"html2text": html2text.FromString,
		"trim":      strings.TrimSpace,
	})
	t, _ = t.Parse(config.TimeslotDescription)

	for _, ts := range timeslots {
		var desc bytes.Buffer

		data := structs.TimeslotTemplateData{
			Timeslot: ts,
			Config:   *config,
		}

		err := t.Execute(&desc, data)

		if err != nil {
			return nil, err
		}

		calEvents = append(calEvents, CalendarEvent{
			ID:          fmt.Sprintf("TS-%d", ts.TimeslotID),
			Title:       ts.Title,
			Description: desc.String(),
			StartTime:   ts.StartTime,
			EndTime:     ts.StartTime.Add(ts.Duration),
			CreatedTime: ts.Submitted,
		})
	}

	return calEvents, nil
}

func myRadioEventsToCalendarEvents(events []myradio.Event, config *structs.Config) ([]CalendarEvent, error) {
	var calEvents []CalendarEvent

	t := template.New("event template")
	t.Funcs(template.FuncMap{
		"html2text": html2text.FromString,
		"trim":      strings.TrimSpace,
	})
	t, _ = t.Parse(config.EventDescription)

	for _, ev := range events {
		var desc bytes.Buffer

		data := structs.EventDescription{
			Description: ev.HTMLDescription,
		}

		err := t.Execute(&desc, data)

		if err != nil {
			return nil, err
		}

		calEvents = append(calEvents, CalendarEvent{
			ID:          fmt.Sprintf("EV-%d", ev.ID),
			Title:       ev.Title,
			Description: desc.String(),
			StartTime:   ev.StartTime,
			EndTime:     ev.EndTime,
			CreatedTime: ev.StartTime.AddDate(0, -1, 0),
		})
	}

	return calEvents, nil
}
