package main

import (
	"fmt"
	"github.com/UniversityRadioYork/ury-ical/utils"
	"github.com/UniversityRadioYork/ury-ical/web"
	"log"
)

func main() {

	log.SetFlags(log.Llongfile)

	//Get the config from the config.yaml file
	config, err := utils.GetConfigFromFile("./config.toml")

	s, err := web.NewServer(config)

	if err != nil {
		log.Fatal(err)
	}

	l := fmt.Sprintf("%s:%d", config.Server.Address, config.Server.Port)

	log.Printf("Listening on: %s", l)

	s.Run(l)

}
