package domain

import (
	"net/http"
)

type Transaction struct {
	id            int64  `json:"id"`
	groundSurface int64  `json:"ground_surface" validate:"required"`
	buildingType  string `json:"building_type"`
	postalCode    string `json:"postal_code"`
	mutationType  string `json:"mutation_type"`
}

type TransactionServiceDomain interface {
	GetAll(http.ResponseWriter, *http.Request)
}

type TransactionRepositoryDomain interface {
	//GetAll() string
}
