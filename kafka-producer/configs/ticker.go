package configs

//Ticker define config of time to start cronjob
type Ticker struct {
	DayOfWeek int `default:""`
	Year      int `default:""`
	Month     int `default:""`
	Day       int `default:""`
	Hours     int `default:""`
	Minute    int `default:""`
}
