package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/golang/glog"

	"fmt"
)

var (
	region                = "eu-west-1"
	topicArn              = ""
	aws_access_key_id     = ""
	aws_secret_access_key = ""
	token                 = ""
)

func main() {
	message := "It's a test."

	if message == "" || topicArn == "" {
		glog.Error("You must supply a message and topic ARN")
		return
	}

	creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Credentials: creds,
			Region:      &region,
		},
	}))

	svc := sns.New(sess)

	result, err := svc.Publish(&sns.PublishInput{
		Message:  aws.String(message),
		TopicArn: &topicArn,
	})

	if err != nil {
		glog.Error(err.Error())
	}

	fmt.Println(*result.MessageId)
}
