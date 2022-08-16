package handlers

type Response struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Relations    string   `json:"relations"`
	PlacesDates  map[string][]string
}

type Relations struct {
	Index []struct {
		Dateslocations map[string][]string `json: "datesLocations"`
	} `json:"index"`
}

type AllResponse struct {
	AllArtists   []Response
	FoundArtists []Response
}

type ErrorBody struct {
	Status  int
	Message string
}
