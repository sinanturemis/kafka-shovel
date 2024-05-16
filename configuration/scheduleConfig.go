package configuration

import "time"

type ScheduleConfig struct {
	Interval time.Duration `json:"interval"`
}
