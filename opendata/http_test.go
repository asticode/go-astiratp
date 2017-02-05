package astiopendata_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/asticode/go-astiratp/opendata"
	"github.com/rs/xlog"
	"github.com/stretchr/testify/assert"
)

func mockVersionResponse(i string) io.ReadCloser {
	return ioutil.NopCloser(strings.NewReader(fmt.Sprintf(`<soapenv:Envelope xmlns:soapenv="http://www.w3.org/2003/05/soap-envelope"><soapenv:Body><ns2:getVersionResponse xmlns:ns2="http://wsiv.ratp.fr"><ns2:return>%s</ns2:return></ns2:getVersionResponse></soapenv:Body></soapenv:Envelope>`, i)))
}

func TestSendWithMaxRetries(t *testing.T) {
	var count int
	d := astiopendata.Client{
		Logger:     xlog.NopLogger,
		RetrySleep: time.Nanosecond,
	}
	astiopendata.Send = func(req *http.Request, httpClient *http.Client) (*http.Response, error) {
		count++
		if count == 1 {
			return &http.Response{StatusCode: http.StatusInternalServerError, Body: mockVersionResponse("1")}, nil
		} else if count == 2 {
			return &http.Response{StatusCode: http.StatusGatewayTimeout, Body: mockVersionResponse("2")}, nil
		}
		return &http.Response{StatusCode: http.StatusOK, Body: mockVersionResponse("3")}, nil
	}
	d.RetryMax = 0
	var e astiopendata.Envelope
	e, err := d.SendWithMaxRetries("", astiopendata.Envelope{})
	assert.Error(t, err)
	count = 0
	d.RetryMax = 1
	e, err = d.SendWithMaxRetries("", astiopendata.Envelope{})
	assert.Error(t, err)
	count = 0
	d.RetryMax = 2
	e, err = d.SendWithMaxRetries("", astiopendata.Envelope{})
	assert.NoError(t, err)
	assert.Equal(t, astiopendata.Version("3"), e.Body.GetVersionResponse.Version)
}
