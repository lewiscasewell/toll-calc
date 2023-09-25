package main

import (
	"time"

	"github.com/lewiscasewell/tolling/types"
	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next Aggregator
}

func NewLogMiddleware(next Aggregator) *LogMiddleware {
	return &LogMiddleware{
		next: next,
	}
}

func (lm *LogMiddleware) AggregateDistance(distance types.Distance) (err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"distance": distance,
			"duration": time.Since(start),
			"err":      err,
		}).Info("Aggregating distance")
	}(time.Now())
	err = lm.next.AggregateDistance(distance)
	return err
}

func (lm *LogMiddleware) CalculateInvoice(obuID int) (invoice types.Invoice, err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"obuID":    obuID,
			"duration": time.Since(start),
			"err":      err,
		}).Info("Calculating invoice")
	}(time.Now())
	invoice, err = lm.next.CalculateInvoice(obuID)
	return
}
