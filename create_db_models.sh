#!/usr/bin/env bash

go install github.com/volatiletech/sqlboiler/v4@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mssql@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-sqlite@latest

DBS=("sqlite3" "psql" "mssql" "mysql")

for db in "${DBS[@]}"; do
    new_dir=./models/$db
    mkdir -p $new_dir

    sqlboiler --output $new_dir $db
    sed -i 's/t.Parallel()//' $new_dir/*
    # Only needed for $b=sqlite3
    cp bntp_sqlite.db $new_dir
    go test -v $new_dir
done
