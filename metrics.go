package main

import (
	"strconv"
	"strings"

	pipeline "github.com/mattn/go-pipeline"
)

// Values はメトリクス名とその数値の kv map
type Values map[string]float64

func getCPUMetrics() (Values, error) {
	iostatLabel := []string{"cpu.user", "cpu.system", "cpu.idol"}
	rawOutput, err := pipeline.Output(
		[]string{"iostat"},
		[]string{"sed", "-n", "3P"},
		[]string{"awk", "{print $4 \" \" $5 \" \" $6}"},
	)
	if err != nil {
		return nil, err
	}

	output := string(rawOutput)
	fields := strings.Fields(output)
	ret := make(map[string]float64, len(iostatLabel))
	for i := range iostatLabel {
		value, err := strconv.ParseFloat(fields[i], 64)
		if err != nil {
			return nil, err
		}

		ret[iostatLabel[i]] = value
	}

	return ret, nil
}

func getMemoryMetics() (Values, error) {
	vmstatLabel := []string{"memory.free", "memory.active", "memory.inactive", "memory.total"}
	rawOutput, err := pipeline.Output(
		[]string{"vm_stat", "-c", "1", "1"},
		[]string{"sed", "-n", "3P"},
		[]string{"awk", "{print $1*4096 \" \" $2*4096 \" \" $3*4096 \" \" ($1+$2+$3)*4096}"},
	)
	if err != nil {
		return nil, err
	}

	output := string(rawOutput)
	fields := strings.Fields(output)
	ret := make(map[string]float64, len(vmstatLabel))
	for i := range vmstatLabel {
		value, err := strconv.ParseFloat(fields[i], 64)
		if err != nil {
			return nil, err
		}

		ret[vmstatLabel[i]] = value
	}

	return ret, nil
}

func getNetworkMetics() (Values, error) {
	netstatLabel := []string{"network.ibyte", "network.obyte"}
	rawOutput, err := pipeline.Output(
		[]string{"netstat", "-inb"},
		[]string{"grep", "en0"},
		[]string{"head", "-n", "1"},
		[]string{"awk", "{print $7 \" \" $10}"},
	)
	if err != nil {
		return nil, err
	}

	output := string(rawOutput)
	fields := strings.Fields(output)
	ret := make(map[string]float64, len(netstatLabel))
	for i := range netstatLabel {
		value, err := strconv.ParseFloat(fields[i], 64)
		if err != nil {
			return nil, err
		}

		ret[netstatLabel[i]] = value
	}

	return ret, nil
}

func getDiskMetrics() (Values, error) {
	dfLabel := []string{"disk.total", "disk.used", "disk.free"}
	rawOutput, err := pipeline.Output(
		[]string{"df"},
		[]string{"grep", "/dev/disk1"},
		[]string{"head", "-n", "1"},
		[]string{"awk", "{print $2 \" \" $3 \" \" $4}"},
	)
	if err != nil {
		return nil, err
	}

	output := string(rawOutput)
	fields := strings.Fields(output)
	ret := make(map[string]float64, len(dfLabel))
	for i := range dfLabel {
		value, err := strconv.ParseFloat(fields[i], 64)
		if err != nil {
			return nil, err
		}

		ret[dfLabel[i]] = value
	}

	return ret, nil
}
