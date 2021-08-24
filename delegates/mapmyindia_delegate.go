package delegates

import (
	"encoding/json"
	"errors"

	"github.com/covid19/clients"
	"github.com/covid19/domain"
	"github.com/covid19/helpers"
)

func FetchRevGeoCode(lat, lng string) domain.Result {
	revGeoCodeJSON, err := clients.FetchRevGeoCode(lat, lng)
	helpers.CheckErr(err)

	var revGeoCodeResponse domain.RevGeoCodeResponse
	err = json.Unmarshal([]byte(revGeoCodeJSON), &revGeoCodeResponse)
	helpers.CheckErr(err)

	if revGeoCodeResponse.ResponseCode != 200 || len(revGeoCodeResponse.Results) == 0 {
		panic(errors.New("unable to find state from geocode"))
	}

	return revGeoCodeResponse.Results[0]
}
