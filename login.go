package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go-v2/service/sts/types"
	"github.com/urfave/cli/v2"
	"os"
)

func login(c *cli.Context) error {
	fmt.Println("Using Configfile :", configfile)
	fmt.Println("Using Profile :", profile)
	fmt.Println("Assuming Role:", assumeRole)
	if c.NArg() > 0 {
		mfa := c.Args().Get(0)
		// mfa := readFromIO("Please enter MFA Token: ")
		awsprofile, isNew := getProfile()
		if isNew == true {
			mfa = readFromIO("MFA Token Might have expired in the time credentials were entered, Please enter new MFA: ")
		}
		creds := awsLogin(awsprofile, mfa)
		if assumeRole != "" {
			client := getSTSClientFromCredentials(creds)
			identity, err := client.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
			checkErr(err)
			assumeRoleARN := "arn:aws:iam::" + *identity.Account + ":role/" + assumeRole
			username := getUser()
			input := &sts.AssumeRoleInput{
				RoleArn:         &assumeRoleARN,
				RoleSessionName: &username,
			}
			output, err := client.AssumeRole(context.TODO(), input)
			checkErr(err)
			creds = *output.Credentials
		}
		writeCredentialsFile(creds)
	} else {
		fmt.Println("Please Pass MFA Code to login")
	}

	return nil
}

func getProfile() (Profile, bool) {
	accounts, isNew := readConfigFromFile()
	if _, ok := accounts[profile]; !ok {
		fmt.Println("Credentials for selected Profile don't exist, Please add them")
		addProfile(accounts)
	}
	return accounts[profile], isNew
}

func writeCredentialsFile(creds types.Credentials) {
	awsFolder := FilePathClean("~/.aws")
	if _, err := os.Stat(awsFolder); os.IsNotExist(err) {
		os.Mkdir(awsFolder, 0777)
	}
	file := FilePathClean("~/.aws/credentials")
	os.Remove(file)
	f, err := os.Create(file)
	checkErr(err)
	defer f.Close()
	w := bufio.NewWriter(f)
	_, err = w.WriteString("[default]\n")
	_, err = w.WriteString("aws_access_key_id=" + *creds.AccessKeyId + "\n")
	_, err = w.WriteString("aws_secret_access_key=" + *creds.SecretAccessKey + "\n")
	_, err = w.WriteString("aws_session_token=" + *creds.SessionToken + "\n")
	w.Flush()
}

func awsLogin(awsprofile Profile, mfa string) types.Credentials {
	fmt.Println("using mfa", mfa)
	client := getSTSClientFromProfile(awsprofile)
	duration := int32(43200)
	input := sts.GetSessionTokenInput{}
	input.DurationSeconds = &duration
	input.SerialNumber = &awsprofile.Device
	input.TokenCode = &mfa

	output, err := client.GetSessionToken(context.TODO(), &input)
	checkErr(err)

	return *output.Credentials
}
