#!/usr/bin/env bash

# Enable ** glob
shopt -s globstar

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
    if [[ $db == "mssql" ]]; then
        cp schema/bntp_sqlserver_test.sql tables_schema.sql
        cp schema/bntp_sqlserver_test.sql $new_dir/tables_schema.sql
    fi
    go mod tidy
    go get -t github.com/JonasMuehlmann/bntp.go/$new_dir
    go test -v $new_dir
done

REPOSITORIES=(libbookmarks libdocuments liblinks libtags)

# Users won't need these tests anymore
rm ./models/**/*_test.go

# Copy generated models to libraries
for repo in "${REPOSITORIES[@]}"; do
    for db in "${DBS[@]}"; do
        repo_dir="./pkg/$repo/repository/$db"
        mkdir -p "$repo_dir"
        # Not sure if these are mandatory, but keeping the mcan't hurt
        cp -t "$repo_dir/" ./models/$db/{boil_queries,boil_types,boil_table_names,boil_view_names,*_upsert}.go
        # Strip prefix lib
        tmp=${repo#lib}
        # Strip suffix s
        tmp=${tmp%s}
        mv ./models/$db/*$tmp*.go "$repo_dir/"

        sed -i "s/package models/ package ${tmp}_repository/g" "$repo_dir/"*.go
    done
done

# Remove temp dir of generation
rm -r models
