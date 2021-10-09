package transaction

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionRepository struct {
	collection *mongo.Collection
}

func NewTransactionRepository(db *mongo.Database, collectionName string) domain.transactionRepositoryDomain {
	return &transactionRepository{
		collection: db.Collection(collectionName),
	}
}
