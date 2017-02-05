package astiopendata_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/asticode/go-astiratp/opendata"
	"github.com/stretchr/testify/assert"
)

func TestClient_GeoPoint(t *testing.T) {
	var c = astiopendata.New(astiopendata.Configuration{})
	var bodyRequest []byte
	var bodyResponse = []byte(`<?xml version='1.0' encoding='UTF-8'?><soapenv:Envelope xmlns:soapenv="http://www.w3.org/2003/05/soap-envelope"><soapenv:Body><ns2:getGeoPointsResponse xmlns:ns2="http://wsiv.ratp.fr"><ns2:return><ns1:id xmlns:ns1="http://wsiv.ratp.fr/xsd">2095</ns1:id><ns1:name xmlns:ns1="http://wsiv.ratp.fr/xsd">Antony</ns1:name><ns1:type xmlns:ns1="http://wsiv.ratp.fr/xsd">station</ns1:type><ns1:x xmlns:ns1="http://wsiv.ratp.fr/xsd">597398.0</ns1:x><ns1:y xmlns:ns1="http://wsiv.ratp.fr/xsd">2417412.0</ns1:y></ns2:return></ns2:getGeoPointsResponse></soapenv:Body></soapenv:Envelope>`)
	astiopendata.Send = func(req *http.Request, httpClient *http.Client) (resp *http.Response, err error) {
		bodyRequest, err = ioutil.ReadAll(req.Body)
		return &http.Response{Body: ioutil.NopCloser(bytes.NewReader(bodyResponse)), StatusCode: http.StatusOK}, nil
	}
	g, err := c.GeoPoint("2095")
	assert.NoError(t, err)
	assert.Equal(t, "<Envelope xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Body xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><getGeoPoints xmlns=\"http://wsiv.ratp.fr\"><gp xmlns=\"http://wsiv.ratp.fr\"><id xmlns=\"http://wsiv.ratp.fr/xsd\">2095</id><type xmlns=\"http://wsiv.ratp.fr/xsd\">station</type></gp><limit xmlns=\"http://wsiv.ratp.fr\">100000</limit></getGeoPoints></Body></Envelope>", string(bodyRequest))
	assert.Equal(t, astiopendata.GeoPoint{ID: "2095", Name: "Antony", NameSuffix: "", Type: "station", X: 597398, Y: 2417412}, g)
}
