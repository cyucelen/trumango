#!/bin/sh
output=$(go test $(go list ./...) -v -coverprofile cov.txt -timeout 10s -short)
sum=0
counter=0

while read -r line; do
    echo $line
    if [[ $line =~ ^coverage: ]]
    then
        percentage=`echo $line | grep -Eo '[0-9]+([.][0-9]+)?'`
        sum=`awk "BEGIN {print $percentage+$sum}"`
        ((counter++))
    fi
done <<< "$output"

average=`awk "BEGIN {print $sum/$counter}"`
echo "coverage: $average% of statements"
