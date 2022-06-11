package configs

//Ticker define config of time to start cronjob
type Ticker struct {
	DayOfWeek string `default:""`
	Month     string `default:""`
	Day       string `default:""`
	Hours     string `default:""`
	Minute    string `default:""`
}
