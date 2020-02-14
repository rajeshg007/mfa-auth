package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func print(c *cli.Context) error {
	accounts := readConfigFromFile()
	printProfile(accounts)
	return nil
}

func printProfile(accounts map[string]Profile) {
	fmt.Println(accounts)
}