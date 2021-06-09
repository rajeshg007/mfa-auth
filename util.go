package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go-v2/service/sts/types"
	"log"
	"os"
	"os/user"
	"strings"
)

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func getSTSClientFromProfile(profileValue Profile) *sts.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		// Hard coded credentials.
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: profileValue.Keyid, SecretAccessKey: profileValue.Secretkey,
				Source: "Profile Based Credentials",
			},
		}))
	if err != nil {
		log.Fatal(err)
	}
	client := sts.NewFromConfig(cfg)
	return client
}

func getSTSClientFromKeys(keyid string, secret string) *sts.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		// Hard coded credentials.
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: keyid, SecretAccessKey: secret,
				Source: "Profile Based Credentials",
			},
		}))
	if err != nil {
		log.Fatal(err)
	}
	client := sts.NewFromConfig(cfg)
	return client
}

func getSTSClientFromCredentials(creds types.Credentials) *sts.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: *creds.AccessKeyId, SecretAccessKey: *creds.SecretAccessKey, SessionToken: *creds.SessionToken,
				Source: "Profile Based Credentials",
			},
		}))
	if err != nil {
		log.Fatal(err)
	}
	client := sts.NewFromConfig(cfg)
	return client
}

func readFromIO(ques string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(ques)
	readVal, err := reader.ReadString('\n')
	checkErr(err)
	readVal = strings.Replace(readVal, "\n", "", -1)
	return readVal
}

func getUser() string {
	user, _ := user.Current()
	return user.Username
}
