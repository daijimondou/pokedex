package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocDetail struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *Client) GetExplore(name string) (LocDetail, error) {
	actualURL := baseURL + "/location-area/" + name

	if val, ok := c.cache.Get(actualURL); ok {
		locDetailRes := LocDetail{}
		err := json.Unmarshal(val, &locDetailRes)
		if err != nil {
			return LocDetail{}, err
		}

		return locDetailRes, nil
	}

	req, err := http.NewRequest("GET", actualURL, nil)
	if err != nil {
		return LocDetail{}, err
	}
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return LocDetail{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocDetail{}, err
	}

	locRes := LocDetail{}
	if err = json.Unmarshal(data, &locRes); err != nil {
		return LocDetail{}, err
	}
	c.cache.Add(actualURL, data)
	return locRes, nil
}
