package types

import (
	"time"

	"github.com/prometheus/prometheus/model/timestamp"
)

type QueryTimeRange struct {
	StartT               int64 // Start timestamp, in milliseconds since Unix epoch.
	EndT                 int64 // End timestamp, in milliseconds since Unix epoch.
	IntervalMilliseconds int64 // Range query interval, or 1 for instant queries. Note that this is deliberately different to parser.EvalStmt.Interval for instant queries (where it is 0) to simplify some loop conditions.

	StepCount int64 // 1 for instant queries.
}

func NewInstantQueryTimeRange(t time.Time) QueryTimeRange {
	ts := timestamp.FromTime(t)

	return QueryTimeRange{
		StartT:               ts,
		EndT:                 ts,
		IntervalMilliseconds: 1,
		StepCount:            1,
	}
}

func NewRangeQueryTimeRange(start time.Time, end time.Time, interval time.Duration) QueryTimeRange {
	startT := timestamp.FromTime(start)
	endT := timestamp.FromTime(end)
	IntervalMilliseconds := interval.Milliseconds()

	return QueryTimeRange{
		StartT:               startT,
		EndT:                 endT,
		IntervalMilliseconds: IntervalMilliseconds,
		StepCount:            ((endT - startT) / IntervalMilliseconds) + 1,
	}
}

// PointIndex returns the index in the QueryTimeRange that the timestamp, t, falls on.
// t must be in line with IntervalMs (ie the step).
func (q *QueryTimeRange) PointIndex(t int64) int64 {
	return (t - q.StartT) / q.IntervalMilliseconds
}

// IndexTime returns the timestamp that the point index, p, falls on.
// p must be less than StepCount
func (q *QueryTimeRange) IndexTime(p int64) int64 {
	return q.StartT + p*q.IntervalMilliseconds
}
