package main

import (
	//"encoding/json"
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"encoding/csv"
	"io"
	"os"
	// -- Import Packages ---
	// "github.com/le-retour-des-rois/nlpf-back/tree/mongo-docker/src/api/area"
)

type Transaction struct {
	date_mutation             string `bson: "date_mutation,omitempty"`             //1
	nature_mutation           string `bson: "nature_mutation,omitempty"`           //3
	valeur_fonciere           string `bson: "valeur_fonciere,omitempty"`           //4
	code_postal               string `bson: "code_postal,omitempty"`               //9
	nom_commune               string `bson: "nom_commune,omitempty"`               //11
	code_departement          string `bson: "code_departement,omitempty"`          //12
	type_local                string `bson: "type_local,omitempty"`                //30
	nombre_pieces_principales string `bson: "nombre_pieces_principales,omitempty"` //32
	surface_terrain           string `bson: "surface_terrain,omitempty"`           //37
	longitude                 string `bson: "longitude,omitempty"`                 //38
	latitude                  string `bson: "latitude,omitempty"`                  //39
}

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

func parser() {

	db := connectDB()

	// Open the file
	csvfile, err := os.Open("../full.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	reader := csv.NewReader(csvfile)
	//reader := csv.NewReader(bufio.NewReader(csvfile))

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		newRecord := Transaction{
			date_mutation:             record[1],
			nature_mutation:           record[3],
			valeur_fonciere:           record[4],
			code_postal:               record[9],
			nom_commune:               record[11],
			code_departement:          record[12],
			type_local:                record[30],
			nombre_pieces_principales: record[32],
			surface_terrain:           record[37],
			longitude:                 record[38],
			latitude:                  record[39],
		}
		//fmt.Printf("%+v\n", newRecord)
		collection := db.Collection("Transaction")
		collection.InsertOne(context.TODO(), newRecord)
		//fmt.Printf("Insert of record : %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s\n", record[1], record[3], record[4], record[9], record[11], record[12], record[30], record[32], record[37], record[38], record[39])
	}
}

func main() {
	// db := connectDB()

	parser()

	// --- Router Instanciation --- //
	// mainRouter := mux.NewRouter().StrictSlash(true)

	// // --- AREAS --- //
	// areaRepository := area.NewAreaRepository(db, "books")
	// areaService := area.NewAreaService(areaRepository)

	// mainRouter.HandleFunc("/", homePage)
	// mainRouter.HandleFunc("/areas", areaService.GetAll).Methods("GET")
	//mainRouter.HandleFunc("/areas", postArticles).Methods("POST")
	// log.Fatal(http.ListenAndServe(":8001", mainRouter))
}
