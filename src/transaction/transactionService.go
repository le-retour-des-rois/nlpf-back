package transaction

import (
	"net/http"

	"domain"
)

type TransactionService struct {
	transactionRepository domain.transactionRepositoryDomain
}

func NewTransactionService(ar domain.transactionRepositoryDomain) domain.transactionRepositoryDomain {
	return &transactionService{
		transactionRepository: ar,
	}
}

func (as *transactionService) GetAll(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, as.transactionRepository.GetAll())
}
