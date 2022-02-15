#!/usr/bin/env sh

rm test/data/bntp_test.db
sqlite3 test/data/bntp_test.db <bntp.sql
