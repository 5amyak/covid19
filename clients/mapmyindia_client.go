package clients

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/covid19/helpers"
)

const mapMyIndiaBaseURL = "https://apis.mapmyindia.com/advancedmaps/v1"
const mapMyIndiaKey = "d1d7d0bbc7c20ca472a5582ca72caea3"

func FetchRevGeoCode(lat, lng string) []byte {
	const revGeoCodePath = "rev_geocode"

	url := fmt.Sprintf("%v/%v/%v/?lat=%v&lng=%v",
		mapMyIndiaBaseURL, mapMyIndiaKey, revGeoCodePath, lat, lng)
	resp, err := http.Get(url)
	helpers.CheckErr(err)

	body, err := ioutil.ReadAll(resp.Body)
	helpers.CheckErr(err)
	return body
}
