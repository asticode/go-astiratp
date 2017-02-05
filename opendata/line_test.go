package astiopendata_test

import (
	"net/http"
	"testing"

	"bytes"
	"io/ioutil"

	"github.com/asticode/go-astiratp/opendata"
	"github.com/stretchr/testify/assert"
)

func TestLineRealm_UnmarshalText(t *testing.T) {
	var r astiopendata.LineRealm
	err := r.UnmarshalText([]byte("r"))
	assert.NoError(t, err)
	assert.Equal(t, astiopendata.LineRealmRealTime, r)
	err = r.UnmarshalText([]byte("invalid"))
	assert.Error(t, err)
}

func TestClient_Lines(t *testing.T) {
	var c = astiopendata.New(astiopendata.Configuration{})
	var bodyRequest []byte
	var bodyResponse = []byte(`<?xml version='1.0' encoding='UTF-8'?><soapenv:Envelope xmlns:soapenv="http://www.w3.org/2003/05/soap-envelope"><soapenv:Body><ns2:getLinesResponse xmlns:ns2="http://wsiv.ratp.fr"><ns2:return><ns1:code xmlns:ns1="http://wsiv.ratp.fr/xsd">1</ns1:code><ns1:codeStif xmlns:ns1="http://wsiv.ratp.fr/xsd">100110001</ns1:codeStif><ns1:id xmlns:ns1="http://wsiv.ratp.fr/xsd">62</ns1:id><ns1:image xmlns:ns1="http://wsiv.ratp.fr/xsd">m1.gif</ns1:image><ns1:name xmlns:ns1="http://wsiv.ratp.fr/xsd">La Defense / Chateau de Vincennes</ns1:name><ns1:realm xmlns:ns1="http://wsiv.ratp.fr/xsd">t</ns1:realm><reseau xmlns="http://wsiv.ratp.fr/xsd"><code>metro</code><id>1-metro</id><image>p_met.gif</image><name>Métro</name></reseau></ns2:return></ns2:getLinesResponse></soapenv:Body></soapenv:Envelope>`)
	astiopendata.Send = func(req *http.Request, httpClient *http.Client) (resp *http.Response, err error) {
		bodyRequest, _ = ioutil.ReadAll(req.Body)
		return &http.Response{Body: ioutil.NopCloser(bytes.NewReader(bodyResponse)), StatusCode: http.StatusOK}, nil
	}
	l, err := c.Lines()
	assert.NoError(t, err)
	assert.Equal(t, "<Envelope xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Body xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><getLines xmlns=\"http://wsiv.ratp.fr\"></getLines></Body></Envelope>", string(bodyRequest))
	assert.Equal(t, []astiopendata.Line{astiopendata.Line{Code: "1", CodeSTIF: "100110001", ID: "62", Image: "m1.gif", Name: "La Defense / Chateau de Vincennes", Network: astiopendata.Network{Code: astiopendata.NetworkCodeMetro, ID: "1-metro", Image: "p_met.gif", Name: "Métro"}, Realm: astiopendata.LineRealmTheoretical}}, l)
}
