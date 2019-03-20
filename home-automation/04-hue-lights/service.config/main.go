package main

import (
	"log"
	"time"

	"github.com/jakewright/muxinator"

	"github.com/jakewright/tutorials/home-automation/04-hue-lights/service.config/controller"
	"github.com/jakewright/tutorials/home-automation/04-hue-lights/service.config/domain"
	"github.com/jakewright/tutorials/home-automation/04-hue-lights/service.config/service"
)

func main() {
	config := domain.Config{}

	configService := service.ConfigService{
		Config:   &config,
		Location: "/data/config.yaml",
	}

	go configService.Watch(time.Second * 30)

	c := controller.Controller{
		Config: &config,
	}

	router := muxinator.NewRouter()
	router.Get("/read/{serviceName}", c.ReadConfig)
	log.Fatal(router.ListenAndServe(":80"))
}
