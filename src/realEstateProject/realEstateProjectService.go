package realEstateProject

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"

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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")

	if r.Method == http.MethodPost {
		project := domain.RealEstateProject{}
		err := json.NewDecoder(r.Body).Decode(&project)
		project.Id = primitive.NewObjectID()
		fmt.Println("response struct:", project)
		if err != nil {
			fmt.Println(err)
		}
		as.RealEstateProjectRepository.AddProject(project)
	}
}

func (as *RealEstateProjectService) GetAll(w http.ResponseWriter, r *http.Request) {
	var arr = as.RealEstateProjectRepository.GetAll()
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Content-Type", "application/json")
	j, _ := json.Marshal(arr)
	fmt.Fprintf(w, "%+v\n", string(j))
}

func (as *RealEstateProjectService) GetOne(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	var temp = params["id"]
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Content-Type", "application/json")
	var transactions = as.RealEstateProjectRepository.GetOne(temp)
	j, _ := json.Marshal(transactions)
	fmt.Fprintf(w, "%+v\n", string(j))
}

func (as *RealEstateProjectService) DeleteProject(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	var temp = params["id"]
	fmt.Println("id", temp)
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Content-Type", "application/json")
	as.RealEstateProjectRepository.DeleteProject(temp)
}
