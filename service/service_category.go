package service

import (
	"errors"

	"github.com/MuhammadSarimin/belajar-unit-test-golang.git/entity"
	"github.com/MuhammadSarimin/belajar-unit-test-golang.git/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (s CategoryService) Get(id string) (*entity.Category, error) {
	category := s.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("category not found")
	}
	return category, nil
}
