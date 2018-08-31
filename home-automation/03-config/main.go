package main

import (
	"log"
	"time"

	"github.com/jakewright/muxinator"

	"github.com/jakewright/tutorials/home-automation/03-config/controller"
	"github.com/jakewright/tutorials/home-automation/03-config/domain"
	"github.com/jakewright/tutorials/home-automation/03-config/service"
)

func main() {
	config := domain.Config{}

	configService := service.ConfigService{
		Config:   &config,
		Location: "config.yaml",
	}

	go configService.Watch(time.Second * 30)

	c := controller.Controller{
		Config: &config,
	}

	router := muxinator.NewRouter()
	router.Get("/read/{serviceName}", c.ReadConfig)
	log.Fatal(router.ListenAndServe(":8080"))
}
