package pkge

import (
	"github.com/google/go-querystring/query"
)

const baseURL = "https://api.pkge.net/v1"

type Client struct {
	httpClient *HttpClient
}

type PackageParams struct {
	TrackNumber string `url:"trackNumber"`
	CourierID   int    `url:"courierId,omitempty"`
}

func New(apiKey string) *Client {
	return &Client{
		httpClient: NewHttpClient(baseURL, apiKey),
	}
}

// Request to receive activated delivery services
func (c *Client) ActivatedDeliveryServices() ([]Courier, error) {
	var data CouriersResponse
	err := c.httpClient.Do("GET", "/couriers/enabled", nil, &data)
	if err != nil {
		return nil, err
	}
	return data.Couriers, nil
}

// Request to receive all available delivery services
func (c *Client) AllDeliveryServices() ([]Courier, error) {
	var data CouriersResponse
	err := c.httpClient.Do("GET", "/couriers/", nil, &data)
	if err != nil {
		return nil, err
	}
	return data.Couriers, nil
}

// Definition of delivery service
func (c *Client) DefinitionDeliveryService(trackNumber string) (*CouriersResponse, error) {
	var data CouriersResponse
	opt := PackageParams{TrackNumber: trackNumber}
	params, _ := query.Values(opt)
	err := c.httpClient.Do("GET", "/couriers/detect", params, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// Adding a package
func (c *Client) AddPackage(trackNumber string, courierID int) (*PackageResponse, error) {
	var data PackageResponse
	opt := PackageParams{TrackNumber: trackNumber, CourierID: courierID}
	params, _ := query.Values(opt)
	err := c.httpClient.Do("POST", "/packages", params, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// Updating a package
func (c *Client) UpdatePackage(trackNumber string) (*Response, error) {
	var data Response
	opt := PackageParams{TrackNumber: trackNumber}
	params, _ := query.Values(opt)
	err := c.httpClient.Do("POST", "/packages/update", params, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// package information
func (c *Client) PackageInfo(trackNumber string) (*Package, error) {
	var data PackageResponse
	opt := PackageParams{TrackNumber: trackNumber}
	params, _ := query.Values(opt)
	err := c.httpClient.Do("GET", "/packages", params, &data)
	if err != nil {
		return nil, err
	}
	return &data.Payload, nil
}

// Modifying package information
func (c *Client) ModifyPackageInfo(trackNumber string) (*Package, error) {
	var data PackageResponse
	opt := PackageParams{TrackNumber: trackNumber}
	params, _ := query.Values(opt)
	err := c.httpClient.Do("PUT", "/packages?", params, &data)
	if err != nil {
		return nil, err
	}
	return &data.Payload, nil
}

// Deleting a package
func (c *Client) DeletePackage(trackNumber string) (*Response, error) {
	var data Response
	opt := PackageParams{TrackNumber: trackNumber}
	params, _ := query.Values(opt)
	err := c.httpClient.Do("DELETE", "/packages", params, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// list of packages
func (c *Client) Packages() ([]Package, error) {
	var data PackageList
	err := c.httpClient.Do("GET", "/packages/list", nil, &data)
	if err != nil {
		return nil, err
	}
	return data.Payload, nil
}
