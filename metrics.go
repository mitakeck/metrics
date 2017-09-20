package main

import (
	"log"
	"strings"

	"github.com/k0kubun/pp"
	pipeline "github.com/mattn/go-pipeline"
)

var iostatLabel = []string{"user", "system", "idle"}

func showCPUMetrics() {
	rawOutput, err := pipeline.Output(
		[]string{"iostat"},
		[]string{"sed", "-n", "3P"},
		[]string{"awk", "{print $4 \" \" $5 \" \" $6}"},
	)
	if err != nil {
		log.Fatal(err)
	}

	output := string(rawOutput)
	fields := strings.Fields(output)
	pp.Print(fields)
}

func showMemoryMetics() {
	rawOutput, err := pipeline.Output(
		[]string{"vm_stat", "-c", "1", "1"},
		[]string{"sed", "-n", "3P"},
		[]string{"awk", "{print $1*4096 \" \" $2*4096 \" \" $3*4096 \" \" ($1+$2+$3)*4096}"},
	)
	if err != nil {
		log.Fatal(err)
	}

	output := string(rawOutput)
	fields := strings.Fields(output)
	pp.Print(fields)
}

func showNetworkMetics() {
	rawOutput, err := pipeline.Output(
		[]string{"netstat", "-inb"},
		[]string{"grep", "en0"},
		[]string{"head", "-n", "1"},
		[]string{"awk", "{print $7 \" \" $10}"},
	)
	if err != nil {
		log.Fatal(err)
	}

	output := string(rawOutput)
	fields := strings.Fields(output)
	pp.Print(fields)
}

func showDiskMetrics() {
	rawOutput, err := pipeline.Output(
		[]string{"df"},
		[]string{"grep", "/dev/disk1"},
		[]string{"head", "-n", "1"},
		[]string{"awk", "{print $2 \" \" $3 \" \" $4}"},
	)
	if err != nil {
		log.Fatal(err)
	}

	output := string(rawOutput)
	fields := strings.Fields(output)
	pp.Print(fields)
}
