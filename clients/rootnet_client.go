package clients

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const rootNetBaseURL = "https://api.rootnet.in"

func FetchCovidStats() ([]byte, error) {
	const covidStatsPath = "covid19-in/stats/latest"

	resp, err := http.Get(fmt.Sprintf("%v/%v", rootNetBaseURL, covidStatsPath))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}
