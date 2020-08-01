package Lib

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type MongoClient struct {
	Host    string
	Port    string
	Context context.Context
}

func NewMongoClient(mongoStruct *MongoClient) *MongoClient {
	return &MongoClient{mongoStruct.Host, mongoStruct.Port, mongoStruct.Context}
}

func (m *MongoClient) Client() (*mongo.Client, error) {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + m.Host + ":" + m.Port))

	if err != nil {
		return client, err
	}

	err = client.Connect(m.Context)
	if err != nil {
		return client, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		return client, err
	}

	return client, err
}
