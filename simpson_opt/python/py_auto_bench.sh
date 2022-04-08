#!/bin/bash
for (( i=1; i < 100; i++ ))
do
        echo "$i iteration "
        python test.py >> py_res_bench.csv
done

