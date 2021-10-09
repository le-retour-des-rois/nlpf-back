/*package realEstateProject

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database, collectionName string) domain.RepositoryDomain {
	return &Repository{
		collection: db.Collection(collectionName),
	}
}*/