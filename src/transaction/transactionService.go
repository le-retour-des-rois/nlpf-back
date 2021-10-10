package transaction

import (
	"encoding/json"
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
	//min_prix    int64
	//max_prix    int64
	Nom_commune string

	//type_local  string
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

	as.TransactionRepository.GetInfo(TransactionRequest.Nom_commune)
}
