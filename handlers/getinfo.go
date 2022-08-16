package handlers

import (
	"encoding/json"
	"net/http"
)

func GetData(first, second string) ([]Response, int) {
	var result []Response
	var rels Relations
	artists, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, http.StatusInternalServerError
	}
	defer artists.Body.Close()
	if err := json.NewDecoder(artists.Body).Decode(&result); err != nil {
		return nil, http.StatusInternalServerError
	}
	relations, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return nil, http.StatusInternalServerError
	}
	defer relations.Body.Close()
	if err := json.NewDecoder(relations.Body).Decode(&rels); err != nil {
		return nil, http.StatusInternalServerError
	}

	for i, v := range rels.Index {
		result[i].PlacesDates = v.Dateslocations
	}

	return result, http.StatusOK
}
