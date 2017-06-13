package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/cloudfoundry/garden-runc-release/src/github.com/gogo/protobuf/io"
	"strings"
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
	findById("onboardsupport", "pablodepacas@hotmail.com", "124816")
}

func assertEqual(t *testing.T, expected string, actual string) {
	if actual != expected {
		t.Errorf("expected %s but was %s", expected, actual)
	}
}
