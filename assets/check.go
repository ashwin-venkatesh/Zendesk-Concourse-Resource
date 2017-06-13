package main

import (
	"io/ioutil"
	"os"
)

func main() {
	inbytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	payload := parseInputJson(inbytes)
	_ = findById("onboardsupport", "pablodepacas@hotmail.com", "124816", payload.Version.Ref)
}
