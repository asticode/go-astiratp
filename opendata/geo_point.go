package astiopendata

import (
	"encoding/xml"
	"fmt"
	"time"
)

// GeoPoint represents a geographical point
type GeoPoint struct {
	ID         string       `xml:"id"`
	Name       string       `xml:"name"`
	NameSuffix string       `xml:"nameSuffix"`
	Type       GeoPointType `xml:"type"`
	X          float64      `xml:"x"`
	Y          float64      `xml:"y"`
}

// GeoPointType represents a geo point type
type GeoPointType string

// GeoPoint types
const (
	GeoPointTypeAddress GeoPointType = "address"
	GeoPointTypeSpot    GeoPointType = "spot"
	GeoPointTypeStation GeoPointType = "station"
)

// GeoPoint types mapping
var geoPointTypeMapping = map[string]GeoPointType{
	"adresse": GeoPointTypeAddress,
	"lieu":    GeoPointTypeSpot,
	"station": GeoPointTypeStation,
}

// UnmarshalText allows DirectionWay to implement the TextUnmarshaler interface
func (t *GeoPointType) UnmarshalText(b []byte) error {
	var ok bool
	if *t, ok = geoPointTypeMapping[string(b)]; !ok {
		return fmt.Errorf("Invalid geo point type %s", string(b))
	}
	return nil
}

// GetGeoPointsRequest represents the getGeoPoints request
type GetGeoPointsRequest struct {
	GeoPoint GeoPointRequestFilter `xml:"http://wsiv.ratp.fr gp"`
	Limit    int                   `xml:"http://wsiv.ratp.fr limit"`
}

// GeoPointRequestFilter represents the gp filter
type GeoPointRequestFilter struct {
	ID   string       `xml:"http://wsiv.ratp.fr/xsd id"`
	Type GeoPointType `xml:"http://wsiv.ratp.fr/xsd type"`
}

// GetGeoPointsResponse represents the getGeoPoints response
type GetGeoPointsResponse struct {
	GeoPoints []GeoPoint `xml:"http://wsiv.ratp.fr return"`
	XMLName   xml.Name   `xml:"http://wsiv.ratp.fr getGeoPointsResponse"`
}

// GeoPoint returns the station geo point of a specific ID
func (c *Client) GeoPoint(id string) (g GeoPoint, err error) {
	// Log
	var m = "Retrieve Open Data RATP geo points"
	c.Logger.Debugf("[Start] %s", m)
	defer func(now time.Time) {
		c.Logger.Debugf("[End] %s in %s", m, time.Since(now))
	}(time.Now())

	// Build request
	var req = Envelope{Body: &EnvelopeBody{GetGeoPointsRequest: &GetGeoPointsRequest{GeoPoint: GeoPointRequestFilter{ID: id, Type: GeoPointTypeStation}, Limit: 100000}}}

	// Send request
	var resp Envelope
	if resp, err = c.SendWithMaxRetries("getGeoPoints", req); err != nil {
		return
	}

	// Update geo point
	if len(resp.Body.GetGeoPointsResponse.GeoPoints) > 0 && resp.Body.GetGeoPointsResponse.GeoPoints[0].ID == id {
		g = resp.Body.GetGeoPointsResponse.GeoPoints[0]
		return
	}
	err = ErrNotFound
	return
}
