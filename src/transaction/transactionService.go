package transaction

import (
	"net/http"

	domain "test/domain"
)

type TransactionService struct {
	TransactionRepository domain.TransactionRepositoryDomain
}

func NewTransactionService(ar domain.TransactionRepositoryDomain) domain.TransactionRepositoryDomain {
	return &TransactionService{
		TransactionRepository: ar,
	}
}

func (as *TransactionService) GetAll(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, as.TransactionRepository.GetAll())
}
