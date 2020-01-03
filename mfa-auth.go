package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"sort"
)

var configfile string
var profile string

func main() {
	commonFlags := []cli.Flag{
		&cli.StringFlag{
			Name:        "configfile",
			Aliases:     []string{"c"},
			Value:       "~/.mfa-auth",
			Usage:       "Base Configfile for mfa-auth",
			Destination: &configfile,
			EnvVars:     []string{"MFA_AUTH_CONFIG"},
		},
		&cli.StringFlag{
			Name:        "profile",
			Aliases:     []string{"p"},
			Value:       "default",
			Usage:       "pass the flag to use a non default profile",
			Destination: &profile,
			EnvVars:     []string{"MFA_AUTH_PROFILE"},
		},
	}
	initFlags := []cli.Flag{
		&cli.StringFlag{
			Name:        "configfile",
			Aliases:     []string{"c"},
			Value:       "~/.mfa-auth",
			Usage:       "Base Configfile for mfa-auth",
			Destination: &configfile,
			EnvVars:     []string{"MFA_AUTH_CONFIG"},
		},
		&cli.StringFlag{
			Name:        "profile",
			Aliases:     []string{"p"},
			Value:       "default",
			Usage:       "pass the flag to use a non default profile",
			Destination: &profile,
			EnvVars:     []string{"MFA_AUTH_PROFILE"},
		},
		&cli.BoolFlag{
			Name:    "reinitialize",
			Aliases: []string{"r"},
			Usage:   "recreate config file",
		},
	}
	app := &cli.App{
		Flags: commonFlags,
		Action: func(c *cli.Context) error {
			return login(c)
		},
		Commands: []*cli.Command{
			{
				Name:    "init",
				Aliases: []string{"i"},
				Usage:   "create an new config file",
				Flags:   initFlags,
				Action: func(c *cli.Context) error {
					return mfa_init(c)
				},
			},
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "Add a new Profile to config file",
				Flags:   commonFlags,
				Action: func(c *cli.Context) error {
					return add(c)
				},
			},
			{
				Name:    "remove",
				Aliases: []string{"a"},
				Usage:   "Remove a Profile to config file",
				Flags:   commonFlags,
				Action: func(c *cli.Context) error {
					return remove(c)
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	checkErr(err)
}

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
