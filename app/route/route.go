package route

import (
	"bufio"
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
	Lat  float64
	Long float64
}

func (r *Route) LoadPositions() error {
	if r.ID == "" {
		return errors.New("Route ID not informed")
	}

	file, err := os.Open("destinations" + r.ID + ".txt")
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
