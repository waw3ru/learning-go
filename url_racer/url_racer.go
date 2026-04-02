package urlracer

import (
	"errors"
	"net/http"
	"time"
)

var TimeoutError = errors.New("timeout error")

func Racer(a, b string, d time.Duration) (winner string, e error) {
	return ConfigRacer(a, b, d)
}

func ConfigRacer(a, b string, d time.Duration) (winner string, e error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(d):
		return "", TimeoutError
	}
}

func ping(url string) chan struct{} {
	c := make(chan struct{})
	go func() {
		defer close(c)
		http.Get(url)
	}()
	return c
}
