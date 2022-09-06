package main

import (
	"os"
	"reflect"
	"strings"

	"github.com/JonasMuehlmann/bntp.go/cmd"
)

func main() {
	cli, err := cmd.NewCli(cmd.WithAll())
	// The config subcommand does not need the error, so it can be ignored, this is also needed for setting up the initial database in the first place.
	if err != nil && strings.Contains(err.Error(), "data_source") && len(os.Args) > 1 && os.Args[1] == "config" {
	} else if err != nil {
		cli.Logger.Error(err)
	}

	err = cli.Execute()
	if err != nil {
		cli.StdErr.Write([]byte("ENDLOG\n"))

		errMessage := map[string]any{"errorType": reflect.TypeOf(err).Name(), "error": err}

		errSerialized, err := cli.BNTPBackend.Marshallers[cli.OutFormat].Marshall(errMessage)
		if err != nil {
			cli.Logger.Error(err)
			cli.StdErr.Write([]byte(err.Error()))
		}

		cli.StdErr.Write([]byte(errSerialized))
	}
}
