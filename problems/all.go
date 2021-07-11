package main

import (
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
)

func main() {
	entries, err := ioutil.ReadDir("./problems")
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		if entry.Name() == "all.go" {
			continue
		}
		path := filepath.Join("problems", entry.Name())
		solver := exec.Command("go", "run", "./cmd/solver", "--as-service", path)
		log.Printf("Solving %s", path)
		err := solver.Run()
		if err != nil {
			log.Printf("Problem with %s: %s", path, err)
		}
	}
}
