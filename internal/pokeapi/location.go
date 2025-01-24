package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocations(url *string) (Locations, error) {
	actualURL := baseURL + "/location-area"
	if url != nil {
		actualURL = *url
	}

	if val, ok := c.cache.Get(actualURL); ok {
		locationsRes := Locations{}
		err := json.Unmarshal(val, &locationsRes)
		if err != nil {
			return Locations{}, err
		}

		return locationsRes, nil
	}

	req, err := http.NewRequest("GET", actualURL, nil)
	if err != nil {
		return Locations{}, err
	}
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return Locations{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Locations{}, err
	}

	locRes := Locations{}
	if err = json.Unmarshal(data, &locRes); err != nil {
		return Locations{}, err
	}
	c.cache.Add(actualURL, data)
	return locRes, nil
}
