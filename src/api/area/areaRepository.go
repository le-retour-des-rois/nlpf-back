package area

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/le-retour-des-rois/nlpf-back/tree/mongo-docker/src/domain"
)

type areaRepository struct {
	collection *mongo.Collection
}

func NewAreaRepository(db *mongo.Database) (domain.AreaRepository, error) {
	return &AreaRepository{collection: db.Collection(collectionName), hash: encrypt.Hash{}}, nil
}

/*func (r *AreaRepository) GetAll() ([]Area, error) {

}*/

func (r *areaRepository) GetAll() (string, error) {
	return "Coucou la team"
}