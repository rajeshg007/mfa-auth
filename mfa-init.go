package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"os/user"
)

func mfa_init(c *cli.Context) error {
	_, err := os.Stat(FilePathClean(configfile))
	if err != nil {
		create_configfile()
	} else {
		fmt.Println("Config File is already present, exiting program")
		fmt.Println("To add new Profiles use the command mfa-auth add --profile (under construction)")
	}
	return nil
}

func create_configfile() {
	fmt.Println("config file doesnot exist, new file is being created")
	_, err := os.Create(FilePathClean(configfile))
	checkErr(err)
	accounts = make(map[string]Profile)
	addProfile(accounts)
}
func FilePathClean(s string) string {
	usr, err := user.Current()
	checkErr(err)
	if string(s[0]) == "~" {
		s = usr.HomeDir + s[1:]
	}
	return s
}

type Profile struct {
	Keyid     string
	Secretkey string
	Device    string
}

var accounts map[string]Profile
