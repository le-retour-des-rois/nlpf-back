package transaction

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
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

func (r *TransactionRepository) GetInfo(nom_commune string, type_local string, min_prix int64, max_prix int64) []domain.Transaction {

	var res []domain.Transaction

	filterCursor, err := r.collection.Find(context.TODO(), bson.M{
		"$and": []bson.M{
			bson.M{"type_local": type_local},
			bson.M{"nom_commune": nom_commune},
			bson.M{"valeur_fonciere": bson.M{"$gte": min_prix}},
			bson.M{"valeur_fonciere": bson.M{"$lte": max_prix}},
		}})

	if err != nil {
		log.Fatal(err)
	}
	var transactionFiltered []bson.M
	if err = filterCursor.All(context.TODO(), &transactionFiltered); err != nil {
		log.Fatal(err)
	}

	for _, transaction := range transactionFiltered {
		var structure domain.Transaction
		//bson.Unmarshal(project, &structure)
		bsonBytes, _ := bson.Marshal(transaction)
		bson.Unmarshal(bsonBytes, &structure)
		res = append(res, structure)
	}

	fmt.Println(transactionFiltered)

	return res
}
