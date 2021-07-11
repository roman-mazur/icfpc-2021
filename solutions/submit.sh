#!/bin/bash

API_TOKEN=5018e0e7-6d55-4647-966f-a67fefacbdd4

while [[ $# -gt 0 ]]
do
  case $1 in
    -f|--filename)
      FILENAME=$2
      if [[ "$FILENAME" == "" ]]
      then
        echo "Missing argument to -f"
        exit 2
      fi
      shift 2
      ;;
    -p|--problem)
      PROBLEM_ID=$2
      if [[ "$PROBLEM_ID" == "" ]]
      then
        echo "Missing argument to -p"
        exit 2
      fi
      shift 2
      ;;
    *)
      echo "Usage: $0 -p PROBLEM_ID -f ./path/to/solution.json"
  esac
done

if [[ "$FILENAME" == "" ]]
then
  FILENAME=$(ls *problem.$PROBLEM_ID* | head -n 1)
fi

if [[ "$FILENAME" == "" ]] || [[ "$PROBLEM_ID" == "" ]]
then
  echo "Missing parameter"
  exit 2
fi

curl https://poses.live/api/problems/$PROBLEM_ID/solutions -H "Authorization: Bearer $API_TOKEN" --data "@$FILENAME"
