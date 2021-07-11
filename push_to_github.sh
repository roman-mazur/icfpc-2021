#!/bin/bash

git config difftool.prompt false
git config diff.tool diffmerge
git difftool

while [[ true ]]
do
  git pull
  git add solutions
  git commit -m "Add new batch of solutions"
  git push
  sleep 360
done
