package times

import "time"

func StringToDuration(t string) (time.Duration, error) {
	return time.ParseDuration(t)
}
