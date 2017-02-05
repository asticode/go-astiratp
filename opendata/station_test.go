package astiopendata_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/asticode/go-astiratp/opendata"
	"github.com/stretchr/testify/assert"
)

func TestClient_Stations(t *testing.T) {
	var c = astiopendata.New(astiopendata.Configuration{})
	var bodyRequest []byte
	var bodyResponse = []byte(`<?xml version='1.0' encoding='UTF-8'?><soapenv:Envelope xmlns:soapenv="http://www.w3.org/2003/05/soap-envelope"><soapenv:Body><ns2:getStationsResponse xmlns:ns2="http://wsiv.ratp.fr"><ns2:return><stations xmlns="http://wsiv.ratp.fr/xsd"><geoPointA><id>2048</id><name>Argentine</name></geoPointA><geoPointR><id>2098</id><name>Argentine</name></geoPointR><id>2048-2098</id><line><code>1</code><codeStif>100110001</codeStif><id>62</id><image>m1.gif</image><name>La Defense / Chateau de Vincennes</name><realm>t</realm><reseau><code>metro</code><id>1-metro</id><image>p_met.gif</image><name>Métro</name></reseau></line><name>Argentine</name><stationArea><id>114</id></stationArea></stations></ns2:return></ns2:getStationsResponse></soapenv:Body></soapenv:Envelope>`)
	astiopendata.Send = func(req *http.Request, httpClient *http.Client) (resp *http.Response, err error) {
		bodyRequest, err = ioutil.ReadAll(req.Body)
		return &http.Response{Body: ioutil.NopCloser(bytes.NewReader(bodyResponse)), StatusCode: http.StatusOK}, nil
	}
	s, err := c.Stations()
	assert.NoError(t, err)
	assert.Equal(t, "<Envelope xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Body xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><getStations xmlns=\"http://wsiv.ratp.fr\"><limit xmlns=\"http://wsiv.ratp.fr\">100000</limit></getStations></Body></Envelope>", string(bodyRequest))
	assert.Equal(t, []astiopendata.Station{astiopendata.Station{ID: "2048-2098", GeoPointOneWay: astiopendata.GeoPoint{ID: "2048", Name: "Argentine"}, GeoPointReturnTrip: astiopendata.GeoPoint{ID: "2098", Name: "Argentine"}, Line: astiopendata.Line{Code: "1", CodeSTIF: "100110001", ID: "62", Image: "m1.gif", Name: "La Defense / Chateau de Vincennes", Network: astiopendata.Network{Code: astiopendata.NetworkCodeMetro, ID: "1-metro", Image: "p_met.gif", Name: "Métro"}, Realm: astiopendata.LineRealmTheoretical}, Name: "Argentine", StationArea: astiopendata.StationArea{ID: "114"}}}, s)
}
