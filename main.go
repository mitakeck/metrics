package main

import (
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("metrics", "A command-line metrics checker.")

	cpu     = app.Command("cpu", "show cpu infomation")
	memory  = app.Command("memory", "show memory infomation")
	disk    = app.Command("disk", "show disk infomation")
	network = app.Command("network", "show network information")
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case cpu.FullCommand():
		showCPUMetrics()
	case memory.FullCommand():
		showMemoryMetics()
	case network.FullCommand():
		showNetworkMetics()
	case disk.FullCommand():
		showDiskMetrics()
	}
}
