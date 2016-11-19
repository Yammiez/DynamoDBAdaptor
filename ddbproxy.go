package ddbproxy

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"log"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"fmt"
	"reflect"
)

func NewDDBProxy(config DDBProxyConfig) DDBProxy {
	// Create session
	sess, err := 	session.NewSession(&aws.Config{
		Region: aws.String(config.Region),
		Endpoint: aws.String(config.Endpoint),

		// Assuming default profile defined in shared credentials file
		Credentials: credentials.NewSharedCredentials(config.SharedCredentialsLocation, "default"),
	});

	if err != nil {
		log.Fatal("Failed to create AWS session. Exiting.")
	}

	// Create DynamoDB client with sess
	ddbClient := dynamodb.New(sess)
	fmt.Println("Created DynamoDB Client of type: ", reflect.TypeOf(ddbClient))
	return DDBProxy{ddbClient};
}

func (ddb *DDBProxy) ListTables () *dynamodb.ListTablesOutput{
	svc := ddb.client;

	fmt.Println("Testing dynamodb db Connection with Endpoint: ", svc.Endpoint);

	// aws dynamodb list-tables
	// Test connection
	params := &dynamodb.ListTablesInput{
		Limit: aws.Int64(100),
	}

	resp, err := svc.ListTables(params)

	if err != nil {
		fmt.Println("Failed calling ListTables. ERROR: ", err.Error());
		panic("")
	}

	fmt.Println(resp)
	fmt.Println("=================")
	return resp
}