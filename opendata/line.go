package astiopendata

import (
	"encoding/xml"
	"fmt"
	"time"
)

// Line represents a line
type Line struct {
	Code     string    `xml:"http://wsiv.ratp.fr/xsd code"`
	CodeSTIF string    `xml:"http://wsiv.ratp.fr/xsd codeStif"`
	ID       string    `xml:"http://wsiv.ratp.fr/xsd id"`
	Image    string    `xml:"http://wsiv.ratp.fr/xsd image"`
	Name     string    `xml:"http://wsiv.ratp.fr/xsd name"`
	Network  Network   `xml:"http://wsiv.ratp.fr/xsd reseau"`
	Realm    LineRealm `xml:"http://wsiv.ratp.fr/xsd realm"`
}

// LineRealm represents a line realm
type LineRealm string

// Line realms
const (
	LineRealmRealTime    LineRealm = "real time"
	LineRealmTheoretical LineRealm = "theoretical"
)

// Line realms mapping
var lineRealmsMapping = map[string]LineRealm{
	"r": LineRealmRealTime,
	"t": LineRealmTheoretical,
}

// UnmarshalText allows LineRealm to implement the TextUnmarshaler interface
func (r *LineRealm) UnmarshalText(b []byte) error {
	var ok bool
	if *r, ok = lineRealmsMapping[string(b)]; !ok {
		return fmt.Errorf("Invalid line realm %s", string(b))
	}
	return nil
}

// GetLinesRequest represents the getLines request
type GetLinesRequest struct{}

// GetLinesResponse represents getLines response
type GetLinesResponse struct {
	Lines   []Line   `xml:"http://wsiv.ratp.fr return"`
	XMLName xml.Name `xml:"http://wsiv.ratp.fr getLinesResponse"`
}

// Lines returns the lines
func (c *Client) Lines() (l []Line, err error) {
	// Log
	var m = "Retrieve Open Data RATP lines"
	c.Logger.Debugf("[Start] %s", m)
	defer func(now time.Time) {
		c.Logger.Debugf("[End] %s in %s", m, time.Since(now))
	}(time.Now())

	// Build request
	var req = Envelope{Body: &EnvelopeBody{GetLinesRequest: &GetLinesRequest{}}}

	// Send request
	var resp Envelope
	if resp, err = c.SendWithMaxRetries("getLines", req); err != nil {
		return
	}
	l = resp.Body.GetLinesResponse.Lines
	return
}
