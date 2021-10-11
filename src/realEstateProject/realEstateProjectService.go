package realEstateProject

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	domain "test/domain"
)

type RealEstateProjectService struct {
	RealEstateProjectRepository domain.RealEstateProjectRepositoryDomain
}

func NewRealEstateProjectService(ar domain.RealEstateProjectRepositoryDomain) domain.RealEstateProjectServiceDomain {
	return &RealEstateProjectService{
		RealEstateProjectRepository: ar,
	}
}

func (as *RealEstateProjectService) AddProject(w http.ResponseWriter, r *http.Request) {
	var project domain.RealEstateProject
	err := json.NewDecoder(r.Body).Decode(&project)
	project.Id = primitive.NewObjectID()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	as.RealEstateProjectRepository.AddProject(project)
}

func (as *RealEstateProjectService) GetAll(w http.ResponseWriter, r *http.Request) {
	var arr = as.RealEstateProjectRepository.GetAll()
	for _, project := range arr {
		fmt.Fprintf(w, "%+v\n", project)
	}
}

func (as *RealEstateProjectService) GetOne(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	var temp = params["id"]
	var project = as.RealEstateProjectRepository.GetOne(temp)
	fmt.Fprintf(w, "%+v\n", project)
}

func (as *RealEstateProjectService) DeleteProject(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	var temp = params["id"]
	fmt.Println("id", temp)
	as.RealEstateProjectRepository.DeleteProject(temp)
}
