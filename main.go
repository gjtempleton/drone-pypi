package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	_ "github.com/joho/godotenv/autoload"
)

var build = "0" // build number set at compile-time

func main() {
	app := cli.NewApp()
	app.Name = "pypi plugin"
	app.Usage = "pypi publish plugin"
	app.Action = run
	app.Version = fmt.Sprintf("0.0.%s", build)
	app.Flags = []cli.Flag{

		//
		// repo args
		//
		cli.StringFlag{
			Name:   "repository",
			Usage:  "pypi repository URL",
			EnvVar: "PLUGIN_REPOSITORY",
		},
		cli.StringFlag{
			Name:   "username",
			Usage:  "pypi username",
			EnvVar: "PLUGIN_USERNAME",
		},
		cli.StringFlag{
			Name:   "password",
			Usage:  "pypi password",
			EnvVar: "PLUGIN_PASSWORD",
		},
		cli.StringSliceFlag{
			Name:   "distributions",
			Usage:  "distribution types to deploy",
			EnvVar: "PLUGIN_DISTRIBUTIONS",
		},
	}

	app.Run(os.Args)
}

func run(c *cli.Context) {
	plugin := Plugin{
		Repository:    c.String("repository"),
		Username:      c.String("username"),
		Password:      c.String("password"),
		Distributions: c.StringSlice("distributions"),
	}

	if err := plugin.Exec(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
