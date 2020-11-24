package repository

import (
	"context"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/domarcio/bexs/src/entity"
)

// DynamoDBRepo dynamodb repository
type DynamoDBRepo struct {
	client dynamodbiface.DynamoDBAPI
}

// NewAirportDynamoDB create new repository
func NewAirportDynamoDB(client dynamodbiface.DynamoDBAPI) *DynamoDBRepo {
	return &DynamoDBRepo{
		client: client,
	}
}

// Get airport by code
func (c *DynamoDBRepo) Get(ctx context.Context, code string) (*entity.Airport, error) {
	return nil, nil
}
