package pixoo64

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/netip"
)

// Client is used to send and receive the data to a pixoo64  device.
type Client struct {
	addr       netip.Addr
	httpClient *http.Client
}

// Post sends data to the pixoo64 device.
func (c *Client) Post(data []byte) (string, error) {
	request, err := http.NewRequest("POST", "http://"+c.addr.String()+"/post", bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	response, err := c.httpClient.Do(request)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = response.Body.Close()
	}()

	if response.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			return "", err
		}
		bodyString := string(bodyBytes)
		return bodyString, nil
	} else {
		return "", fmt.Errorf("unsupported response status code: %d", response.StatusCode)
	}
}
