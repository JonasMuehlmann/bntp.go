package main

//go:generate go run ./tools/schema_converter
//go:generate go run ./tools/generate_external_cli_documentation
//go:generate go run ./tools/generate_domain_models
//go:generate go run ./tools/generate_repository_interfaces
//go:generate go run ./tools/generate_sql_repositories
//go:generate go run ./tools/generate_sql_repository_filter_converters
//go:generate go run ./tools/generate_sql_repository_model_converters
//go:generate go run ./tools/generate_sql_repository_updater_converters

import (
	"github.com/JonasMuehlmann/bntp.go/bntp/backend"
	"github.com/JonasMuehlmann/bntp.go/cmd"
)

func main() {
	backend := new(backend.Backend)
	cmd.Execute(backend)
}
