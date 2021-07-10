package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/roman-mazur/icfpc-2021/data"
)

func WriteSolution(sol data.Solution, number int) {
	if solFile, err := os.Create(filepath.Join("solutions", fmt.Sprintf("%d.json", number))); err == nil {
		defer solFile.Close()
		if err := json.NewEncoder(solFile).Encode(sol); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
}
