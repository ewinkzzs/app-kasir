package services

import (
	"app-kasir/models"
	"app-kasir/repositories"
)

type ReportService struct {
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetSalesSummaryToday() (*models.SalesSummary, error) {
	return s.repo.GetSalesSummaryToday()
}

func (s *ReportService) GetSalesSummaryByDate(
	startDate, endDate string,
) (*models.SalesSummary, error) {
	return s.repo.GetSalesSummaryByDate(startDate, endDate)
}
