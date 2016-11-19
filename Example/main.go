package main

import (
	"fmt"
	"github.com/Yammiez/DynamoDBAdaptor"
)

func main() {
	fmt.Printf("Hello World!")
	ddbClient := ddbproxy.NewDDBProxy( ddbproxy.DDBProxyConfig{"us-west-2", "https://dynamodb.us-west-2.amazonaws.com", "W:\\GoWork\\key\\AWS\\credentials"})
	println(ddbClient.ListTables())

}
