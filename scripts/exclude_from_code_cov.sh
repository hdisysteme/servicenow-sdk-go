#!/bin/sh

# Assuming relative paths in exclude-from-code-coverage.txt
# Prepend "github.com/hdisysteme/servicenow-sdk-go" to each path
basePath="github.com/hdisysteme/servicenow-sdk-go"

while read p || [ -n "$p" ]
do
    full_path="$basePath/$p"
    sed -i '' "/${full_path//\//\\/}/d" ./coverage.out
done < ./exclude-from-code-coverage.txt
