package realEstateProject

import (
	"go.mongodb.org/mongo-driver/mongo"

	domain "test/domain"
)

type RealEstateProjectRepository struct {
	collection *mongo.Collection
}

func NewRealEstateProjectRepository(db *mongo.Database, collectionName string) domain.RealEstateProjectRepositoryDomain {
	return &RealEstateProjectRepository{
		collection: db.Collection(collectionName),
	}
}
