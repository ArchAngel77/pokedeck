package pokeapi

type PokemonCity struct {
    Count    int            `json:"count"`
    Next     *string        `json:"next"`
    Previous *string        `json:"previous"`
    Results  []CityLocation `json:"results"`
}

type CityLocation struct {
    Name string `json:"name"`
    Url  string `json:"url"`
}

type Location struct {
    Name	      string			`json:"name"`
    PokemonEncounters []PokemonEncounter	`json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name 	       string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
        Weight         int    `json:"weight"`
        Stats []struct {
            BaseStat int `json:"base_stat"`
            Stat     struct {
                Name string `json:"name"`
            } `json:"stat"`
        } `json:"stats"`
        Types []struct {
            Type struct {
                Name string `json:"name"`
            } `json:"type"`
        } `json:"types"`
}

