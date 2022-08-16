package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var All *template.Template

func init() {
	var err error
	All, err = template.ParseFiles("ui/htmlfiles/error.html", "ui/htmlfiles/index.html", "ui/htmlfiles/more-info.html")
	if err != nil {
		log.Fatal(err)
	}
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		errorHeader(w, http.StatusMethodNotAllowed)
		return
	}
	search_word := r.FormValue("search-bar")

	res, status := Search(search_word)
	if status != 200 {
		errorHeader(w, status)
		return
	}

	err := All.ExecuteTemplate(w, "index.html", res)
	if err != nil {
		errorHeader(w, http.StatusInternalServerError)
		return
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var myResponse AllResponse
	if r.URL.Path != "/" {
		errorHeader(w, http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		errorHeader(w, http.StatusMethodNotAllowed)
		return
	}
	result, status := GetData("https://groupietrackers.herokuapp.com/api/artists", "https://groupietrackers.herokuapp.com/api/relation")
	if status != 200 {
		errorHeader(w, status)
		return
	}
	myResponse.AllArtists = result
	myResponse.FoundArtists = result
	err := All.ExecuteTemplate(w, "index.html", myResponse)
	if err != nil {
		errorHeader(w, http.StatusInternalServerError)
		return
	}
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/artist/")
	num, err := strconv.Atoi(id)
	if err != nil {
		errorHeader(w, http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		errorHeader(w, http.StatusMethodNotAllowed)
		return
	}
	result, status := GetData("https://groupietrackers.herokuapp.com/api/artists", "https://groupietrackers.herokuapp.com/api/relation")
	if status != 200 {
		errorHeader(w, http.StatusNotFound)
		return
	}

	for i := range result {
		if result[i].Id == 0 {
			errorHeader(w, http.StatusNotFound)
			return
		}
	}

	err = All.ExecuteTemplate(w, "more-info.html", result[num-1])
	if err != nil {
		errorHeader(w, http.StatusInternalServerError)
		return
	}
}

func errorHeader(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	errH := setError(status)
	err := All.ExecuteTemplate(w, "error.html", errH)
	if err != nil {
		errorHeader(w, http.StatusInternalServerError)
		return
	}
}

func setError(status int) *ErrorBody {
	return &ErrorBody{
		Status:  status,
		Message: http.StatusText(status),
	}
}
