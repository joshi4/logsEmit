package main

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

func main() {
	c := 0
	tok := "default"
	for {
		c++
		msg := fmt.Sprintf("log-test:%d", c)
		tok = sendMsg(msg, tok)
		fmt.Println(msg)
		time.Sleep(1 * time.Second)
	}
}

func sendMsg(msg, tok string) string {

	cl := cloudwatchlogs.New(session.New(awsConfigWithSharedCredentials("us-west-2")))
	params := &cloudwatchlogs.PutLogEventsInput{
		LogEvents: []*cloudwatchlogs.InputLogEvent{ // Required
			{ // Required
				Message:   aws.String(msg),                     // Required
				Timestamp: aws.Int64(time.Now().Unix() * 1000), // Required
			},
		},
		LogGroupName:  aws.String("test"),  // Required
		LogStreamName: aws.String("test2"), // Required
		SequenceToken: aws.String(tok),
	}

	resp, err := cl.PutLogEvents(params)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return aws.StringValue(resp.NextSequenceToken)
}

func awsConfigWithSharedCredentials(region string) *aws.Config {
	return &aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewSharedCredentials("", os.Getenv("AWS_PROFILE")),
	}
}
