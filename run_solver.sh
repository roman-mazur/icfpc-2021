#!/bin/bash

while [ true ]
do
  for i in problems/problem.*
  do
    go build ./cmd/solver
    ./solver -as-service -iterations $(( $RANDOM % 10000 )) -gen-size $(( $RANDOM % 2048 )) $i
  done
done
