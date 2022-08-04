package service

import (
	"testing"

	"github.com/MuhammadSarimin/belajar-unit-test-golang.git/entity"
	"github.com/MuhammadSarimin/belajar-unit-test-golang.git/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repo = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var service = CategoryService{Repository: repo}

func TestCategoryService_GetNotFound(t *testing.T) {

	repo.Mock.On("FindById", "1").Return(nil)
	category, err := service.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)

}

func TestCategoryService_GetSuccess(t *testing.T) {

	category := entity.Category{
		Id:   "2",
		Name: "Sarimin",
	}

	repo.Mock.On("FindById", "2").Return(category)
	res, err := service.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, category.Id, res.Id)
	assert.Equal(t, category.Name, res.Name)
}
