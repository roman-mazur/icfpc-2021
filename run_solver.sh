#!/bin/bash

function shuffle() {
  while IFS= read -r line
  do
    printf "%06d %s\n" $RANDOM "$line"
  done | sort -n | cut -c8-

}

while [ true ]
do
  for i in $(ls problems/problem.* | shuffle)
  do
    go build ./cmd/solver

    echo "Solver against $i"
    ./solver -as-service -iterations $(( $RANDOM % 5000 )) -gen-size $(( $RANDOM % 256 )) $i
  done
done
