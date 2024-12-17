package routes

import (
	"microservice/graphql"

	"github.com/gin-gonic/gin"
)

func Locations(c *gin.Context) {
	var graphql graphql.Query
	stations, err := graphql.Stations()
	if err != nil {
		c.Abort()
		_ = c.Error(err)
		return
	}

	c.JSON(200, stations)
}
