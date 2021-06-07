package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

func add(c *cli.Context) error {
	accounts, _ := readConfigFromFile()
	addProfile(accounts)
	return nil
}

func addProfile(accounts map[string]Profile) {
	fmt.Println("Using Profile: " + profile)
	keyid := readFromIO("Enter Access Key ID\t:")
	secret := readFromIO("Enter Secret Access Key\t:")
	device := getDeviceId(keyid, secret)
	fmt.Println("Using Device ", device)
	accounts[profile] = Profile{keyid, secret, device}
	writeConfigToFile(accounts)
}

func getDeviceId(keyid string, secret string) string {
	client := getSTSClientFromKeys(keyid, secret)
	identity, err := client.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
	checkErr(err)
	serial := *identity.Arn
	serial = strings.Replace(serial, ":user/", ":mfa/", -1)
	return serial
}

func writeConfigToFile(accounts map[string]Profile) {
	encodeFile, err := os.Create(FilePathClean(configfile))
	defer encodeFile.Close()
	encoder := gob.NewEncoder(encodeFile)
	if err = encoder.Encode(accounts); err != nil {
		fmt.Println(err)
	}
}

func readConfigFromFile() (map[string]Profile, bool) {
	absPath := FilePathClean(configfile)
	isNew := false
	_, err := os.Stat(absPath)
	if err != nil {
		create_configfile()
		isNew = true
	}
	decodeFile, err := os.Open(absPath)
	checkErr(err)
	defer decodeFile.Close()

	decoder := gob.NewDecoder(decodeFile)
	accounts := make(map[string]Profile)
	decoder.Decode(&accounts)
	return accounts, isNew
}
