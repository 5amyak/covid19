package delegates

import (
	"encoding/json"

	"github.com/covid19/clients"
	"github.com/covid19/domain"
)

func FetchRevGeoCode(lat, lng string) domain.Result {
	revGeoCodeJSON, _ := clients.FetchRevGeoCode(lat, lng)

	var revGeoCodeResponse domain.RevGeoCodeResponse
	json.Unmarshal([]byte(revGeoCodeJSON), &revGeoCodeResponse)

	return revGeoCodeResponse.Results[0]
}
