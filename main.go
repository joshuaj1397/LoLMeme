package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/joshuaj1397/LoLMemes/controller"
)

var tpl *template.Template

type formData struct {
	SummonerName string
	ChampionName string
	Region       string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func index(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func memed(w http.ResponseWriter, req *http.Request) {
	fd := formData{}

	if req.Method == http.MethodPost {
		fd.SummonerName = req.FormValue("summonerName")
		fd.ChampionName = req.FormValue("championName")
		fd.Region = req.FormValue("region")
	}

	// Gets the recent performance
	perf, err := controller.GetRecentPerformance(&fd.Region, fd.SummonerName)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	errTemplate := tpl.ExecuteTemplate(w, "memed.gohtml", perf)
	if errTemplate != nil {
		log.Println(errTemplate)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/memed", memed)
	http.ListenAndServe(":8080", nil)
}
