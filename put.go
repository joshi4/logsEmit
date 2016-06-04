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
	c := 49
	tok, err := tokenForStream("test")
	if err != nil {
		fmt.Println(err)
		return
	}

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

func tokenForStream(stream string) (string, error) {
	cl := cloudwatchlogs.New(session.New(util.AwsConfigWithSharedCredentials("us-west-2")))
	params := &cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName:        aws.String("test"),
		LogStreamNamePrefix: aws.String(stream),
	}

	tok := ""
	getTokenForStream := func(p *cloudwatchlogs.DescribeLogStreamsOutput, lp bool) bool {
		for _, ls := range p.LogStreams {
			if aws.StringValue(ls.LogStreamName) == stream {
				tok = aws.StringValue(ls.UploadSequenceToken)
				return false
			}
		}
		if lp {
			return false
		}
		return true
	}

	err := cl.DescribeLogStreamsPages(params, getTokenForStream)

	if err == nil && tok == "" {
		err = fmt.Errorf("No token found for stream:%s", stream)
	}
	return tok, err
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
