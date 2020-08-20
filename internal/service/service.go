package service

import (
	"github.com/mokichi13/dynamo/internal/model"
	"github.com/mokichi13/dynamo/internal/repository"
)

type Server interface {
	Run() error
}

type server struct {
	repository.Repository
}

func (s *server) Run() error {
	return s.createData()
}

func NewService() (Server, error) {
	r, err := repository.NewRepository()
	if err != nil {
		return nil, err
	}
	return &server{
		Repository: r,
	}, nil
}

func (s *server) createData() error {
	m := testModel()
	return s.Repository.Put(m)
}

func testModel() model.Model {
	objID := "2"

	return model.Model{
		Object: &model.Object1{
			ID:   objID,
			Name: "hoge",
			Age:  25,
		},
		Details: []model.Detail{
			&model.Detail1{
				ID:     "detail1",
				Object: objID,
				Job:    "enginner",
			},
			&model.Detail2{
				ID:         "detail2",
				Object:     objID,
				MiddleName: "jack",
			},
		},
	}
}
