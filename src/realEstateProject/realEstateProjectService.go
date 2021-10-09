package realEstateProject

import (
	"net/http"

	domain "test/domain"
)

type RealEstateProjectService struct {
	RealEstateProjectRepository domain.RealEstateProjectRepositoryDomain
}

func NewRealEstateProjectService(ar domain.RealEstateProjectRepositoryDomain) domain.RealEstateProjectRepositoryDomain {
	return &RealEstateProjectService{
		RealEstateProjectRepository: ar,
	}
}

func (as *RealEstateProjectService) GetAll(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, as.RealEstateProjectRepository.GetAll())
}
