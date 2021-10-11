package transaction

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

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

/*func (repository *TransactionRepository) FindbyAll(min_prix int64, max_prix int64, nom_commune string, type_local string) *domain.Transaction {
	result := domain.Transaction{}
	session := repository.collection.Session(nil)
	coll := session.DB(os.Getenv("MONGODB_DATABASE")).C("user")
	err := coll.Find(bson.M{"_id": id}).One(&result)
	//err := repository.collection.Find(bson.M{"nom_commune": nom_commune}).One(&result)

	if err != nil {
		return nil, err
	}

	return &result
}*/

func (r *TransactionRepository) GetInfo(nom_commune string, type_local string, min_prix int64, max_prix int64) {
	opts := options.Find()
	opts.SetSort(bson.D{{"valeur_fonciere", -1}})
	filterCursor, err := r.collection.Find(context.TODO(),
		bson.D{
			{"nom_commune", nom_commune},
			{"type_local", type_local},
			{"valeur_fonciere", bson.D{
					{"$gt", min_prix},
				}},
			{"valeur_fonciere", bson.D{
					{"$lt", max_prix},
				}},
		}, opts)

	if err != nil {
		log.Fatal(err)
	}
	var transactionFiltered []bson.M
	if err = filterCursor.All(context.TODO(), &transactionFiltered); err != nil {
		log.Fatal(err)
	}
	fmt.Println(transactionFiltered)
}
