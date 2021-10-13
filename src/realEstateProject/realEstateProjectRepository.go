package realEstateProject

import (
	"context"

	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"log"

	domain "test/domain"
)

type RealEstateProjectRepository struct {
	projectCollection     *mongo.Collection
	transactionCollection *mongo.Collection
}

func NewRealEstateProjectRepository(db *mongo.Database, projectCollectionName string, transactionCollectionName string) domain.RealEstateProjectRepositoryDomain {
	return &RealEstateProjectRepository{
		projectCollection:     db.Collection(projectCollectionName),
		transactionCollection: db.Collection(transactionCollectionName),
	}
}

func (as *RealEstateProjectRepository) AddProject(project domain.RealEstateProject) {
	as.projectCollection.InsertOne(context.TODO(), project)
}

func (as *RealEstateProjectRepository) GetAll() []domain.RealEstateProjectBack {
	var temp []domain.RealEstateProjectBack
	filterCursor, err := as.projectCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}
	var res []bson.M
	if err = filterCursor.All(context.TODO(), &res); err != nil {
		log.Fatal(err)
	}

	for _, project := range res {
		var structure domain.RealEstateProject
		var structureFinal domain.RealEstateProjectBack
		// Transform the project receive from the DB in the right struct
		bsonBytes, _ := bson.Marshal(project)
		bson.Unmarshal(bsonBytes, &structure)

		var tempID = structure.Id.Hex()
		structureFinal.Id = tempID
		structureFinal.Max_prix = structure.Max_prix
		structureFinal.Min_prix = structure.Min_prix
		structureFinal.Nom_commune = structure.Nom_commune
		structureFinal.Type_local = structure.Type_local
		temp = append(temp, structureFinal)
	}

	return temp
}

func (as *RealEstateProjectRepository) GetOne(id string) []domain.Transaction {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}

	result := as.projectCollection.FindOne(context.TODO(), bson.M{"id": objectId})
	var project domain.RealEstateProject
	result.Decode(&project)

	fmt.Println("toto: ", project)

	var res []domain.Transaction

	filterCursor, err := as.transactionCollection.Find(context.TODO(), bson.M{
		"$and": []bson.M{
			bson.M{"type_local": project.Type_local},
			bson.M{"nom_commune": project.Nom_commune},
			bson.M{"valeur_fonciere": bson.M{"$gte": project.Min_prix}},
			bson.M{"valeur_fonciere": bson.M{"$lte": project.Max_prix}},
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
		bsonBytes, _ := bson.Marshal(transaction)
		bson.Unmarshal(bsonBytes, &structure)
		res = append(res, structure)
	}

	fmt.Println("transactionFiltered ", transactionFiltered)
	fmt.Println("res ", res)

	return res
}

func (as *RealEstateProjectRepository) DeleteProject(id string) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}
	fmt.Println(objectId)

	result, err := as.projectCollection.DeleteOne(context.TODO(), bson.M{"id": objectId})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)
}
