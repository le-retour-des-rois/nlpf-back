package area

import (
	"net/http"
	"fmt"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"

	//"github.com/le-retour-des-rois/nlpf-back/tree/mongo-docker/src/domain"
	"domain" 	
)

type areaService struct {
	areaRepository 	domain.AreaRepository
}

func NewAreaService(ar domain.AreaRepository) domain.AreaService {
	return &areaService{
		areaRepository:    ar,
	}
}

/*func (r *AreaRepository) GetAll() ([]Area, error) {

}*/

func (as *areaService) GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, as.areaRepository.GetAll())
}