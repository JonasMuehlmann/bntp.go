#!/usr/bin/env sh

rm model/repository/**/*_repository.go || true
go run ./tools/generate_sql_repositories
go run ./tools/generate_repository_interfaces && git restore model/repository/document_content_repository.go model/repository/fs/document_content_repository.go
