package delegates

import (
	"encoding/json"

	"github.com/covid19/clients"
	"github.com/covid19/domain"
)

func FetchRegionalCases() domain.CovidStatsResponse {
	covidStatsJson, _ := clients.FetchCovidStats()

	var covidStatsResponse domain.CovidStatsResponse
	json.Unmarshal([]byte(covidStatsJson), &covidStatsResponse)

	return covidStatsResponse
}
