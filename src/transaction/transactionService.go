package transaction

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	domain "test/domain"
)

type TransactionService struct {
	TransactionRepository domain.TransactionRepositoryDomain
}

func NewTransactionService(ar domain.TransactionRepositoryDomain) domain.TransactionServiceDomain {
	return &TransactionService{
		TransactionRepository: ar,
	}
}

/*func (as *TransactionService) GetAll(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, as.TransactionRepository.GetAll())
}*/

type TransactionDTO struct {
	Min_prix    int64
	Max_prix    int64
	Nom_commune string
	Type_local  string
}

func (as *TransactionService) GetInfo(w http.ResponseWriter, r *http.Request) {

	var TransactionRequest TransactionDTO

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &TransactionRequest)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "%+v", as.TransactionRepository.GetInfo(TransactionRequest.Nom_commune, TransactionRequest.Type_local, TransactionRequest.Min_prix, TransactionRequest.Max_prix))
}
