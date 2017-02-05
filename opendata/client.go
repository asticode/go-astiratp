package astiopendata

import (
	"errors"
	"net/http"
	"time"

	"github.com/rs/xlog"
)

// Constants
const (
	endpointURL = "http://opendata-tr.ratp.fr/wsiv/services/Wsiv"
)

// Vars
var (
	ErrNotFound = errors.New("Not found")
)

// Client represents an entity capable of interacting with the Open Data RATP API
type Client struct {
	Client               *http.Client
	Logger               xlog.Logger
	MaxRequestsPerSecond int
	RetryMax             int
	RetrySleep           time.Duration
}

// New creates a new Client
func New(c Configuration) *Client {
	var o = &Client{
		Client:     &http.Client{},
		Logger:     xlog.NopLogger,
		RetryMax:   c.RetryMax,
		RetrySleep: c.RetrySleep,
	}
	return o
}
