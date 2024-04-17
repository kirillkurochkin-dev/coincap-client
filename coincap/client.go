package coincap

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Client struct {
	client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("Timeout can't be zero")
	}

	return &Client{
		client: &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				fmt.Println(req.Response.Status)
				fmt.Println("REDIRECT")
				return nil
			},

			Transport: &loggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},

			Timeout: timeout,
		},
	}, nil
}

func (c Client) GetAssets() ([]AssetData, error) {
	resp, err := c.client.Get("https://api.coincap.io/v2/assets")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r assetsResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return r.Data, nil
}

func (c Client) GetAsset(name string) (AssetData, error) {
	resp, err := c.client.Get("https://api.coincap.io/v2/assets/" + name)
	if err != nil {
		return AssetData{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return AssetData{}, err
	}

	var r assetResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return AssetData{}, err
	}

	return r.Data, nil
}
