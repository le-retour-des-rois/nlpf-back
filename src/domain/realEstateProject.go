package domain

import (
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Struct to put data or receive in the DB
type RealEstateProject struct {
	Id          primitive.ObjectID
	Max_prix    int64
	Min_prix    int64
	Nom_commune string
	Type_local  string
}

// Struct to send to the front
type RealEstateProjectBack struct {
	Id          string
	Max_prix    int64
	Min_prix    int64
	Nom_commune string
	Type_local  string
}

type RealEstateProjectServiceDomain interface {
	// Get all projects in the DB
	GetAll(http.ResponseWriter, *http.Request)
	// Get one project in the DB (int64)
	GetOne(http.ResponseWriter, *http.Request)
	// Delete one project in the DB (int64)
	DeleteProject(http.ResponseWriter, *http.Request)
	// Add one project in the DB (int64, int64, string, string)
	AddProject(http.ResponseWriter, *http.Request)
}

type RealEstateProjectRepositoryDomain interface {
	// Get all projects in the DB
	GetAll() []RealEstateProjectBack
	// Get one project in the DB (int64)
	GetOne(id string) []Transaction
	// Delete one project in the DB (int64)
	DeleteProject(id string)
	// Add one project in the DB (int64, int64, string, string)
	AddProject(RealEstateProject)
}
