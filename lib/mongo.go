package lib

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoDBOptions struct {
	DSN string
}

type MongoDB struct {
	DB *mongo.Database
}

func NewMongoDB(opts MongoDBOptions) *MongoDB {
	bsonOpts := &options.BSONOptions{
		UseJSONStructTags:   true,
		NilSliceAsEmpty:     true,
		ObjectIDAsHexString: true,
	}
	clientOpts := options.Client().
		ApplyURI(opts.DSN).
		SetBSONOptions(bsonOpts)
	client, err := mongo.Connect(clientOpts)
	if err != nil {
		log.Fatalf("Failed to connect mongodb: %s", err)
	}
	database := client.Database("sturdy")
	return &MongoDB{DB: database}
}

func (m *MongoDB) Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := m.DB.Client().Disconnect(ctx); err != nil {
		log.Fatalf("MongoDB error in closing the connection: %s", err)
	}
}
