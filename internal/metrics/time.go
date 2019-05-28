package metrics

import "time"

type (
	//TimeService wraps common time related operations
	TimeService interface {
		NewTicker(duration time.Duration) *time.Ticker
		Now() time.Time
	}

	timeService struct {
	}
)

// MakeTimeService creates a new time service
func MakeTimeService() TimeService {
	return &timeService{}
}

func (ts *timeService) NewTicker(duration time.Duration) *time.Ticker {
	return time.NewTicker(duration)
}

func (ts *timeService) Now() time.Time {
	return time.Now()
}
