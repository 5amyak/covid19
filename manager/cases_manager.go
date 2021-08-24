package manager

import (
	"fmt"
	"net/http"

	"github.com/covid19/delegates"
	"github.com/covid19/repository"
	"github.com/labstack/echo/v4"
)

func UpdateCaseCount(c echo.Context) error {
	defer handleServerError(c)
	covidStatsResponse := delegates.FetchRegionalCases()
	for _, v := range covidStatsResponse.Data.RegionalCases {
		repository.UpdateCaseCount(&v, covidStatsResponse.LastRefreshed)
	}

	return c.String(http.StatusOK,
		fmt.Sprintf("Data of %v states and union territories updated!", len(covidStatsResponse.Data.RegionalCases)))
}

func GetCaseCount(c echo.Context) error {
	defer handleServerError(c)
	geoCodeInfo := delegates.FetchRevGeoCode(c.QueryParam("lat"), c.QueryParam("lng"))
	stateCaseCount := repository.FetchCaseCount(geoCodeInfo.State)
	stateCaseCount["active"] = stateCaseCount["confirmed"].(int32) - stateCaseCount["deaths"].(int32) - stateCaseCount["discharged"].(int32)

	aggCaseCount := repository.FetchAggregatedCaseCount()
	aggCaseCount["totalActive"] = aggCaseCount["totalConfirmed"].(int32) - aggCaseCount["totalDeaths"].(int32) - aggCaseCount["totalDischarged"].(int32)
	lastRefreshedOn := stateCaseCount["lastRefreshedOn"]

	delete(aggCaseCount, "_id")
	delete(stateCaseCount, "_id")
	delete(stateCaseCount, "lastRefreshedOn")
	delete(stateCaseCount, "state")

	caseCountResponse := make(map[string]interface{})
	caseCountResponse["India"] = aggCaseCount
	caseCountResponse[geoCodeInfo.State] = stateCaseCount
	caseCountResponse["lastUpdatedOn"] = lastRefreshedOn

	return c.JSON(http.StatusOK, caseCountResponse)
}

func handleServerError(c echo.Context) error {
	if r := recover(); r != nil {
		if err, ok := r.(error); ok {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, err.Error())
		}
	}

	return nil
}
