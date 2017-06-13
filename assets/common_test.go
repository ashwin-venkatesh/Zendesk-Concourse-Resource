package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/cloudfoundry/garden-runc-release/src/github.com/gogo/protobuf/io"
	"strings"
)

func TestCanParseInput(t *testing.T) {
	input := `
{"version":
  {"ref": "1"}
}
`
	expected := "1"
	payload, err := parseInputJson(strings.NewReader(input))
	actual := payload.Version.Ref

	assertEqual(t, expected, actual)
	assert.Nil(t, err)
}

func TestCanSeedValue(t *testing.T) {
	input := []byte("{}")
	expected := "1"
	payload, err := parseInputJson(strings.NewReader(input))
	actual := payload.Version.Ref

	assertEqual(t, expected, actual)
	assert.Nil(t, err)
}

func assertEqual(t *testing.T, expected string, actual string) {
	if actual != expected {
		t.Errorf("expected %s but was %s", expected, actual)
	}
}
