package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var temp *template.Template

type tData struct {
	Test string
}

func init() {
	fMap := template.FuncMap{}
	templatePath, err := filepath.Abs(*fTemplates)
	if err != nil {
		log.Fatal(err)
	}
	temp, err = template.New("templates").Funcs(fMap).ParseGlob(filepath.Join(templatePath, "*.template"))
	if err != nil {
		log.Fatal(err)
	}
}

func startWeb(port int) {
	log.Printf("Starting the web ui on port %d\n", port)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	td := &tData{
		Test: "LOL",
	}
	err := temp.ExecuteTemplate(w, "home.template", td)
	if err != nil {
		log.Fatal(err)
	}
}
