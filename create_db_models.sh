#!/usr/bin/env bash

DBS=("sqlite3" "psql" "mssql" "mysql")

for db in "${DBS[@]}"; do
    new_dir=models/$db
    mkdir -p $new_dir

    sqlboiler --output $new_dir $db
    xargs sed -i $new_dir/* 's/t.Parallel()//'
    go test -v $new_dir
done
