package bd

import "go.mongodb.org/mongo-driver/mongo"

type BdUSer struct {
	db *mongo.Database
}

func newBbMongo(db *mongo.Database) *BdUSer {
	return &BdUSer{db}
}
