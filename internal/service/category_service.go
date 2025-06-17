package service

import (
	"final-project/internal/domain"
	"final-project/internal/repository"
)

type CategoryService interface {
	Create(name string) error
	GetAll() ([]domain.Category, error)
	Update(id uint, name string) error
	Delete(id uint) error
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo}
}

func (s *categoryService) Create(name string) error {
	category := &domain.Category{Name: name}
	return s.repo.Create(category)
}

func (s *categoryService) GetAll() ([]domain.Category, error) {
	return s.repo.FindAll()
}

func (s *categoryService) Update(id uint, name string) error {
	cat, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	cat.Name = name
	return s.repo.Update(cat)
}

func (s *categoryService) Delete(id uint) error {
	cat, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(cat)
}
