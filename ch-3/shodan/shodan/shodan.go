package shodan

// BaseURL for the api calls
const BaseURL = "https://api.shodan.io"

// Client stores data for a Shodan client
type Client struct {
	apiKey string
}

// New Creates a new Shodan client
func New(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}
