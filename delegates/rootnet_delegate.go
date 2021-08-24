package delegates

import (
	"encoding/json"
	"errors"

	"github.com/covid19/clients"
	"github.com/covid19/domain"
	"github.com/covid19/helpers"
)

func FetchRegionalCases() domain.CovidStatsResponse {
	covidStatsJson, err := clients.FetchCovidStats()
	helpers.CheckErr(err)

	var covidStatsResponse domain.CovidStatsResponse
	err = json.Unmarshal([]byte(covidStatsJson), &covidStatsResponse)
	helpers.CheckErr(err)

	if !covidStatsResponse.Success {
		helpers.CheckErr(errors.New("unable to fetch covid stats"))
	}

	return covidStatsResponse
}
