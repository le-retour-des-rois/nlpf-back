package realEstateProject

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strconv"
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
	// w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	// w.Header().Set("Content-Type", "application/json")
	// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
	nom_comm := r.FormValue("Nom_commune")
	max_prix, _ := strconv.ParseInt(r.FormValue("Max_prix"), 0, 64)
	min_prix, _ := strconv.ParseInt(r.FormValue("Min_prix"), 0, 64)
	type_local := r.FormValue("Type_local")

	var project domain.RealEstateProject
	// err := json.NewDecoder(r.Body).Decode(&project)
	project.Id = primitive.NewObjectID()
	project.Nom_commune = nom_comm
	project.Min_prix = min_prix
	project.Max_prix = max_prix
	project.Type_local = type_local

	as.RealEstateProjectRepository.AddProject(project)
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
	var project = as.RealEstateProjectRepository.GetOne(temp)
	j, _ := json.Marshal(project)
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
