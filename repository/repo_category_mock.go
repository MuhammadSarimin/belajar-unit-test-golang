package repository

import (
	"github.com/MuhammadSarimin/belajar-unit-test-golang.git/entity"
	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (r *CategoryRepositoryMock) FindById(id string) *entity.Category {
	args := r.Mock.Called(id)

	if args.Get(0) == nil {
		return nil
	}

	category := args.Get(0).(entity.Category)
	return &category
}
