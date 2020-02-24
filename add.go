package main

import (
	"bufio"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
	"log"
	"strings"
)

func add(c *cli.Context) error {
	accounts := readConfigFromFile()
	addProfile(accounts)
	return nil
}

func addProfile(accounts map[string]Profile) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Using Profile: " + profile)
	fmt.Print("Enter Access Key ID\t:")
	keyid, err := reader.ReadString('\n')
	keyid = strings.Replace(keyid, "\n", "", -1)
	checkErr(err)
	fmt.Print("Enter Secret Access Key\t:")
	secret, err := reader.ReadString('\n')
	secret = strings.Replace(secret, "\n", "", -1)
	checkErr(err)
	device := getDeviceId(keyid, secret)
	fmt.Println("Using Device ", device)
	accounts[profile] = Profile{keyid, secret, device}
	fmt.Println(accounts)
	writeConfigToFile(accounts)
}

type caller_identity struct {
	Account string `json:"Account"`
	UserId  string `json:"UserId"`
	Arn     string `json:"Arn"`
}

func getDeviceId(keyid string, secret string) string {
	cmd := exec.Command("aws", "sts", "get-caller-identity")
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "AWS_ACCESS_KEY_ID="+keyid)
	cmd.Env = append(cmd.Env, "AWS_SECRET_ACCESS_KEY="+secret)
	out, err := cmd.CombinedOutput()
	checkErr(err)
	var gci caller_identity
	output := string(out)
	output = strings.Replace(output, "\n", "", -1)

	err = json.Unmarshal([]byte(output), &gci)
	if err != nil {
		log.Fatal("Unable to validate keys, Please verify them and update again")
		log.Fatal(err)
	}
	serial := gci.Arn
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

func readConfigFromFile() map[string]Profile {
	absPath := FilePathClean(configfile)
	_, err := os.Stat(absPath)
	if err != nil {
		create_configfile()
		fmt.Print("MFA Token Might have expired in the time credentials were entered, Please enter new MFA: ")
		reader := bufio.NewReader(os.Stdin)
		mfa, err = reader.ReadString('\n')
		checkErr(err)
		mfa = strings.Replace(mfa, "\n", "", -1)
	}
	decodeFile, err := os.Open(absPath)
	checkErr(err)
	defer decodeFile.Close()

	decoder := gob.NewDecoder(decodeFile)
	accounts := make(map[string]Profile)
	decoder.Decode(&accounts)
	return accounts
}
