package main

import "net/http"

// API is the object that is responsible for serving the API
type API struct {
	Port string
}

// NewAPI creates a new instance of the API
func NewAPI(port string) {
	return &API{
		Port: port,
	}
}

// Serve starts a webserver with the different handlers
func (api *API) Serve() error {
	http.HandleFunc("/", rootHandler)
	err := http.ListenAndServe(api.Port, nil)
	return err
}
