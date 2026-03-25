
package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c Client) GetLocation(locationName string) (Location, error) {
        url := "https://pokeapi.co/api/v2/location-area/" + locationName
        if val, ok := c.cache.Get(url); ok {
          locationResp := Location{}
          err := json.Unmarshal(val, &locationResp)
          if err != nil {
            return Location{}, err
          }
          return locationResp, nil
        }
        resp, err := http.Get(url)
        if err != nil {
                return Location{}, err
        }
        defer resp.Body.Close()
        body, err := io.ReadAll(resp.Body)
        if err != nil {
                return Location{}, err
        }
        c.cache.Add(url, body)
        locationResp := Location{}
        err = json.Unmarshal(body, &locationResp)
        if err != nil {
                return Location{}, err
        }
        return locationResp, nil
}
