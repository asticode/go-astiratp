package astiopendata

import (
	"encoding/xml"
	"time"
)

// Mission represents a mission
type Mission struct {
	Code             string
	Direction        Direction
	ID               string
	Line             Line
	EndOfLineStation Station
	Perturbations    []string // TODO Parse message
	Stations         []MissionStation
}

// MissionStation represents a mission station
type MissionStation struct {
	Date     time.Time
	Message  string // TODO Parse message
	Platform string
	Station  Station
	Stop     bool
}

// ResponseMissions represents response of the getMission action
type ResponseMissions struct {
	Missions []Mission `xml:"http://wsiv.ratp.fr return"`
	XMLName  xml.Name  `xml:"http://wsiv.ratp.fr getMissionResponse"`
}

// TODO Missions returns the missions
func (c *Client) Missions() (m []Mission, err error) {
	return
}
