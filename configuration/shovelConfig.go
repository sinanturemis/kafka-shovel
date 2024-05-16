package configuration

type ShovelConfig struct {
	ErrorTopic     string         `json:"errorTopic"`
	RetryTopic     string         `json:"retryTopic"`
	Active         bool           `json:"active"`
	RetryCount     int            `json:"retryCount"`
	Interval       string         `json:"interval"`
	ScheduleConfig ScheduleConfig `json:"scheduleConfig"`
}
