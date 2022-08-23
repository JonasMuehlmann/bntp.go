package main

import (
	"github.com/JonasMuehlmann/bntp.go/cmd"
)

func main() {
	cli := cmd.NewCli(cmd.WithAll())
	cli.Execute()
}
