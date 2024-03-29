package main

import (
	"github.com/urfave/cli/v2"
	"os"
	"sort"
)

var configfile string
var profile string
var assumeRole string

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
		&cli.BoolFlag{
			Name:    "reinitialize",
			Aliases: []string{"r"},
			Usage:   "recreate config file",
		},
	}
	loginFlags := []cli.Flag{
		&cli.StringFlag{
			Name:        "assume-role",
			Aliases:     []string{"ar"},
			Destination: &assumeRole,
			Value:       "",
			Usage:       "login and assume a role",
		},
	}
	app := &cli.App{
		Flags: append(loginFlags, commonFlags...),
		Action: func(c *cli.Context) error {
			return login(c)
		},
		Commands: []*cli.Command{
			{
				Name:    "init",
				Aliases: []string{"i"},
				Usage:   "create an new config file",
				Flags:   append(initFlags, commonFlags...),
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
			// {
			// 	Name:    "remove",
			// 	Aliases: []string{"r"},
			// 	Usage:   "Remove a Profile to config file",
			// 	Flags:   commonFlags,
			// 	Action: func(c *cli.Context) error {
			// 		return remove(c)
			// 	},
			// },
			{
				Name:    "update",
				Aliases: []string{"u"},
				Usage:   "Update an existing Profile in config file",
				Flags:   commonFlags,
				Action: func(c *cli.Context) error {
					return update(c)
				},
			},
			{
				Name:    "print_config",
				Aliases: []string{"p"},
				Usage:   "Print the contents of config file",
				Flags:   commonFlags,
				Action: func(c *cli.Context) error {
					return print(c)
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	checkErr(err)
}
