package clients

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const mapMyIndiaBaseURL = "https://apis.mapmyindia.com/advancedmaps/v1"
const mapMyIndiaKey = "d1d7d0bbc7c20ca472a5582ca72caea3"

func FetchRevGeoCode(lat, lng string) ([]byte, error) {
	const revGeoCodePath = "rev_geocode"

	url := fmt.Sprintf("%v/%v/%v/?lat=%v&lng=%v",
		mapMyIndiaBaseURL, mapMyIndiaKey, revGeoCodePath, lat, lng)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}
