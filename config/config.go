package config

import (
	"encoding/json"
	"sync"
	"os"
	"log"
)

type Site struct {
	Domain string
}

type Configuration struct {
	Site Site
}

var config *Configuration
var once sync.Once

func LoadConfig() *Configuration {
    once.Do(func() {
        file, err := os.Open("config.json")
        if err != nil {
            log.Fatalln("Cannot open config file", err)
        }
        decoder := json.NewDecoder(file)
        config = &Configuration{}
        err = decoder.Decode(config)
        if err != nil {
            log.Fatalln("Cannot get configuration from file", err)
        }
    })
    return config

}