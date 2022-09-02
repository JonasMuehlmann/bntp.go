package main

import (
	"reflect"

	"github.com/JonasMuehlmann/bntp.go/cmd"
)

func main() {
	cli := cmd.NewCli(cmd.WithAll())

	err := cli.Execute()
	if err != nil {
		cli.Logger.Error(err)
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
