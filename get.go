package main

import (
	"fmt"

	"github.com/joshi4/logsEmit/util"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

func main() {

	if err := getLogs(); err != nil {
		fmt.Println(err)
		return
	}
}

func getLogs() error {
	cl := cloudwatchlogs.New(session.New(util.AwsConfigWithSharedCredentials("us-west-2")))
	params := &cloudwatchlogs.GetLogEventsInput{
		LogGroupName:  aws.String("test"), // Required
		LogStreamName: aws.String("test"), // Required
	}

	printPage := func(p *cloudwatchlogs.GetLogEventsOutput, lastPage bool) bool {
		printLogs(p.Events)
		return true
	}

	err := cl.GetLogEventsPages(params, printPage)
	if err != nil {
		return err
	}
	return nil
}

func printLogs(events []*cloudwatchlogs.OutputLogEvent) {
	for _, e := range events {
		fmt.Println(aws.StringValue(e.Message))
	}
}
