package main

import (
	"os"

	"github.com/JonasMuehlmann/bntp.go/bntp/backend"
	"github.com/JonasMuehlmann/bntp.go/cmd"
)

func main() {
	backend := new(backend.Backend)

	cmd.Execute(backend, os.Stderr)
}
