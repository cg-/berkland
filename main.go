package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/cg-/berkland/cl"
)

var fPort = flag.Int("port", 80, "port for web ui")
var fTemplates = flag.String("templates", "web/templates", "path for web templates")

var clp *cl.CraigsListParser

func init() {
	flag.Parse()
	clp = cl.NewCraigsListParser()
}

func main() {
	fmt.Println("Let's predict some earthquakes.")
	go startWeb(*fPort)

	for {
		time.Sleep(time.Second * 1)
	}
}
