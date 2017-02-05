package astiopendata

import (
	"bytes"
	"compress/gzip"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

// Send allows testing functions using it
var Send = func(req *http.Request, httpClient *http.Client) (*http.Response, error) {
	return httpClient.Do(req)
}

// Send sends a request to the Open Data RATP API
func (c *Client) Send(action string, e Envelope) (resp *http.Response, err error) {
	// Marshal body
	var b []byte
	if b, err = xml.Marshal(e); err != nil {
		c.Logger.Error(err)
		return
	}

	// Create request
	var req *http.Request
	if req, err = http.NewRequest(http.MethodPost, endpointURL, bytes.NewReader(b)); err != nil {
		c.Logger.Error(err)
		return
	}
	defer req.Body.Close()

	// Add headers
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Content-Type", "application/soap+xml")
	req.Header.Add("SOAPAction", action)

	// Send request
	c.Logger.Debugf("Sending RATP %s request to %s with body %s and headers %+v", http.MethodPost, endpointURL, string(b), req.Header)
	if resp, err = Send(req, c.Client); err != nil {
		c.Logger.Error(fmt.Sprintf("%s for request to %s", err, req.URL))
	}
	return
}

// SendWithMaxRetries tries sending a request to the Open Data RATP API and retries in case of specific conditions
func (c *Client) SendWithMaxRetries(action string, e Envelope) (o Envelope, err error) {
	// We start at s.RetryMax + 1 so that it runs at least once even if RetryMax == 0
	var b []byte
	for retriesLeft := c.RetryMax + 1; retriesLeft > 0; retriesLeft-- {
		// Send request
		var resp, retry = &http.Response{}, false
		if resp, err = c.Send(action, e); err != nil {
			// If error is temporary, retry
			if netError, ok := err.(net.Error); ok && netError.Temporary() {
				retry = true
			} else {
				return
			}
		}

		// Read response body since it will be used in any case
		if resp != nil {
			// Make sure the body is closed
			defer resp.Body.Close()

			// Decompress the body
			var reader = resp.Body
			if resp.Header.Get("Content-Encoding") == "gzip" {
				if reader, err = gzip.NewReader(resp.Body); err != nil {
					c.Logger.Error(err)
					return
				}
				defer reader.Close()
			}

			// Read the body content
			if b, err = ioutil.ReadAll(reader); err != nil {
				c.Logger.Error(err)
				return
			}
		}

		// Retry if internal server or if error is temporary
		if retry || resp.StatusCode >= http.StatusInternalServerError {
			// Sleep
			if retriesLeft > 1 {
				c.Logger.Debugf("Sleeping %s and retrying... (%d retries left and body %s)", c.RetrySleep, retriesLeft-1, string(b))
				time.Sleep(c.RetrySleep)
			}
			continue
		}

		// Parse response if conditions for retrying were not met
		if resp.StatusCode < 200 || resp.StatusCode > 299 {
			err = fmt.Errorf("Invalid status code %v and body %s", resp.StatusCode, string(b))
			return
		}

		// Unmarshal
		o = Envelope{}
		if err = xml.Unmarshal(b, &o); err != nil {
			c.Logger.Error(err)
			return
		}
		return
	}

	// Max retries limit reached
	err = fmt.Errorf("Max retries %d reached for action %s with body %s", c.RetryMax, action, string(b))
	c.Logger.Error(err)
	return
}
