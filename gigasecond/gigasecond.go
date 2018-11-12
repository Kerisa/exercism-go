package gigasecond

import "time"

func AddGigasecond(t time.Time) time.Time {
    result := t.Add(1e9 * time.Second)
	return result
}
