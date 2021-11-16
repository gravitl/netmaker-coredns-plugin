package netmaker

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	APIURL     string
	APIToken   string
	HttpClient *http.Client
}

func NewClient(APIURL string, APIToken string) *Client {
	return &Client{
		APIURL:     APIURL,
		APIToken:   APIToken,
		HttpClient: http.DefaultClient,
	}
}

func (nm Client) DNS(ctx context.Context) ([]DNSEntry, error) {
	url := fmt.Sprintf("%v/api/dns", nm.APIURL)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", nm.APIToken))

	resp, err := nm.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	entries := make([]DNSEntry, 0)
	err = json.NewDecoder(resp.Body).Decode(&entries)
	if err != nil {
		return nil, err
	}

	return entries, nil
}
