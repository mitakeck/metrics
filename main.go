package main

import (
	"fmt"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	app    = kingpin.New("metrics", "A command-line metrics checker.")
	output = app.Flag("output", "specify output format (text, csv, json)").Default("text").Enum("text", "csv", "json")

	cpu     = app.Command("cpu", "show cpu infomation")
	memory  = app.Command("memory", "show memory infomation")
	disk    = app.Command("disk", "show disk infomation")
	network = app.Command("network", "show network information")
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
	case cpu.FullCommand():
		return getCPUMetrics()
	case memory.FullCommand():
		return getMemoryMetics()
	case network.FullCommand():
		return getNetworkMetics()
	case disk.FullCommand():
		return getDiskMetrics()
	}

	return nil, fmt.Errorf("metrics not found")
}
