package model

import "time"

const DEFAULT_DATE_TIME_FORMATTER = "2006-01-02"
const REPRESENTATION_DATE_TIME_FORMATTER = "02/01/2006"

type Date struct {
	Day time.Time
}

func (d *Date) FormattedDate(pattern *string) string {
	return d.Day.Format(patternOrDefault(pattern))
}

func DateFrom(date string, pattern *string) (*Date, error) {
	parse, err := time.Parse(patternOrDefault(pattern), date)
	return &Date{Day: parse}, err
}

func patternOrDefault(pattern *string) string {
	actualPattern := DEFAULT_DATE_TIME_FORMATTER
	if pattern != nil {
		actualPattern = *pattern
	}
	return actualPattern
}
