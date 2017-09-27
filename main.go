package main

import (
	"fmt"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	app    = kingpin.New("metrics", "A command-line metrics checker.")
	output = app.Flag("output", "specify output format (text, csv, json)").Default("text").Enum("text", "csv", "json")

	cpuOpt     = app.Command("cpu", "show cpu infomation")
	memoryOpt  = app.Command("memory", "show memory infomation")
	diskOpt    = app.Command("disk", "show disk infomation")
	networkOpt = app.Command("network", "show network information")
)

func main() {
	values, err := getMetrics(kingpin.MustParse(app.Parse(os.Args[1:])))
	if err != nil {
		fmt.Print(err)
		return
	}

	switch *output {
	case "text":
		outputText(values)
	case "csv":
		outputCsv(values)
	case "json":
		outputJSON(values)
	}
}

func getMetrics(comname string) (Values, error) {
	switch comname {
	case cpuOpt.FullCommand():
		return getCPUMetrics()
	case memoryOpt.FullCommand():
		return getMemoryMetics()
	case networkOpt.FullCommand():
		return getNetworkMetics()
	case diskOpt.FullCommand():
		return getDiskMetrics()
	}

	return nil, fmt.Errorf("metrics not found")
}
