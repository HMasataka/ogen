package api

import (
	"context"
	"net/http"

	"github.com/HMasataka/ogen/ogen"
)

type PetService struct {
	ogen.UnimplementedHandler
}

func (n *PetService) CreatePets(ctx context.Context, req *ogen.Pet) error {
	return nil
}

func (n *PetService) ListPets(ctx context.Context, params ogen.ListPetsParams) (*ogen.PetsHeaders, error) {
	var pet ogen.PetsHeaders
	return &pet, nil
}

func (n *PetService) ShowPetById(ctx context.Context, params ogen.ShowPetByIdParams) (*ogen.Pet, error) {
	err := ogen.Error{
		Code:    1,
		Message: "error",
	}

	errorStatusCode := ogen.ErrorStatusCode{
		StatusCode: http.StatusBadRequest,
		Response:   err,
	}

	return nil, &errorStatusCode
}
