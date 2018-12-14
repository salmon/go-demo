package mongo

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
	"github.com/sirupsen/logrus"
	"time"
	"transwarp/tos-app-market/pkg/setting"
)

type MongoClient struct {
	client *mongo.Client
}

var mongoClient *MongoClient

func GetDefaultRedisClient() *MongoClient {
	var err error

	if mongoClient == nil {
		mongoClient.client, err = mongo.NewClient(fmt.Sprintf("mongodb://%s", setting.Config.MongoConfig.Addr))
		if err != nil {
			logrus.Errorf("GetDefaultRedisClient mongoClient error %v", err)
		}
		ctx, _ := context.WithTimeout(context.Background(), 60 * time.Second)
		err = mongoClient.client.Connect(ctx)
		if err != nil {
			logrus.Errorf("GetDefaultRedisClient mongoClient error %v", err)
		}
		err = mongoClient.client.Ping(ctx, readpref.Primary())
		if err != nil {
			logrus.Errorf("GetDefaultRedisClient mongoClient error %v", err)
		}
	}

	return mongoClient
}

func (mongoClient *MongoClient) GetClient() *mongo.Client {
	return mongoClient.client
}
