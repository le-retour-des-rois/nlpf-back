package domain

import (
	"net/http"
)

type Transaction struct {
	Valeur_fonciere           int64
	Nom_commune               string
	Type_local                string
	Code_postal               int64
	Nombre_pieces_principales int64
	Surface_terrain           int64
	Code_departement          int64
}

type TransactionServiceDomain interface {
	GetInfo(http.ResponseWriter, *http.Request)
}

type TransactionRepositoryDomain interface {
	//GetAll() string
	GetInfo(nom_commune string, type_local string, min_prix int64, max_prix int64) []Transaction
}
