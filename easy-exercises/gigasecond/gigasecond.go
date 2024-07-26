// Package gigasecond contains AddGigasecond method
// https://golang.org/doc/effective_go.html#commentary
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond adds 1000 000 000 seconds to the given time
func AddGigasecond(t time.Time) time.Time {

	t = t.Add(time.Second * 1000000000) 
	return t
}
