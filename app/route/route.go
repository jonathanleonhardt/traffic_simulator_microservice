package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	ID       string
	ClientID string
	Position []Position
}

type Position struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type PartialRoutePosition struct {
	ID       string    `json:"routeID"`
	ClientID string    `json:"clientId"`
	Position []float64 `json:"position"`
	Finished bool      `json:"finished"`
}

func (r *Route) LoadPositions() error {
	if r.ID == "" {
		return errors.New("Route ID not informed")
	}

	file, err := os.Open("destinations/" + r.ID + ".txt")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		lat, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return err
		}
		long, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			return err
		}
		r.Position = append(r.Position, Position{
			Lat:  lat,
			Long: long,
		})
	}

	return nil
}

func (r *Route) ExportJsonPositions() ([]string, error) {
	var route PartialRoutePosition
	var result []string
	total := len(r.Position)

	for k, v := range r.Position {
		route.ID = r.ID
		route.ClientID = r.ClientID
		route.Position = []float64{v.Lat, v.Long}
		route.Finished = false
		if total-1 == k {
			route.Finished = true
		}
		jsonRoute, err := json.Marshal(route)
		if err != nil {
			return nil, err
		}
		result = append(result, string(jsonRoute))
	}
	return result, nil
}
