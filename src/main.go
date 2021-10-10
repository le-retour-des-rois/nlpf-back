package main

import (
	//"encoding/json"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	// -- Import Packages ---

	transaction "test/transaction"
)

func connectDB() (db *mongo.Database) {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

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
	return client.Database("test-db")
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func main() {
	db := connectDB()

	//parser()

	// --- Router Instanciation --- //
	mainRouter := mux.NewRouter().StrictSlash(true)

	// --- Transactions --- //
	transactionRepository := transaction.NewTransactionRepository(db, "products")
	transactionService := transaction.NewTransactionService(transactionRepository)

	// --- RealEstate --- //
	//projectRepository := realEstate.NewRealEstateProjectRepository(db, "project")
	//projectService := realEstate.NewRealEstateProjectService(projectRepository)

	mainRouter.HandleFunc("/", HomePage).Methods("GET")
	mainRouter.HandleFunc("/transaction", transactionService.GetInfo).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", mainRouter))

}
