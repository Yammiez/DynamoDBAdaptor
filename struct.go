package ddbproxy

import "github.com/aws/aws-sdk-go/service/dynamodb"

type LikunObject struct{
	name string
	age int
}


type DDBProxy struct{
	client *dynamodb.DynamoDB
}

type DDBProxyConfig struct{
	Region string
	Endpoint string
	SharedCredentialsLocation string
}

