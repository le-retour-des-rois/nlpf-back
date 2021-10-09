package main

import (
	//"encoding/json"
	"context"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// -- Import Packages ---
)

func connectDB() (db *mongo.Database) {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://root:example@localhost:27018")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client.Database("testdb")
}

func main() {
	db := connectDB()

	//parser()

	// --- Router Instanciation --- //
	mainRouter := mux.NewRouter().StrictSlash(true)

	// --- Transactions --- //
	transactionRepository := transaction.newTransactionRepository(db, "transactions")
	transactionService := transaction.newTransactionService(transactionRepository)

	// --- RealEstate --- //
	//projectRepository := domain.RealEstateProjectRepository(db, "project")
	//projectService := domain.RealEstateProjectService(projectRepository)

	mainRouter.HandleFunc("/areas", transactionService.GetAll).Methods("GET")
}
