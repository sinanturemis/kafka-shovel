package enums

type RunningState int

const (
	Stopping RunningState = iota
	Stopped  RunningState = iota
	Running  RunningState = iota
)
