package main

import (
	"context"
	"log"
	"net/http"

	"github.com/HMasataka/ogen/ogen"
)

type petService struct {
	ogen.UnimplementedHandler
}

func (n *petService) CreatePets(ctx context.Context, req *ogen.Pet) error {
	return nil
}

func (n *petService) ListPets(ctx context.Context, params ogen.ListPetsParams) (*ogen.PetsHeaders, error) {
	var pet ogen.PetsHeaders
	return &pet, nil
}

func (n *petService) ShowPetById(ctx context.Context, params ogen.ShowPetByIdParams) (*ogen.Pet, error) {
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

func main() {
	service := &petService{}

	s, err := ogen.NewServer(service)
	if err != nil {
		log.Fatalln(err)
	}

	if err := http.ListenAndServe(":8080", s); err != nil {
		log.Fatalln(err)
	}
}
