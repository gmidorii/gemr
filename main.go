package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/emr"
)

const az = "ap-northeast-1"

func main() {
	session := session.New(&aws.Config{Region: aws.String(az)})
	svc := emr.New(session)

	instancConfig := emr.JobFlowInstancesConfig{
		Ec2KeyName:                  aws.String("emr-common"),
		Ec2SubnetId:                 aws.String("subnet-2e7e6f67"),
		HadoopVersion:               aws.String("2.8.4"),
		InstanceCount:               aws.Int64(1),
		KeepJobFlowAliveWhenNoSteps: aws.Bool(true),
		MasterInstanceType:          aws.String("m4.large"),
		SlaveInstanceType:           aws.String("m4.large"),
		TerminationProtected:        aws.Bool(false),
	}

	param := emr.RunJobFlowInput{
		Applications: []*emr.Application{
			{
				Name: aws.String("Hadoop"),
			},
			{
				Name: aws.String("Hive"),
			},
		},
		Instances:         &instancConfig,
		JobFlowRole:       aws.String("EMR_EC2_DefaultRole"),
		LogUri:            aws.String("s3://aws-logs-000001-ap-northeast-1/emr"),
		Name:              aws.String("Sample Cluster"),
		ReleaseLabel:      aws.String("emr-4.6.0"),
		ServiceRole:       aws.String("EMR_DefaultRole"),
		VisibleToAllUsers: aws.Bool(true),
	}

	resp, err := svc.RunJobFlow(&param)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.JobFlowId)
}
