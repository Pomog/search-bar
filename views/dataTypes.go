package views

import (
	"groupie-tracker/get_info"
)

type DataForFilter struct {
	MinYear     int
	MaxYear     int
	TeamMembers int
	Location    string
}

type ViewData struct {
	Artists         []get_info.Artist
	LenArtists      int
	Name            string
	MaxMembers      int
	MaxMemberOption []int
	MinCreationDate int
	MaxCreationDate int
	MinYearAlbum    int
	MaxYearAlbum    int
	ListOfLocs      []string
	SelectedLoc     string
}
