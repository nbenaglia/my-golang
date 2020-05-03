package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type workingInfo struct {
	tableName    *string
	safeAccounts []string
}

func main() {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	info := workingInfo{tableName: aws.String("UserOrdersTable")}

	// Create DynamoDB client
	service := dynamodb.New(session)
	describeTable(service, &info)
	//delete(service, &info)
}

func keys(info *workingInfo) {
	queryInput := dynamodb.QueryInput{
		IndexName:            aws.String("Username"),
		ProjectionExpression: aws.String("Username"),
		TableName:            info.tableName,
	}
}

func delete(service *dynamodb.DynamoDB, info *workingInfo) {
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

func describeTable(service *dynamodb.DynamoDB, info *workingInfo) ([]*dynamodb.KeySchemaElement, error) {
	describeInput := &dynamodb.DescribeTableInput{TableName: info.tableName}
	table, err := service.DescribeTable(describeInput)
	if err != nil {
		fmt.Println("Got error calling DescribeTable")
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Printf("Items %d\n", table.Table.ItemCount)
	fmt.Printf("Creation date %s\n", table.Table.CreationDateTime.Format("2006-01-02T15:04:05"))
	return table.Table.KeySchema, nil
}

// 1. pass tableName
// 2. check keys
// 3. retrieve key list
// 4. iterate and delete each item (go routine)
