package main

import (
	"flag"
	"github.com/Convocation-of-Empyreans-Development/shortlink/config"
	"github.com/Convocation-of-Empyreans-Development/shortlink/endpoints"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

var flagConfig = flag.String("config", "config.json", "path to config file")

func main() {
	redirects := config.LoadConfig(*flagConfig)
	router := httprouter.New()
	router.GET("/:path", endpoints.RouteRequest(redirects.Mapping))
	log.Fatal(http.ListenAndServe(":10001", router))
}
