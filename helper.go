package main

import (
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

var awsRegion = ""
var snsArn = ""

func sendSMS(quote string) {
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(awsRegion)}))
	svc := sns.New(sess)
	params := &sns.PublishInput{
		Message:  aws.String(quote),
		TopicArn: aws.String(snsArn),
	}
	resp, err := svc.Publish(params)

	if err != nil {
		log.Println("Failed to publish the msg:", err)
	}

	log.Println("Message was published successfully!", resp)
}

func join(quotes ...string) string {
	var sb strings.Builder
	for _, q := range quotes {
		sb.WriteString(q)
	}
	return sb.String()

}

func equal(a, b []Media) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v.Id != b[i].Id {
			return false
		}
	}
	return true
}

func checkVariables() {

}
