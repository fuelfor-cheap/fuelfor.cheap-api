package main

// ProjectZeroThreeAPIResponse is the top level response from the api
type ProjectZeroThreeAPIResponse struct {
	Updated int64                                 `json:"updated"`
	Regions []*ProjectZeroThreeAPIResponseRegions `json:"regions"`
}

// ProjectZeroThreeAPIResponseRegions contains all the states
type ProjectZeroThreeAPIResponseRegions struct {
	Region string                                     `json:"region"`
	Prices []*ProjectZeroThreeAPIResponsePriceRegions `json:"prices"`
}

// ProjectZeroThreeAPIResponsePriceRegions details the fuel price and region information
type ProjectZeroThreeAPIResponsePriceRegions struct {
	Type      string  `json:"type"`
	Price     float64 `json:"price"`
	Name      string  `json:"name"`
	State     string  `json:"state"`
	Postcode  string  `json:"postcode"`
	Suburb    string  `json:"suburb"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}
