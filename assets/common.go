package main

import (
	"encoding/json"
	"fmt"
	"os"
	"net/http"
	"bytes"
)

type Payload struct {
	Source  struct {
		Uri        string
		Branch     string
		PrivateKey string
	}
	Version struct {
		Ref string
	}
}

func parseInputJson(inbytes []byte) *Payload {
	payload := &Payload{Version: struct {Ref string}{Ref: "1"}}

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

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	return buf.String()
}

type ZendeskResponse struct {
	Results []struct {
		Created_At string
	}
	Count int
}

func parseZendDeskFindByIdJson(jsonString string) *ZendeskResponse{
	response:= &ZendeskResponse{}

	if err := json.Unmarshal([]byte(jsonString), response); err != nil {
		fmt.Fprintf(os.Stderr, "error parsing response as JSON: %s", err)
		os.Exit(1)
	}

	return response
}

func getLastDateById(subdomain string, userName string, password string, id string) string {
	jsonResponse := findById(subdomain, userName, password, id)
	zendeskResponse := parseZendDeskFindByIdJson(jsonResponse)
	date := zendeskResponse.Results[0].Created_At
	return date
}
