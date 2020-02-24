package main

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	// "reflect"
	"bufio"
	"os/exec"
	"strings"
)

var mfa string

func login(c *cli.Context) error {
	fmt.Println("Using Configfile :", configfile)
	fmt.Println("Using Profile :", profile)
	if c.NArg() > 0 {
		mfa = c.Args().Get(0)
		awsprofile := getProfile()
		awsLogin(awsprofile)
		writeCredentialsFile()
	} else {
		fmt.Println("Please Pass MFA Code to login")
	}

	return nil
}

func getProfile() Profile {
	accounts := readConfigFromFile()
	if _, ok := accounts[profile]; !ok {
		fmt.Println("Credentials for selected Profile don't exist, Please add them")
		addProfile(accounts)
	}
	return accounts[profile]
}

func writeCredentialsFile() {
	awsFolder := FilePathClean("~/.aws")
	if _, err := os.Stat(awsFolder); os.IsNotExist(err) {
		os.Mkdir(awsFolder,0777)
	}
	file := FilePathClean("~/.aws/credentials")
	os.Remove(file)
	f, err := os.Create(file)
	checkErr(err)
	defer f.Close()
	w := bufio.NewWriter(f)
	_, err = w.WriteString("[default]\n")
	_, err = w.WriteString("aws_access_key_id=" + gst.Credentials.AccessKeyId + "\n")
	_, err = w.WriteString("aws_secret_access_key=" + gst.Credentials.SecretAccessKey + "\n")
	_, err = w.WriteString("aws_session_token=" + gst.Credentials.SessionToken + "\n")
	w.Flush()
}

func awsLogin(awsprofile Profile) {
	fmt.Println(awsprofile)
	fmt.Println("using mfa",mfa)
	cmd := exec.Command("aws", "sts", "get-session-token", "--serial-number", awsprofile.Device, "--token-code", mfa)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "AWS_ACCESS_KEY_ID="+awsprofile.Keyid)
	cmd.Env = append(cmd.Env, "AWS_SECRET_ACCESS_KEY="+awsprofile.Secretkey)
	out, err := cmd.CombinedOutput()
	checkErr(err)

	output := string(out)
	output = strings.Replace(output, "\n", "", -1)

	err = json.Unmarshal([]byte(output), &gst)
	checkErr(err)

}

type get_session_token struct {
	Credentials session `json:"Credentials"`
}

type session struct {
	SecretAccessKey string `json:"SecretAccessKey"`
	SessionToken    string `json:"SessionToken"`
	Expiration      string `json:"Expiration"`
	AccessKeyId     string `json:"AccessKeyId"`
}

var gst get_session_token
