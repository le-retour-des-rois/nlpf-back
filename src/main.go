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
	realEstate "test/realEstate"
	transaction "test/transaction"
)

func connectToProject() (db *mongo.Database) {
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

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func main() {
	db := connectToProject()

	// --- Router Instanciation --- //
	mainRouter := mux.NewRouter().StrictSlash(true)

	// --- Transactions --- //
	transactionRepository := transaction.NewTransactionRepository(db, "products")
	transactionService := transaction.NewTransactionService(transactionRepository)

	// --- RealEstate --- //
	projectRepository := realEstate.NewRealEstateProjectRepository(db, "projectRE")
	projectService := realEstate.NewRealEstateProjectService(projectRepository)

	mainRouter.HandleFunc("/project", projectService.GetAll).Methods("GET")
	mainRouter.HandleFunc("/project", projectService.AddProject).Methods("POST", "OPTIONS")
	mainRouter.HandleFunc("/project/{id}", projectService.DeleteProject).Methods("DELETE")
	mainRouter.HandleFunc("/project/{id}", projectService.GetOne).Methods("GET")
	mainRouter.HandleFunc("/transaction", transactionService.GetInfo).Methods("GET")
	mainRouter.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8081", mainRouter))
}
