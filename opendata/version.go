package astiopendata

import (
	"encoding/xml"
	"time"
)

// Version represents the version
type Version string

// GetVersionRequest represents the getVersion request
type GetVersionRequest struct{}

// GetVersionResponse represents getVersion response
type GetVersionResponse struct {
	Version Version  `xml:"http://wsiv.ratp.fr return"`
	XMLName xml.Name `xml:"http://wsiv.ratp.fr getVersionResponse"`
}

// Version returns the version
func (c *Client) Version() (v Version, err error) {
	// Log
	var m = "Retrieve Open Data RATP version"
	c.Logger.Debugf("[Start] %s", m)
	defer func(now time.Time) {
		c.Logger.Debugf("[End] %s in %s", m, time.Since(now))
	}(time.Now())

	// Build request
	var req = Envelope{Body: &EnvelopeBody{GetVersionRequest: &GetVersionRequest{}}}

	// Send request
	var resp Envelope
	if resp, err = c.SendWithMaxRetries("getVersion", req); err != nil {
		return
	}
	v = resp.Body.GetVersionResponse.Version
	return
}
