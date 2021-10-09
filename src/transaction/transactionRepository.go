package transaction

import (
	"go.mongodb.org/mongo-driver/mongo"

	domain "test/domain"
)

type TransactionRepository struct {
	collection *mongo.Collection
}

func NewTransactionRepository(db *mongo.Database, collectionName string) domain.TransactionRepositoryDomain {
	return &TransactionRepository{
		collection: db.Collection(collectionName),
	}
}
