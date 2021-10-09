package area

import (
	"fmt"
	"net/http"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"

	//"github.com/le-retour-des-rois/nlpf-back/tree/mongo-docker/src/domain"
	"domain"
)

type transactionService struct {
	transactionRepository domain.TransactionRepository
}

func NewAreaService(ar domain.TransactionRepository) domain.TransactionService {
	return &transactionService{
		transactionRepository: ar,
	}
}

/*func (r *AreaRepository) GetAll() ([]Area, error) {

}*/

func (as *transactionService) GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, as.transactionRepository.GetAll())
}
