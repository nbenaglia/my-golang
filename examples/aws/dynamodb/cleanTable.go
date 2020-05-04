package main

import (
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type workingInfo struct {
	tableName   *string
	role 		*string
	region 		*string
}

func main() {
	info := workingInfo{
		tableName: aws.String("UserOrderTable"),
		role: aws.String("arn:aws:iam::<ACCOUNT_NUMBER>:role/myrole"),
		region: aws.String("eu-west-1"),
	}

	// Create DynamoDB client
	service, err := getDynamoClient(info.role, info.region)
	if err != nil {
		log.Println("Got error calling dynamoClient")
		log.Println(err.Error())
		return
	}
	describeTable(service, &info)
	//delete(service, &info)
}

func getDynamoClient(roleArn *string, region *string) (*dynamodb.DynamoDB, error) {
	log.Printf("Initializing new dynamo client for role %s", *roleArn)

	if *roleArn == "" {
		sess := session.Must(session.NewSession())
		client := dynamodb.New(sess, &aws.Config{Region: region})
		return client, nil
	}

	sess := session.Must(session.NewSession())
	credentials := stscreds.NewCredentials(sess, *roleArn)
	client := dynamodb.New(sess, &aws.Config{Credentials: credentials, Region: region})

	log.Printf("Client is %s", client.ServiceName)
	return client, nil
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
		log.Println("Got error calling DeleteItem")
		log.Println(err.Error())
		return
	}
	log.Printf("Delete executed.")
}

func describeTable(service *dynamodb.DynamoDB, info *workingInfo) ([]*dynamodb.KeySchemaElement, error) {
	describeInput := &dynamodb.DescribeTableInput{TableName: info.tableName}
	table, err := service.DescribeTable(describeInput)
	if err != nil {
		log.Println("Got error calling DescribeTable")
		log.Println(err.Error())
		return nil, err
	}
	log.Printf("Items %d\n", table.Table.ItemCount)
	log.Printf("Creation date %s\n", table.Table.CreationDateTime.Format("2006-01-02T15:04:05"))
	return table.Table.KeySchema, nil
}

// 1. pass tableName
// 2. check keys
// 3. retrieve key list
// 4. iterate and delete each item (go routine)
