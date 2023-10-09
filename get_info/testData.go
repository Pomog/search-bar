package get_info

func getArtists() string {

	// JSON data representing a slice of artists
	return `
[
	{
		"id": 1,
		"image": "https://groupietrackers.herokuapp.com/api/images/queen.jpeg",
		"name": "Queen",
		"members": [
			"Freddie Mercury",
			"Brian May",
			"John Deacon",
			"Roger Meddows-Taylor",
			"Mike Grose",
			"Barry Mitchell",
			"Doug Fogie"
		],
		"creationDate": 1970,
		"firstAlbum": "14-12-1973",
		"locations": "https://groupietrackers.herokuapp.com/api/locations/1",
		"concertDates": "https://groupietrackers.herokuapp.com/api/dates/1",
		"relations": "https://groupietrackers.herokuapp.com/api/relation/1"
	},
	{
		"id": 2,
		"image": "https://groupietrackers.herokuapp.com/api/images/soja.jpeg",
		"name": "SOJA",
		"members": [
			"Jacob Hemphill",
			"Bob Jefferson",
			"Ryan \"Byrd\" Berty",
			"Ken Brownell",
			"Patrick O'Shea",
			"Hellman Escorcia",
			"Rafael Rodriguez",
			"Trevor Young"
		],
		"creationDate": 1997,
		"firstAlbum": "05-06-2002",
		"locations": "https://groupietrackers.herokuapp.com/api/locations/2",
		"concertDates": "https://groupietrackers.herokuapp.com/api/dates/2",
		"relations": "https://groupietrackers.herokuapp.com/api/relation/2"
	},
	{
		"id": 3,
		"image": "https://groupietrackers.herokuapp.com/api/images/pinkfloyd.jpeg",
		"name": "Pink Floyd",
		"members": [
			"David Gilmour",
			"Syd Barrett",
			"Roger Waters",
			"Richard Wright",
			"Nick Mason",
			"Bob Klose"
		],
		"creationDate": 1965,
		"firstAlbum": "05-08-1967",
		"locations": "https://groupietrackers.herokuapp.com/api/locations/3",
		"concertDates": "https://groupietrackers.herokuapp.com/api/dates/3",
		"relations": "https://groupietrackers.herokuapp.com/api/relation/3"
	}
]
`

}
