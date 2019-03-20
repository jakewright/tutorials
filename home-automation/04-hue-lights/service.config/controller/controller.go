package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jakewright/tutorials/home-automation/04-hue-lights/service.config/domain"
)

// Controller exports the handlers for the endpoints
type Controller struct {
	Config *domain.Config
}

// ReadConfig writes the config for the given service to the ResponseWriter
func (c *Controller) ReadConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	serviceName, ok := vars["serviceName"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "serviceName not set")
		return
	}

	config, err := c.Config.Get(serviceName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	rsp, err := json.Marshal(&config)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(rsp))
}
