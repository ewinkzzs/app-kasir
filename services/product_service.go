package services

import (
	"app-kasir/models"
	"app-kasir/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAll(name string) ([]models.ProductResponse, error) {
	return s.repo.GetAll(name)
}

func (s *ProductService) GetByID(id int) (*models.ProductResponse, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) Create(p *models.Product) error {
	return s.repo.Create(p)
}

func (s *ProductService) Update(p *models.Product) error {
	return s.repo.Update(p)
}

func (s *ProductService) Delete(id int) error {
	return s.repo.Delete(id)
}
