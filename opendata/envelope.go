package astiopendata

import "encoding/xml"

// Envelope represents the envelope of the SOAP API
type Envelope struct {
	Body    *EnvelopeBody `xml:"http://www.w3.org/2003/05/soap-envelope Body"`
	XMLName xml.Name      `xml:"http://www.w3.org/2003/05/soap-envelope Envelope"`
}

// EnvelopeBody represents the envelope body
type EnvelopeBody struct {
	GetGeoPointsRequest  *GetGeoPointsRequest  `xml:"http://wsiv.ratp.fr getGeoPoints"`
	GetGeoPointsResponse *GetGeoPointsResponse `xml:"http://wsiv.ratp.fr getGeoPointsResponse"`
	GetLinesRequest      *GetLinesRequest      `xml:"http://wsiv.ratp.fr getLines"`
	GetLinesResponse     *GetLinesResponse     `xml:"http://wsiv.ratp.fr getLinesResponse"`
	GetStationsRequest   *GetStationsRequest   `xml:"http://wsiv.ratp.fr getStations"`
	GetStationsResponse  *GetStationsResponse  `xml:"http://wsiv.ratp.fr getStationsResponse"`
	GetVersionRequest    *GetVersionRequest    `xml:"http://wsiv.ratp.fr getVersion"`
	GetVersionResponse   *GetVersionResponse   `xml:"http://wsiv.ratp.fr getVersionResponse"`
}
