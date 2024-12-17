package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go"

	q "microservice/graphql"
	apiErrors "microservice/internal/errors"
)

type measurementParameters struct {
	From    time.Time `from:"from" binding:"ltfield=Until"`
	Until   time.Time `from:"until" binding:"gtfield=From"`
	Station *string   `from:"station"`
}

func Measurements(c *gin.Context) {
	var params measurementParameters
	if err := c.ShouldBind(&params); err != nil {
		c.Abort()
		apiErrors.ErrInvalidMeasurementParameters.Emit(c)
		return
	}

	measurements, err := q.Query{}.Measurements(&q.MeasurementArguments{
		From:    &graphql.Time{Time: params.From},
		Until:   &graphql.Time{Time: params.Until},
		Station: params.Station,
	})

	if err != nil {
		c.Abort()
		_ = c.Error(err)
		return
	}

	c.JSON(200, measurements)
}
