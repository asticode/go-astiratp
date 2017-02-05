package astiopendata

import (
	"encoding/xml"
	"time"
)

// Station represents a station
type Station struct {
	ID                 string      `xml:"id"`
	GeoPointOneWay     GeoPoint    `xml:"geoPointA"`
	GeoPointReturnTrip GeoPoint    `xml:"geoPointR"`
	Line               Line        `xml:"line"`
	Name               string      `xml:"name"`
	StationArea        StationArea `xml:"stationArea"`
}

// StationArea represents a station area
type StationArea struct {
	ID string `xml:"id"`
}

// GetStationsRequest represents the getStations request
type GetStationsRequest struct {
	Limit int `xml:"http://wsiv.ratp.fr limit"`
}

// GetStationsResponse represents the getStations response
type GetStationsResponse struct {
	Result  ResultStation `xml:"http://wsiv.ratp.fr return"`
	XMLName xml.Name      `xml:"http://wsiv.ratp.fr getStationsResponse"`
}

// ResultStation represents the result of the response of the getStations action
type ResultStation struct {
	Stations []Station `xml:"http://wsiv.ratp.fr/xsd stations"`
}

// Stations returns the stations
func (c *Client) Stations() (s []Station, err error) {
	// Log
	var m = "Retrieve Open Data RATP stations"
	c.Logger.Debugf("[Start] %s", m)
	defer func(now time.Time) {
		c.Logger.Debugf("[End] %s in %s", m, time.Since(now))
	}(time.Now())

	// Build request
	var req = Envelope{Body: &EnvelopeBody{GetStationsRequest: &GetStationsRequest{Limit: 100000}}}

	// Send request
	var resp Envelope
	if resp, err = c.SendWithMaxRetries("getStations", req); err != nil {
		return
	}
	s = resp.Body.GetStationsResponse.Result.Stations
	return
}
