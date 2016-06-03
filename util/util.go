package util

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

func AwsConfigWithSharedCredentials(region string) *aws.Config {
	return &aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewSharedCredentials("", os.Getenv("AWS_PROFILE")),
	}
}
