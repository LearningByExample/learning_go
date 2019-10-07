package racer

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string, err error) {

	err = nil

	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		winner = a
	} else {
		winner = b
	}

	return
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func RacerWithSelect(a, b string) (winner string, err error) {

	err = nil

	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(5 * time.Second):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}

	return "", fmt.Errorf("unknow error")
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()

	return ch
}
