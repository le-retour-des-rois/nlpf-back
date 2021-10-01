package area

import (
	"context"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"	

	//"github.com/le-retour-des-rois/nlpf-back/tree/mongo-docker/src/domain"
	"domain"
)

type areaRepository struct {
	collection *mongo.Collection
}

func NewAreaRepository(db *mongo.Database, collectionName string) domain.AreaRepository {
	return &areaRepository{
		collection: db.Collection(collectionName),
	}
}

/*func (r *AreaRepository) GetAll() ([]Area, error) {

}*/

type Book struct {
	Title     string
	Author    string
	ISBN      string
	Publisher string
	Copies     int
  }

func (r *areaRepository) GetAll() string {
	book1 := Book{"Animal Farm3", "George Orwell", "0451526341", "Example", 100}
	r.collection.InsertOne(context.TODO(), book1)
	return "Coucou la team"
}