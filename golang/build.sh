#!/bin/bash

for project in $(find * -type d) ; do
  cd "$project" || exit
  echo "Building Go project: $project"
  go build -o ../"$project"_bin .
  echo "Build for $project is complete"
  cd .. || exit
done
