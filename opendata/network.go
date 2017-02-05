package astiopendata

import "fmt"

// Network represents a network
type Network struct {
	Code  NetworkCode `xml:"code"`
	ID    string      `xml:"id"`
	Image string      `xml:"image"`
	Name  string      `xml:"name"`
}

// NetworkCode represents a network code
type NetworkCode string

// Network codes
const (
	NetworkCodeMetro          NetworkCode = "Metro"
	NetworkCodeOptile         NetworkCode = "Optile"
	NetworkCodeOtherBuses     NetworkCode = "Other buses"
	NetworkCodeRATPBus        NetworkCode = "RATP Bus"
	NetworkCodeRATPNoctilien  NetworkCode = "RATP Noctilien"
	NetworkCodeRER            NetworkCode = "RER"
	NetworkCodeSNCFNoctilien  NetworkCode = "SNCF Noctilien"
	NetworkCodeSNCFTransilien NetworkCode = "SNCF Transilien"
	NetworkCodeTramway        NetworkCode = "Tramway"
)

// Network codes mapping
var networkCodesMapping = map[string]NetworkCode{
	"autre":         NetworkCodeOtherBuses,
	"busratp":       NetworkCodeRATPBus,
	"metro":         NetworkCodeMetro,
	"noctilienratp": NetworkCodeRATPNoctilien,
	"noctiliensncf": NetworkCodeSNCFNoctilien,
	"optile":        NetworkCodeOptile,
	"rer":           NetworkCodeRER,
	"sncf":          NetworkCodeSNCFTransilien,
	"tram":          NetworkCodeTramway,
}

// NetworkCodeFromString returns a network code from a string
func NetworkCodeFromString(i string) (c NetworkCode, err error) {
	var ok bool
	if c, ok = networkCodesMapping[i]; !ok {
		err = fmt.Errorf("Invalid network code %s", i)
	}
	return
}

// UnmarshalText allows NetworkCode to implement the TextUnmarshaler interface
func (c *NetworkCode) UnmarshalText(b []byte) (err error) {
	*c, err = NetworkCodeFromString(string(b))
	return
}
