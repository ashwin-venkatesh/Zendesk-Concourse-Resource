package main

import (
	"testing"
	"fmt"
)

func TestCanParseInput(t *testing.T) {
	input := []byte(`
{"version":
  {"ref": "1"}
}
`)
	expected := "1"
	payload := parseInputJson(input)
	actual := payload.Version.Ref

	assertEqual(t, expected, actual)
}

func TestCanSeedValue(t *testing.T) {
	input := []byte("{}")
	expected := "1"
	payload := parseInputJson(input)
	actual := payload.Version.Ref

	assertEqual(t, expected, actual)
}

func TestCanFetchId(t *testing.T) {
	findById("onboardsupport", "pablodepacas@hotmail.com", "124816", "1")
}

func TestParseZendDeskFindByIdJson(t *testing.T) {
	expectedDate:="expected_date"
	json := `
{
	"results": [
		{
			"created_at": "%s"
		}
	],
	"count": 1
}
`
	actual := parseZendDeskFindByIdJson(fmt.Sprintf(json, expectedDate))
	actualDate := actual.Results[0].Created_At

	assertEqual(t, expectedDate, actualDate)
}

func assertEqual(t *testing.T, expected string, actual string) {
	if actual != expected {
		t.Errorf("expected %s but was %s", expected, actual)
	}
}
