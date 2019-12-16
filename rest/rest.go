package main

import (
	"github.com/gorilla/mux"
	"github.com/vlad/rest/apis"
	"github.com/vlad/rest/services"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	var u services.UnitService
	var e services.EmplService
	apis.ServeResource(router, u, e)
}
