package get_info

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	LocDate      [][]string
	FinalRel     map[string][]string
}
type Relation struct {
	Id            int                 `json:"id"`
	DatesLocation map[string][]string `json:"datesLocations"`
}

const (
	ArtistsAPI   = "https://groupietrackers.herokuapp.com/api/artists"
	RelationsAPI = "https://groupietrackers.herokuapp.com/api/relation"
)

func GetInfoFromAPI(api string, target interface{}) error {
	//Make an http get request to the api
	resp, errAPIrequest := http.Get(api)
	if errAPIrequest != nil {
		log.Println("ERROR making get request:", errAPIrequest, "\nFrom: ", api)
		return errAPIrequest
	}
	defer resp.Body.Close()

	//check the HTTP status code
	if resp.StatusCode != http.StatusOK {
		errStatusCode := fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
		log.Println(errStatusCode)
		return errStatusCode
	}

	//Decode the JSON response into a slice
	errJSONdecodeAtrists := json.NewDecoder(resp.Body).Decode(&target)
	if errJSONdecodeAtrists != nil {
		log.Println("Error decoding JSON:", errJSONdecodeAtrists)
		return errJSONdecodeAtrists
	}

	return nil
}
