package pokeapi

const (
	baseURL = "https://pokeapi.co/api/v2"
)

type Locations struct {
	Count   int     `json:"count"`
	NextURL *string `json:"next"`
	PrevURL *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}
