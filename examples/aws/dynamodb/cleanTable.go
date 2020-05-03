package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type dynamoInfo struct {
	tableName *string
}

func main() {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	info := dynamoInfo{tableName: aws.String("UserOrdersTable")}

	// Create DynamoDB client
	service := dynamodb.New(session)
	//delete(service, &info)
	describeTable(service, &info)
}

func delete(service *dynamodb.DynamoDB, info *dynamoInfo) {
	input := &dynamodb.DeleteItemInput{
		TableName: info.tableName,
		Key: map[string]*dynamodb.AttributeValue{
			"Username": {
				S: aws.String("alexdebrie"),
			},
			"OrderId": {
				S: aws.String("20160630-12928"),
			},
		},
	}

	_, err := service.DeleteItem(input)
	if err != nil {
		fmt.Println("Got error calling DeleteItem")
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Delete executed.")
}

func describeTable(service *dynamodb.DynamoDB, info *dynamoInfo) {
	describeInput := &dynamodb.DescribeTableInput{TableName: info.tableName}
	table, err := service.DescribeTable(describeInput)
	if err != nil {
		fmt.Println("Got error calling DescribeTable")
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Items %d\n", table.Table.ItemCount)
	fmt.Printf("Creation date %s\n", table.Table.CreationDateTime.Format("2006-01-02T15:04:05"))
	keySchema := table.Table.KeySchema
	for i, v := range keySchema {
		fmt.Printf("Key %d\n", i)
		fmt.Printf("Values %s\n", v)
	}
}

// 1. pass tableName
// 2. check keys
// 3. retrieve key list
// 4. iterate and delete each item (go routine)
