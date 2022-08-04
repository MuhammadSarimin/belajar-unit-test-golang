package repository

import "github.com/MuhammadSarimin/belajar-unit-test-golang.git/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
