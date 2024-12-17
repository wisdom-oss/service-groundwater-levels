package graphql

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/twpayne/go-geom/encoding/geojson"

	"microservice/internal/db"
	"microservice/types"
)

type Station struct {
	WebsiteID *string         `db:"website_id" json:"websiteID"`
	PublicID  *string         `db:"public_id" json:"publicID"`
	Name      *string         `db:"name" json:"name"`
	Operator  *string         `db:"operator" json:"operator"`
	Location  *types.Location `db:"location" json:"location"`
}

func (Query) Station(args struct{ ID string }) (*Station, error) {
	if strings.TrimSpace(args.ID) == "" {
		return nil, errors.New("empty id")
	}

	rawQuery, err := db.Queries.Raw("get-single-station")
	if err != nil {
		return nil, err
	}

	var station types.Station
	err = pgxscan.Get(context.Background(), db.Pool, &station, rawQuery, args.ID)
	if err != nil {
		return nil, err
	}

	point, err := geojson.Encode(*station.Location)
	if err != nil {
		return nil, err
	}

	raw, _ := point.Coordinates.MarshalJSON()

	var coordinates []float64
	err = json.Unmarshal(raw, &coordinates)
	if err != nil {
		return nil, err
	}

	return &Station{
		WebsiteID: station.WebsiteID,
		PublicID:  station.PublicID,
		Name:      station.Name,
		Operator:  station.Operator,
		Location: &types.Location{
			Type:        "Point",
			Coordinates: coordinates,
		},
	}, nil

}

func (Query) Stations() ([]Station, error) {
	rawQuery, err := db.Queries.Raw("get-measurement-stations")
	if err != nil {
		return nil, err
	}

	var stations []types.Station
	err = pgxscan.Select(context.Background(), db.Pool, &stations, rawQuery)
	if err != nil {
		return nil, err
	}

	var output []Station

	for _, s := range stations {
		point, err := geojson.Encode(*s.Location)
		if err != nil {
			return nil, err
		}

		raw, _ := point.Coordinates.MarshalJSON()

		var coordinates []float64
		err = json.Unmarshal(raw, &coordinates)
		if err != nil {
			return nil, err
		}

		output = append(output, Station{
			WebsiteID: s.WebsiteID,
			PublicID:  s.PublicID,
			Name:      s.Name,
			Operator:  s.Operator,
			Location: &types.Location{
				Type:        "Point",
				Coordinates: coordinates,
			},
		})
	}

	return output, nil

}
