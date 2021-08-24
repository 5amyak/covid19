package domain

type RevGeoCodeResponse struct {
	ResponseCode int      `json:"responseCode"`
	Results      []Result `json:"results"`
}

type Result struct {
	State string `json:"state"`
	Area  string `json:"area"`
}
