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
	parseInputJson(inbytes)
}
