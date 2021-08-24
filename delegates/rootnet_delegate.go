package delegates

import (
	"encoding/json"

	"github.com/covid19/clients"
	"github.com/covid19/domain"
	"github.com/covid19/helpers"
)

func FetchRegionalCases() domain.CovidStatsResponse {
	var err error
	covidStatsJson, err := clients.FetchCovidStats()
	helpers.CheckErr(err)

	var covidStatsResponse domain.CovidStatsResponse
	err = json.Unmarshal([]byte(covidStatsJson), &covidStatsResponse)
	helpers.CheckErr(err)

	return covidStatsResponse
}
