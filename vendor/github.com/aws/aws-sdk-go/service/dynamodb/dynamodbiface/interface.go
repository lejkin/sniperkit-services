// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package dynamodbiface provides an interface to enable mocking the Amazon DynamoDB service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package dynamodbiface

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// DynamoDBAPI provides an interface to enable mocking the
// dynamodb.DynamoDB service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon DynamoDB.
//    func myFunc(svc dynamodbiface.DynamoDBAPI) bool {
//        // Make svc.BatchGetItem request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := dynamodb.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockDynamoDBClient struct {
//        dynamodbiface.DynamoDBAPI
//    }
//    func (m *mockDynamoDBClient) BatchGetItem(input *dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error) {
//        // mock response/functionality
//    }
//
//    TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockDynamoDBClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type DynamoDBAPI interface {
	BatchGetItemRequest(*dynamodb.BatchGetItemInput) (*request.Request, *dynamodb.BatchGetItemOutput)

	BatchGetItem(*dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error)

	BatchGetItemPages(*dynamodb.BatchGetItemInput, func(*dynamodb.BatchGetItemOutput, bool) bool) error

	BatchWriteItemRequest(*dynamodb.BatchWriteItemInput) (*request.Request, *dynamodb.BatchWriteItemOutput)

	BatchWriteItem(*dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error)

	CreateTableRequest(*dynamodb.CreateTableInput) (*request.Request, *dynamodb.CreateTableOutput)

	CreateTable(*dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error)

	DeleteItemRequest(*dynamodb.DeleteItemInput) (*request.Request, *dynamodb.DeleteItemOutput)

	DeleteItem(*dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error)

	DeleteTableRequest(*dynamodb.DeleteTableInput) (*request.Request, *dynamodb.DeleteTableOutput)

	DeleteTable(*dynamodb.DeleteTableInput) (*dynamodb.DeleteTableOutput, error)

	DescribeLimitsRequest(*dynamodb.DescribeLimitsInput) (*request.Request, *dynamodb.DescribeLimitsOutput)

	DescribeLimits(*dynamodb.DescribeLimitsInput) (*dynamodb.DescribeLimitsOutput, error)

	DescribeTableRequest(*dynamodb.DescribeTableInput) (*request.Request, *dynamodb.DescribeTableOutput)

	DescribeTable(*dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error)

	GetItemRequest(*dynamodb.GetItemInput) (*request.Request, *dynamodb.GetItemOutput)

	GetItem(*dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)

	ListTablesRequest(*dynamodb.ListTablesInput) (*request.Request, *dynamodb.ListTablesOutput)

	ListTables(*dynamodb.ListTablesInput) (*dynamodb.ListTablesOutput, error)

	ListTablesPages(*dynamodb.ListTablesInput, func(*dynamodb.ListTablesOutput, bool) bool) error

	ListTagsOfResourceRequest(*dynamodb.ListTagsOfResourceInput) (*request.Request, *dynamodb.ListTagsOfResourceOutput)

	ListTagsOfResource(*dynamodb.ListTagsOfResourceInput) (*dynamodb.ListTagsOfResourceOutput, error)

	PutItemRequest(*dynamodb.PutItemInput) (*request.Request, *dynamodb.PutItemOutput)

	PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)

	QueryRequest(*dynamodb.QueryInput) (*request.Request, *dynamodb.QueryOutput)

	Query(*dynamodb.QueryInput) (*dynamodb.QueryOutput, error)

	QueryPages(*dynamodb.QueryInput, func(*dynamodb.QueryOutput, bool) bool) error

	ScanRequest(*dynamodb.ScanInput) (*request.Request, *dynamodb.ScanOutput)

	Scan(*dynamodb.ScanInput) (*dynamodb.ScanOutput, error)

	ScanPages(*dynamodb.ScanInput, func(*dynamodb.ScanOutput, bool) bool) error

	TagResourceRequest(*dynamodb.TagResourceInput) (*request.Request, *dynamodb.TagResourceOutput)

	TagResource(*dynamodb.TagResourceInput) (*dynamodb.TagResourceOutput, error)

	UntagResourceRequest(*dynamodb.UntagResourceInput) (*request.Request, *dynamodb.UntagResourceOutput)

	UntagResource(*dynamodb.UntagResourceInput) (*dynamodb.UntagResourceOutput, error)

	UpdateItemRequest(*dynamodb.UpdateItemInput) (*request.Request, *dynamodb.UpdateItemOutput)

	UpdateItem(*dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error)

	UpdateTableRequest(*dynamodb.UpdateTableInput) (*request.Request, *dynamodb.UpdateTableOutput)

	UpdateTable(*dynamodb.UpdateTableInput) (*dynamodb.UpdateTableOutput, error)

	WaitUntilTableExists(*dynamodb.DescribeTableInput) error

	WaitUntilTableNotExists(*dynamodb.DescribeTableInput) error
}

var _ DynamoDBAPI = (*dynamodb.DynamoDB)(nil)
