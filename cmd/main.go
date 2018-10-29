package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/darkraiden/switch/pkg/handlers"
	"github.com/darkraiden/switch/pkg/parsing"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the switch application on")
	pathsFile := flag.String("filename", "routes.yaml", "a file containing all the redirect routes in a YAML format")
	flag.Parse()

	b, err := ioutil.ReadFile(*pathsFile)
	handlers.CheckErrors(err)

	routes := parsing.ParseYAML(b)
	routesMap := parsing.BuildMap(routes)

	h := handlers.NewHTTPHandler(routesMap)

	http.ListenAndServe(fmt.Sprintf(":%d", *port), h)
}
