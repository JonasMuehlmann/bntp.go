#!/usr/bin/env bash

DBS=("sqlite3" "psql" "mssql" "mysql")

for db in "${DBS[@]}"; do
    new_dir=models/$db
    mkdir $new_dir
    sqlboiler $db
    ls -1 $new_dir | xargs sed -i 's/t.Parallel()//'
    go test -v $new_dir
done
