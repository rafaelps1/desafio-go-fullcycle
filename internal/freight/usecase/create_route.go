package usecase

import (
	"github.com/rafaelps1/imersao14/go/internal/freight/entity"
)

type CreateRouteInput struct {
	ID             string
	Name           string
	SourceLat      float64
	SourceLng      float64
	DestinationLat float64
	DestinationLng float64
}

type CreateRouteOutput struct {
	ID   string
	Name string
}

type CreateRouteUseCase struct {
	Repositoty entity.RouteRepository
}

func NewCreateRouteUseCase(repository entity.RouteRepository) *CreateRouteUseCase {
	return &CreateRouteUseCase{
		Repositoty: repository,
	}
}

func (uc *CreateRouteUseCase) Execute(input CreateRouteInput) (*CreateRouteOutput, error) {
	source := entity.NewCoord(input.SourceLat, input.SourceLng)
	destination := entity.NewCoord(input.DestinationLat, input.DestinationLng)
	route := entity.NewRoute(input.ID, input.Name, source, destination)
	err := uc.Repositoty.Create(route)
	if err != nil {
		return nil, err
	}

	return &CreateRouteOutput{
		ID:   route.ID,
		Name: route.Name,
	}, nil
}
