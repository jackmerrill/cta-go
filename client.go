package ctago

import "net/url"

type Client struct {
	http *NetworkClient

	Arrivals *ArrivalsService
}

func NewClient(apiKey string) *Client {
	base, err := url.Parse("https://lapi.transitchicago.com/api/1.0")

	if err != nil {
		panic(err)
	}

	c := &Client{
		http: &NetworkClient{
			baseURL: *base,
			apiKey:  apiKey,
		},
	}

	c.Arrivals = &ArrivalsService{
		http: c.http,
	}

	return c
}
