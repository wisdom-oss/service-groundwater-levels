package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/go-chi/chi/v5"
	errorMiddleware "github.com/wisdom-oss/microservice-middlewares/v5/error"

	"microservice/globals"
	"microservice/types"
)

func SingleStation(w http.ResponseWriter, r *http.Request) {
	errorHandler := r.Context().Value(errorMiddleware.ChannelName).(chan<- interface{})
	stationID := chi.URLParam(r, "stationID")
	rawQuery, err := globals.SqlQueries.Raw("get-single-station")
	if err != nil {
		errorHandler <- fmt.Errorf("unable to load SQL query 'get-measurement-stations': %w", err)
		return
	}

	var measurementStation types.Station
	err = pgxscan.Get(r.Context(), globals.Db, &measurementStation, rawQuery, stationID)
	if err != nil {
		errorHandler <- fmt.Errorf("error while parsing measurement stations: %w", err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(measurementStation)
	if err != nil {
		errorHandler <- fmt.Errorf("error while encoding measurement stations: %w", err)
		return
	}
}
