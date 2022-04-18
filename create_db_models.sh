#!/usr/bin/env bash

DBS=("sqlite3" "psql" "mssql" "mysql")

for db in "${DBS[@]}"; do
    new_dir=models/$db
    mkdir -p $new_dir

    sqlboiler --output $new_dir $db
    xargs sed -i 's/t.Parallel()//' $new_dir/*
    go test -v $new_dir
done
