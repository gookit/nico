package main

import (
	"runtime"

	"github.com/gookit/gcli"
	"github.com/gookit/gcli/builtin"
	"github.com/gookit/lako/cmd"
)

// run:
// go run ./cmd/lako
// go build ./cmd/lako && ./lako
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	app := gcli.NewApp(func(app *gcli.App) {
		app.Version = "1.0.6"
		app.Description = "this is lako cli application"
		app.On(gcli.EvtInit, func(data ...interface{}) {
			// do something...
			// fmt.Println("init app")
		})

		// app.SetVerbose(gcli.VerbDebug)
		// app.DefaultCommand("example")
		app.Logo.Text = `   ________    _______
  / ____/ /   /  _/   |  ____  ____
 / /   / /    / // /| | / __ \/ __ \
/ /___/ /____/ // ___ |/ /_/ / /_/ /
\____/_____/___/_/  |_/ .___/ .___/
                     /_/   /_/`
	})

	app.Add(
		cmd.CreateProjectCommand(),
		builtin.GenAutoCompleteScript(),
	)

	app.Run()
}