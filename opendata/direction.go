package astiopendata

import "fmt"

// Direction represents a direction
type Direction struct {
	Name              string
	Line              Line
	EndOfLineStations []Station
	Way               DirectionWay
}

// DirectionWay represents a direction way
type DirectionWay string

// Direction ways
const (
	DirectionWayOneWay     DirectionWay = "one way"
	DirectionWayReturnTrip DirectionWay = "return trip"
)

// Direction ways mapping
var directionWaysMapping = map[string]DirectionWay{
	"A": DirectionWayOneWay,
	"R": DirectionWayReturnTrip,
}

// UnmarshalText allows DirectionWay to implement the TextUnmarshaler interface
func (w *DirectionWay) UnmarshalText(b []byte) error {
	var ok bool
	if *w, ok = directionWaysMapping[string(b)]; !ok {
		return fmt.Errorf("Invalid direction way %s", string(b))
	}
	return nil
}
