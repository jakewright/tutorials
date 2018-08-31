package service

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/jakewright/tutorials/home-automation/03-config/domain"
)

type ConfigService struct {
	Config   *domain.Config
	Location string
}

// Watch reloads the config every d duration
func (s *ConfigService) Watch(d time.Duration) {
	for {
		err := s.Reload()
		if err != nil {
			log.Print(err)
		}

		time.Sleep(d)
	}
}

// Reload reads the config and applies changes
func (s *ConfigService) Reload() error {
	data, err := ioutil.ReadFile(s.Location)
	if err != nil {
		return err
	}

	err = s.Config.SetFromBytes(data)
	if err != nil {
		return err
	}

	return nil
}
