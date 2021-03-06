package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const token = "5018e0e7-6d55-4647-966f-a67fefacbdd4"

var ptrn = regexp.MustCompile("problems_problem.(\\d+)-score-(.+)(_local)?.json")

var doIt = flag.Bool("do-it", false, "Actually do the HTTP calls")

var submitted = make(map[int]int)

func main()  {
	flag.Parse()

	if !*doIt {
		*doIt = strings.HasSuffix(os.Getenv("GITHUB_REF"), "/main")
		if *doIt {
			log.Println("Setting --do-it to true for the GitHub main branch")
		}
	}

	submittedFp := filepath.Join("solutions", "submitted.json")
	if f, err := os.Open(submittedFp); err == nil {
		err = json.NewDecoder(f).Decode(&submitted)
		if err != nil {
			log.Fatal("Cannot decode submitted.json", err)
		}
		f.Close()
		log.Println(submitted)
	}

	entries, err := ioutil.ReadDir("./solutions")
	if err != nil {
		log.Fatal(err)
	}
	solutions := make([]string, 500)
	scores := make([]float64, 500)
	for i := range scores {
		scores[i] = math.Inf(1)
	}

	for _, entry := range entries {
		parsed := ptrn.FindStringSubmatch(entry.Name())
		if len(parsed) < 3 {
			fmt.Println("Cannot parse", entry.Name())
			continue
		}

		number, err := strconv.Atoi(parsed[1])
		if err != nil {
			fmt.Println("Cannot parse problem number for", entry.Name())
			continue
		}
		score, err := strconv.ParseFloat(parsed[2], 64)
		if err != nil {
			fmt.Println("Cannot parse score for", entry.Name())
			continue
		}
		if score < scores[number] {
			solutions[number] = entry.Name()
			scores[number] = score
		}
	}

	ctx := context.Background()
	for number, solution := range solutions {
		if solution != "" {
			score := int(scores[number])
			if score == submitted[number] {
				log.Printf("No better solutions for %d", number)
				continue
			}

			file, err := os.Open(filepath.Join("solutions", solution))
			if err != nil {
				log.Fatal("cannot read", solution)
			}
			reqCtx, cancelFunc := context.WithTimeout(ctx, 5*time.Second)
			req, err := http.NewRequestWithContext(reqCtx, "POST", fmt.Sprintf("https://poses.live/api/problems/%d/solutions", number), file)
			if err != nil {
				log.Fatal(err)
			}
			req.Header.Add("authorization", "Bearer " + token)

			log.Printf("Submitting for %d with score %d", number, score)
			if *doIt {
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					log.Printf("Error submitting %d: %s", number, err)
				} else {
					log.Printf("Status for %d: %d", number, resp.StatusCode)
					if resp.StatusCode == 200 {
						submitted[number] = score
					}
				}
			} else {
				submitted[number] = score
			}

			file.Close()
			cancelFunc()

			if *doIt {
				// Don't attach the contest server.
				time.Sleep(200)
			}
		}
	}

	output, err := os.Create(submittedFp)
	if err != nil {
		log.Fatal("Cannot store submissions", err)
	}
	defer output.Close()
	enc := json.NewEncoder(output)
	enc.SetIndent("", "  ")
	err = enc.Encode(submitted)
	if err != nil {
		log.Fatal("Cannot serialize submissions", err)
	}
}
