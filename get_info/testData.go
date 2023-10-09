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

func getRelations() string {
	return `
    [
        {
            "id": 1,
            "datesLocations": {
                "dunedin-new_zealand": [
                    "10-02-2020"
                ],
                "georgia-usa": [
                    "22-08-2019"
                ],
                "los_angeles-usa": [
                    "20-08-2019"
                ],
                "nagoya-japan": [
                    "30-01-2019"
                ],
                "north_carolina-usa": [
                    "23-08-2019"
                ],
                "osaka-japan": [
                    "28-01-2020"
                ],
                "penrose-new_zealand": [
                    "07-02-2020"
                ],
                "saitama-japan": [
                    "26-01-2020"
                ]
            }
        },
        {
            "id": 2,
            "datesLocations": {
                "noumea-new_caledonia": [
                    "15-11-2019"
                ],
                "papeete-french_polynesia": [
                    "16-11-2019"
                ],
                "playa_del_carmen-mexico": [
                    "05-12-2019",
                    "06-12-2019",
                    "07-12-2019",
                    "08-12-2019",
                    "09-12-2019"
                ]
            }
        },
        {
            "id": 3,
            "datesLocations": {
                "lausanne-switzerland": [
                    "25-09-1994"
                ],
                "london-uk": [
                    "10-05-2007",
                    "02-07-2005",
                    "29-10-1994",
                    "28-10-1994",
                    "27-10-1994",
                    "26-10-1994",
                    "23-10-1994",
                    "22-10-1994",
                    "21-10-1994",
                    "20-10-1994",
                    "19-10-1994",
                    "17-10-1994",
                    "16-10-1994",
                    "15-10-1994",
                    "14-10-1994",
                    "13-10-1994",
                    "12-10-1994"
                ]
            }
        }
    ]
    `
}
