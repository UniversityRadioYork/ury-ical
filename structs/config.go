package structs

import "github.com/UniversityRadioYork/ury-ical/utils/ical"

// Config is a structure containing global website configuration.
//
// See the comments for Server and PageContext for more details.
type Config struct {
	Server              Server         `toml:"server"`
	Calendar            ical.VCalendar `toml:"calendar"`
	ShortName           string         `toml:"shortName"`
	LongName            string         `toml:"longName"`
	Url                 string         `toml:"url"`
	MixcloudUrl         string         `toml:"mixcloudUrl"`
	TimeslotDescription string         `toml:"timeslotDescription"`
	CalendarDescription string         `toml:"calendarDescription"`
}

// Server is a structure containing server configuration.
type Server struct {
	Address string `toml:"address"`
	Port    int    `toml:"port"`
	Timeout int    `toml:"timout"`
}
