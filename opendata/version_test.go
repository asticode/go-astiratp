package astiopendata_test

import (
	"net/http"
	"testing"

	"io/ioutil"

	"bytes"

	"github.com/asticode/go-astiratp/opendata"
	"github.com/stretchr/testify/assert"
)

func TestClient_Version(t *testing.T) {
	var c = astiopendata.New(astiopendata.Configuration{})
	var bodyRequest []byte
	var bodyResponse = []byte(`<?xml version='1.0' encoding='UTF-8'?><soapenv:Envelope xmlns:soapenv="http://www.w3.org/2003/05/soap-envelope"><soapenv:Body><ns2:getVersionResponse xmlns:ns2="http://wsiv.ratp.fr"><ns2:return>2.6.1 / 20170130</ns2:return></ns2:getVersionResponse></soapenv:Body></soapenv:Envelope>`)
	astiopendata.Send = func(req *http.Request, httpClient *http.Client) (resp *http.Response, err error) {
		bodyRequest, _ = ioutil.ReadAll(req.Body)
		return &http.Response{Body: ioutil.NopCloser(bytes.NewReader(bodyResponse)), StatusCode: http.StatusOK}, nil
	}
	v, err := c.Version()
	assert.NoError(t, err)
	assert.Equal(t, "<Envelope xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Body xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><getVersion xmlns=\"http://wsiv.ratp.fr\"></getVersion></Body></Envelope>", string(bodyRequest))
	assert.Equal(t, astiopendata.Version("2.6.1 / 20170130"), v)
}
