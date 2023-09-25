package main

import (
	"time"

	"github.com/lewiscasewell/tolling/types"
	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next DataProducer
}

func NewLogMiddleware(next DataProducer) *LogMiddleware {
	return &LogMiddleware{
		next: next,
	}
}

func (lm *LogMiddleware) ProduceData(data types.OBUData) error {
	go func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"obuID":     data.OBUID,
			"latitude":  data.Latitude,
			"longitude": data.Longitude,
			"duration":  time.Since(start),
		}).Info("Producing data to Kafka")
	}(time.Now())
	return lm.next.ProduceData(data)
}
