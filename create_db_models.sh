#!/usr/bin/env bash

go install github.com/volatiletech/sqlboiler/v4@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mssql@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-sqlite3@latest

DBS=("sqlite3" "mysql" "psql" "mssql")

for db in "${DBS[@]}"; do
    new_dir=./models/$db
    mkdir -p $new_dir

    sqlboiler --output $new_dir $db
    sed -i 's/t.Parallel()//' $new_dir/*
    if [[ $db == "sqlite3" ]]; then
        cp bntp_test.db $new_dir
    fi
    if [[ $db == "mysql" ]]; then
        sed -i 's/ssl-mode/ssl/' $new_dir/*
    fi
    go test -v $new_dir
done
