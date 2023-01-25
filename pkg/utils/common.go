package utils

import "time"

func DoWithTries(fn func() error, attemts int, delay time.Duration) (err error) {
	for attemts > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			attemts--
			continue
		}
		return nil
	}
	return
}
