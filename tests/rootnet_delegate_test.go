package tests

import (
	"testing"

	"github.com/covid19/delegates"
)

func TestRootNetDelegate(t *testing.T) {
	covidStatsResponse := delegates.FetchRegionalCases()

	if !covidStatsResponse.Success {
		t.Errorf("Unable to fetch covid stats. Terminating with response %v", covidStatsResponse)
	}
}
