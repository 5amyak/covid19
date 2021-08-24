package controllers

import (
	"net/http"

	"github.com/covid19/delegates"
	"github.com/covid19/repository"
	"github.com/labstack/echo/v4"
)

func UpdateCaseCount(c echo.Context) error {
	covidStatsResponse := delegates.FetchRegionalCases()
	for _, v := range covidStatsResponse.Data.RegionalCases {
		repository.UpdateCaseCount(&v, covidStatsResponse.LastRefreshed)
	}

	return c.String(http.StatusOK, "DB Updated!")
}

func GetCaseCount(c echo.Context) error {
	geoCodeInfo := delegates.FetchRevGeoCode(c.QueryParam("lat"), c.QueryParam("lng"))
	stateCaseCount := repository.FetchCaseCount(geoCodeInfo.State)
	stateCaseCount["active"] = stateCaseCount["confirmed"].(int32) - stateCaseCount["deaths"].(int32) - stateCaseCount["discharged"].(int32)

	aggCaseCount := repository.FetchAggregatedCaseCount()
	aggCaseCount["totalActive"] = aggCaseCount["totalConfirmed"].(int32) - aggCaseCount["totalDeaths"].(int32) - aggCaseCount["totalDischarged"].(int32)
	lastRefreshedOn := stateCaseCount["lastRefreshedOn"]

	delete(aggCaseCount, "_id")
	delete(stateCaseCount, "_id")
	delete(stateCaseCount, "lastRefreshedOn")

	caseCountResponse := make(map[string]interface{})
	caseCountResponse["India"] = aggCaseCount
	caseCountResponse[stateCaseCount["state"].(string)] = stateCaseCount
	caseCountResponse["lastRefreshedOn"] = lastRefreshedOn

	return c.JSON(http.StatusOK, caseCountResponse)
}
