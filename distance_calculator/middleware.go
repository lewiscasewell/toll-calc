package main

import (
	"time"

	"github.com/lewiscasewell/tolling/types"
	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next CalculatorServicer
}

func NewLogMiddleware(next CalculatorServicer) *LogMiddleware {
	return &LogMiddleware{
		next: next,
	}
}

func (lm *LogMiddleware) CalculateDistance(data types.OBUData) (dist float64, err error) {

	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"obuID":    data.OBUID,
			"duration": time.Since(start),
			"err":      err,
			"distance": dist,
		}).Info("Calculating distance")
	}(time.Now())
	dist, err = lm.next.CalculateDistance(data)
	return
}
