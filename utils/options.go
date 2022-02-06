package utils

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/UniversityRadioYork/ury-ical/structs"
)

// GetConfigFromFile reads the website config from the given path.
//
// path is a filepath, relative to the current working directory, of a
// TOML file marshallable to a structs.Config struct.
//
// Returns a config struct and nil if the config read was successful,
// and an undefined value and non-nil otherwise.
func GetConfigFromFile(path string) (c structs.Config, err error) {
	absPath, _ := filepath.Abs(path)

	b, err := os.ReadFile(absPath)

	if err != nil {
		return
	}

	s := string(b)

	_, err = toml.Decode(s, &c)

	return
}
