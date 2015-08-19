package mpower

import (
	"github.com/jmcvetta/napping"
	"net/http"
	"os"
)

const (
	baseURLLive = "https://app.mpowerpayments.com/api/v1"
	baseURLTest = "https://app.mpowerpayments.com/sandbox-api/v1"
)

// Get an environment variable or return `def` string as the default
//
//    str := env("MP-Master-Key")
func env(name string) string {
	return os.Getenv(name)
}

// MPower holds the setup and store data
// It includes all instances of the MPower API
type MPower struct {
	setup   *Setup
	store   *Store
	baseURL string
	Session *napping.Session
}

// NewRequest is the method provided for all the api services to used
func (mp *MPower) NewRequest(method, url string, payload, result interface{}, header *http.Header) (resp *napping.Response, err error) {
	request := &napping.Request{
		Method:  method,
		Url:     url,
		Payload: payload,
		Result:  result,
	}

	if header != nil {
		mp.Session.Header = header
		defer func() {
			mp.Session.Header = new(http.Header)
			for key, val := range mp.setup.Headers {
				mp.Session.Header.Add(key, val)
			}
		}()
	}

	resp, err = mp.Session.Send(request)
	return
}

// NewMPower creates a new MPower
func NewMPower(setup *Setup, store *Store, mode string) *MPower {
	mp := &MPower{
		setup: setup,
		store: store,
	}

	if mode == "live" {
		mp.baseURL = baseURLLive
	} else {
		mp.baseURL = baseURLTest
	}

	for key, val := range setup.Headers {
		mp.Session.Header.Add(key, val)
	}

	return mp
}
