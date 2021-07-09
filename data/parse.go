package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func ParseProblem(file string) *Problem {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Could not read: %s\n", file)
	}

	var pb Problem
	if err := json.Unmarshal(data, &pb); err != nil {
		log.Fatalf("Invalid file format (%s)\n", err.Error())
	}

	return &pb
}
