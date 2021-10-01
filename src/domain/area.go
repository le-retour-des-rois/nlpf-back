package domain

import (
	"net/http"
)

// Area ...
type Area struct {
	id        			int64    	`json:"id"`
	groundSurface     	int64    	`json:"ground_surface" validate:"required"`
	buildingType		string		`json:"building_type"`
	postalCode			string		`json:"postal_code"`
	mutationType		string		`json:"mutation_type"`
}

// AreaService defines all services related to areas
type AreaService interface {
	/*Fetch(ctx context.Context, cursor string, num int64) ([]Area, string, error)


	
	GetByID(ctx context.Context, id int64) (Area, error)
	GetByGroundSurface(ctx context.Context, groundSurface string) ([]Area, error)
	GetByBuildingType(ctx context.Context, buildingType string) ([]Area, error)
	GetByPostalCode(ctx context.Context, postalCode string) ([]Area, error)*/
	//GetAll() ([]Area, error)
	GetAll(http.ResponseWriter, *http.Request)
}

// AreaRepository defines all repositories related to areas
type AreaRepository interface {
	/*Fetch(ctx context.Context, cursor string, num int64) (res []Area, nextCursor string, err error)

	GetByID(ctx context.Context, id int64) Area
	GetByGroundSurface(ctx context.Context, groundSurface string) []Area
	GetByBuildingType(ctx context.Context, buildingType string) []Area
	GetByPostalCode(ctx context.Context, postalCode string) []Area*/
	GetAll() string

}