package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Source struct {
	Uri        string
	Branch     string
	PrivateKey string
}

type Version struct {
	Ref string
}

type Payload struct {
	Source  Source
	Version Version
}

func parseInputJson(inbytes []byte) *Payload {
	defaultVersion := &Version{Ref: "1"}
	payload := &Payload{Version: *defaultVersion}

	if err := json.Unmarshal(inbytes, payload); err != nil {
		fmt.Fprintf(os.Stderr, "error parsing input as JSON: %s", err)
		os.Exit(1)
	}
	return payload
}
