package routes

import (
	"microservice/graphql"

	"github.com/gin-gonic/gin"
)

func StationData(c *gin.Context) {
	stationID := c.Param("stationID")

	station, err := graphql.Query{}.Station(struct{ StationID string }{stationID})
	if err != nil {
		c.Abort()
		_ = c.Error(err)
		return
	}

	c.JSON(200, station)
}
