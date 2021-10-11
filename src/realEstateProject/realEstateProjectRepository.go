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
	collection *mongo.Collection
}

func NewRealEstateProjectRepository(db *mongo.Database, collectionName string) domain.RealEstateProjectRepositoryDomain {
	return &RealEstateProjectRepository{
		collection: db.Collection(collectionName),
	}
}

func (as *RealEstateProjectRepository) AddProject(project domain.RealEstateProject) {
	as.collection.InsertOne(context.TODO(), project)
}

func (as *RealEstateProjectRepository) GetAll() []domain.RealEstateProjectBack {
	var temp []domain.RealEstateProjectBack
	filterCursor, err := as.collection.Find(context.TODO(), bson.M{})

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

func (as *RealEstateProjectRepository) GetOne(id string) domain.RealEstateProjectBack {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}

	result := as.collection.FindOne(context.TODO(), bson.M{"id": objectId})
	var project domain.RealEstateProject
	var projectFinal domain.RealEstateProjectBack
	result.Decode(&project)

	var tempID = project.Id.Hex()
	projectFinal.Id = tempID
	projectFinal.Max_prix = project.Max_prix
	projectFinal.Min_prix = project.Min_prix
	projectFinal.Nom_commune = project.Nom_commune
	projectFinal.Type_local = project.Type_local
	return projectFinal
}

func (as *RealEstateProjectRepository) DeleteProject(id string) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}
	fmt.Println(objectId)

	result, err := as.collection.DeleteOne(context.TODO(), bson.M{"id": objectId})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)
}
