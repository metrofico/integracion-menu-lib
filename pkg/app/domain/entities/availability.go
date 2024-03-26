package entities

type Availability struct {
	DayOfWeek   string `json:"day_of_week" bson:"day_of_week"`
	TimePeriods string `json:"time_periods" bson:"time_periods"`
}
