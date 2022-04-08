#!/bin/bash
for (( i=1; i < 100; i++ ))
do
        echo "$i iteration "
        go test -bench=. >> go_res_bench
done

