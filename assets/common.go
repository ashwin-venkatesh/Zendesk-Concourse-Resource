package main

import (
	"encoding/json"
	"fmt"
	"os"
	"net/http"
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

func findById(subdomain string, userName string, password string, id string) string {
	url := fmt.Sprintf("https://%s.zendesk.com/api/v2/search.json?query=%s", subdomain, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating request: %s", err)
		os.Exit(1)
	}
	req.SetBasicAuth(userName, password)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error making request: %s", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	return string(resp.Body)
}
