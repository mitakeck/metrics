package main

import (
	"encoding/json"
	"fmt"
)

func outputText(metrics Values) {
	for name := range metrics {
		fmt.Printf("%s\t%f\n", name, metrics[name])
	}
}

func outputCsv(metrics Values) {
	for name := range metrics {
		fmt.Printf("%s, %f\n", name, metrics[name])
	}
}

func outputJSON(metrics Values) {
	jsonBytes, err := json.Marshal(metrics)
	if err != nil {
		fmt.Println("json marshal fail.")
		return
	}
	fmt.Println(string(jsonBytes))
}
