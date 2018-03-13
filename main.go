package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type formData struct {
	SummonerName string
	ChampionName string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
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
	// s, err := GetSummoner(fd.SummonerName)
	// if err != nil {
	// 	log.Println(err)
	// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// 	return
	// }

	err := tpl.ExecuteTemplate(w, "memed.gohtml", fd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func main() {
	log.Println("Freak Face 3.")
	http.HandleFunc("/", index)
	http.HandleFunc("/memed", memed)
	http.ListenAndServe(":8080", nil)
}
