package main

import (
	"log"
	"os"
	"sort"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Version = Version
	app.Name = AppName
	app.Usage = AppUsage
	app.Flags = options
	app.EnableBashCompletion = true

	commonCommands := []cli.Command{
		commandEntry,
	}

	daemonCommands, err := createDaemonCommands(AppName, AppUsage)

	if err != nil {
		log.Fatal(err)
		return
	}

	app.Commands = append(commonCommands, daemonCommands...)

	sort.Sort(cli.CommandsByName(app.Commands))

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
