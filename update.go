package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"encoding/gob"
	"strings"
	"bufio"
)

func update(c *cli.Context) error {
	absPath := FilePathClean(configfile)
	_, err := os.Stat(absPath)
	if err != nil {
		fmt.Println("Configuration not present, please use `init` or `add` to create configuration")
	}
	decodeFile, err := os.Open(absPath)
	checkErr(err)
	defer decodeFile.Close()
	decoder := gob.NewDecoder(decodeFile)
	accounts := make(map[string]Profile)
	decoder.Decode(&accounts)
	oldProfile := accounts[profile]
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Updating Profile: " + profile)
	fmt.Print("Enter Access Key ID ",oldProfile.Keyid,"\t:")
	keyid, err := reader.ReadString('\n')
	keyid = strings.Replace(keyid, "\n", "", -1)
	checkErr(err)
	fmt.Print("Enter Secret Access Key ",oldProfile.Secretkey,"\t:")
	secret, err := reader.ReadString('\n')
	secret = strings.Replace(secret, "\n", "", -1)
	checkErr(err)
	device := getDeviceId(keyid, secret)
	fmt.Println("Using Device ", device)
	accounts[profile] = Profile{keyid, secret, device}
	fmt.Println(accounts)
	writeConfigToFile(accounts)
	fmt.Println("update successful")
	return nil
}
