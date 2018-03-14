package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/joshuaj1397/LoLMemes/riotapi"
)

var tpl *template.Template

type formData struct {
	SummonerName string
	ChampionName string
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
	}

	// Construct a new Summoner
	s, err := riotapi.GetSummoner(fd.SummonerName)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	errTemplate := tpl.ExecuteTemplate(w, "memed.gohtml", s)
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
