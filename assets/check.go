package main

import (
	"io/ioutil"
	"os"
)

type Payload struct {
}

func main() {
	inbytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	common.parseInputJson(inbytes)
}
