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

func (as *RealEstateProjectRepository) GetAll() []domain.RealEstateProject {
	var temp []domain.RealEstateProject
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
		bsonBytes, _ := bson.Marshal(project)
		bson.Unmarshal(bsonBytes, &structure)
		temp = append(temp, structure)
	}
	fmt.Println(res)
	return temp
}

func (as *RealEstateProjectRepository) AddProject(project domain.RealEstateProject) {
	//fmt.Println(project)
	//res, err := as.collection.InsertOne(context.TODO(), project)
	as.collection.InsertOne(context.TODO(), project)
	//fmt.Println("res: ", *res)
	//fmt.Println("err: ", err)
}

func (as *RealEstateProjectRepository) GetOne(id string) domain.RealEstateProject {
	//TODO
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}
	fmt.Println(objectId)

	result := as.collection.FindOne(context.TODO(), bson.M{"_id": objectId})
	var project domain.RealEstateProject
	result.Decode(&project)
	// fmt.Println("project ", project)
	// fmt.Println("result ", result)
	return project
}

func (as *RealEstateProjectRepository) DeleteProject(id int64) {
	//TODO
}
