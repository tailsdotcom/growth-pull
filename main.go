package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go/aws"
)

type auth struct {
	Auth string `json:"auth"`
}
type docker struct {
	Auths map[string]auth `json:"auths"`
}

func main() {
	cfg, err := config.LoadDefaultConfig()
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	ids := os.Getenv("AWS_ECR_REGISTRY_IDS")
	if ids == "" {
		stsSvc := sts.NewFromConfig(cfg)
		callerID, err := stsSvc.GetCallerIdentity(context.Background(), &sts.GetCallerIdentityInput{})
		if err != nil {
			log.Fatalf("unable to get identity, %v", err)
		}
		ids = *callerID.Account
	}
	svc := ecr.NewFromConfig(cfg)
	resp, err := svc.GetAuthorizationToken(context.Background(), &ecr.GetAuthorizationTokenInput{
		RegistryIds: aws.StringSlice(strings.Split(ids, " ")),
	})
	if err != nil {
		log.Fatalf("unable to authorization token, %v", err)
	}
	config := docker{
		Auths: make(map[string]auth),
	}
	for _, repo := range resp.AuthorizationData {
		config.Auths[(*repo.ProxyEndpoint)[8:]] = auth{Auth: *repo.AuthorizationToken}
	}
	out, err := json.Marshal(config)
	if err != nil {
		log.Fatalf("unable marshal json, %v", err)
	}
	err = ioutil.WriteFile(os.Args[1], out, 0644)
	if err != nil {
		log.Fatalf("unable to write output, %v", err)
	}
}
