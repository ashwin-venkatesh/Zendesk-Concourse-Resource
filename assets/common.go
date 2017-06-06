package common

import (
	"encoding/json"
	"fmt"
	"os"
)

func parseInputJson(inbytes []byte) *Payload {
	payload := &Payload{}

	if err := json.Unmarshal(inbytes, payload); err != nil {
		fmt.Fprintf(os.Stderr, "error parsing input as JSON: %s", err)
		os.Exit(1)
	}
	return payload
}
