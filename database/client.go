package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// DbMongo Flag value for mongod db
	DbMongo = 0x12
	// DbPg Flag value for postgres db
	DbPg = 0x13
)

// Database Container
type Database struct {
	client *mongo.Client
}

// New Create a new Database with fresh connection
func New(uri string) Database {
	return Database{CreateClient(uri)}
}

// CreateClient Create a suitable client for db
func CreateClient(uri string) *mongo.Client {
	if uri == "" {
		uri = ""
	}
	_client, _err := mongo.NewClient(options.Client().ApplyURI(uri))

	if _err != nil {
		log.Fatal(_err)
	}
	return _client
}

// GetContext Get context from Client connection
func (db *Database) GetContext() context.Context {
	var ctx context.Context
	var cancel context.CancelFunc

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	err := db.client.Ping(context.TODO(), nil)
	if err != nil {
		db.client.Connect(ctx)
	} else {
		ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	}

	defer cancel()
	defer func() {
		if err := db.client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	return ctx
}

// GetCollection Get context from Client connection
// client Pointer to client connection
func (db *Database) GetCollection(collection string) *mongo.Collection {
	database := db.client.Database("interloper")
	col := database.Collection("dev")

	return col
}
