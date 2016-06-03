package main

import (
	"fmt"
	"time"

	"github.com/joshi4/logsEmit/util"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

func main() {
	c := 0
	tok := "default"
	for {
		c++
		msg := fmt.Sprintf("log-test:%d", c)
		tokn, err := sendMsg(msg, tok)
		if err != nil {
			fmt.Println(err)
			return
		}
		tok = tokn
		fmt.Println(msg)
		time.Sleep(1 * time.Second)
	}
}

func sendMsg(msg, tok string) (string, error) {
	cl := cloudwatchlogs.New(session.New(util.AwsConfigWithSharedCredentials("us-west-2")))
	params := &cloudwatchlogs.PutLogEventsInput{
		LogEvents: []*cloudwatchlogs.InputLogEvent{ // Required
			{ // Required
				Message:   aws.String(msg),                     // Required
				Timestamp: aws.Int64(time.Now().Unix() * 1000), // Required
			},
		},
		LogGroupName:  aws.String("test"), // Required
		LogStreamName: aws.String("test"), // Required
		SequenceToken: aws.String(tok),
	}

	resp, err := cl.PutLogEvents(params)
	if err != nil {
		return "", err
	}
	return aws.StringValue(resp.NextSequenceToken), nil
}