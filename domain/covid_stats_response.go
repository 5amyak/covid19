package domain

import "time"

type CovidStatsResponse struct {
	Success       bool           `json:"success"`
	Data          CovidStatsData `json:"data"`
	LastRefreshed time.Time      `json:"lastRefreshed"`
}

type CovidStatsData struct {
	RegionalCases []RegionalCase `json:"regional"`
}

type RegionalCase struct {
	State      string `json:"loc"`
	Confirmed  int    `json:"totalConfirmed"`
	Discharged int    `json:"discharged"`
	Deaths     int    `json:"deaths"`
}
