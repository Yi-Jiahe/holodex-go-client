package holodex

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const APIBaseURL = "https://holodex.net/api/v2"

type Client struct {
	APIKey string
	Client *http.Client
}

func NewClient(APIKey string) *Client {
	return &Client{
		APIKey: APIKey,
		Client: &http.Client{},
	}
}

func (c *Client) makeAuthenticatedRequest(req *http.Request) (*http.Response, error) {
	req.Header.Add("X-APIKEY", c.APIKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, nil
	}

	return resp, nil
}

func (c *Client) QueryLiveAndUpcoming(channelIds []string) ([]VideoWithChannel, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/users/live", APIBaseURL), nil)
	if err != nil {
		return []VideoWithChannel{}, err
	}

	q := req.URL.Query()
	q.Add("channels", strings.Join(channelIds, ","))
	req.URL.RawQuery = q.Encode()

	resp, err := c.makeAuthenticatedRequest(req)
	if err != nil {
		return []VideoWithChannel{}, err
	}

	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return []VideoWithChannel{}, err
	}

	var videos []VideoWithChannel
	err = json.Unmarshal(respBodyBytes, &videos)
	if err != nil {
		return []VideoWithChannel{}, err
	}

	return videos, nil
}
