package service

import (
	"final-project/dto/request"
	"final-project/internal/domain"
	"final-project/internal/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(r *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: r}
}

func (s *CategoryService) Create(input request.CreateCategoryRequest) error {
	category := domain.Category{
		Name: input.Name,
	}
	return s.repo.Create(&category)
}

func (s *CategoryService) GetAll() ([]domain.Category, error) {
	return s.repo.FindAll()
}

func (s *CategoryService) Update(id uint, input request.UpdateCategoryRequest) error {
	category, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	category.Name = input.Name
	return s.repo.Update(category)
}

func (s *CategoryService) Delete(id uint) error {
	category, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(category)
}
