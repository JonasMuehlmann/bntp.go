package main

//go:generate go run ./tools/generate_config_key_constants
//go:generate go run ./tools/generate_external_cli_documentation

import "github.com/JonasMuehlmann/bntp.go/cmd"

func main() {
	cmd.Execute()
}
