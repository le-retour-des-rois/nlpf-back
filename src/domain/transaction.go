package domain

import (
	"net/http"
)

type Transaction struct {
	//min_prix    int64
	//max_prix    int64
	nom_commune string
	//type_local  string
}

type TransactionServiceDomain interface {
	GetInfo(http.ResponseWriter, *http.Request)
}

type TransactionRepositoryDomain interface {
	//GetAll() string
	GetInfo(nom_commune string)
}
