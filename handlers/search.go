package handlers

import (
	"net/http"
	"strconv"
	"strings"
)

func Search(search_input string) (AllResponse, int) {
	result, err := GetData("https://groupietrackers.herokuapp.com/api/artists", "https://groupietrackers.herokuapp.com/api/relation")
	var myResponse AllResponse
	if err != 200 {
		return myResponse, err
	}
	var data []Response
	var isFound bool = false
	for i := range result {
		if caseIns(result[i].Name, search_input) {
			isFound = true
			data = append(data, result[i])
			continue
		}

		for j := range result[i].PlacesDates {
			if caseIns(j, search_input) {
				isFound = true
				data = append(data, result[i])
				break
			}
		}

		if search_input == result[i].FirstAlbum {
			isFound = true
			data = append(data, result[i])
			continue
		}

		if search_input == strconv.Itoa(result[i].CreationDate) {
			isFound = true
			data = append(data, result[i])
			continue
		}
		for j := range result[i].Members {
			if caseIns(result[i].Members[j], search_input) {
				isFound = true
				data = append(data, result[i])
				continue
			}
		}
	}
	myResponse.AllArtists = result
	myResponse.FoundArtists = checkSearch(data)
	if isFound {
		return myResponse, http.StatusOK
	}
	return myResponse, http.StatusNotFound
}

func checkSearch(data []Response) []Response {
	var newData []Response
	var contains bool = false
	for i := range data {
		for j := range newData {
			if data[i].Name == newData[j].Name {
				contains = true
				break
			}
		}
		if !contains {
			newData = append(newData, data[i])
		}
		contains = false
	}
	return newData
}

func caseIns(str1, str2 string) bool {
	return strings.Contains(
		strings.ToLower(str1),
		strings.ToLower(str2),
	)
}
