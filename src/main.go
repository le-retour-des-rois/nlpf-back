package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	// -- Import Packages ---
	"github.com/le-retour-des-rois/nlpf-back/tree/mongo-docker/src/api/area"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

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

	// --- Router Instanciation --- //
	mainRouter := mux.NewRouter().StrictSlash(true)

	// --- AREAS --- //
	areaRepository 	:= area.NewAreaRepository(db)
	areaService 	:= area.NewAreaService(areaRepository)


	mainRouter.HandleFunc("/", homePage)
	mainRouter.HandleFunc("/areas", areaService.GetAll).Methods("GET")
	//mainRouter.HandleFunc("/areas", postArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8001", mainRouter))
}