package domain

import (
	"net/http"
)

type RealEstateProject struct {
	Id          int64
	Max_prix    int64
	Min_prix    int64
	Nom_commune string
	Type_local  string
}

type RealEstateProjectServiceDomain interface {
	// Get all projects in the DB
	GetAll(http.ResponseWriter, *http.Request)
	// Get one project in the DB (int64)
	//GetAll() string
	// Delete one project in the DB (int64)
	//GetAll() string
	// Add one project in the DB (int64, int64, string, string)
	AddProject(http.ResponseWriter, *http.Request)
}

type RealEstateProjectRepositoryDomain interface {
	// Get all projects in the DB
	GetAll() []RealEstateProject
	// Get one project in the DB (int64)
	//GetAll() string
	// Delete one project in the DB (int64)
	//GetAll() string
	// Add one project in the DB (int64, int64, string, string)
	AddProject(RealEstateProject)
}
