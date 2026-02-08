package repositories

import (
	"app-kasir/models"
	"database/sql"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (repo *ReportRepository) GetSalesSummaryToday() (*models.SalesSummary, error) {
	var summary models.SalesSummary

	// total revenue & total transaksi
	err := repo.db.QueryRow(`
		SELECT 
			COALESCE(SUM(total_amount), 0),
			COUNT(*)
		FROM transactions
		WHERE DATE(created_at) = CURRENT_DATE
	`).Scan(&summary.TotalRevenue, &summary.TotalTransaksi)
	if err != nil {
		return nil, err
	}

	// produk terlaris
	err = repo.db.QueryRow(`
		SELECT p.name, COALESCE(SUM(td.quantity), 0) as qty
		FROM transaction_details td
		JOIN products p ON p.id = td.product_id
		JOIN transactions t ON t.id = td.transaction_id
		WHERE DATE(t.created_at) = CURRENT_DATE
		GROUP BY p.name
		ORDER BY qty DESC
		LIMIT 1
	`).Scan(&summary.ProdukTerlaris.Nama, &summary.ProdukTerlaris.QtyTerjual)

	// kalau belum ada transaksi hari ini
	if err == sql.ErrNoRows {
		summary.ProdukTerlaris = models.BestProduct{}
		return &summary, nil
	}
	if err != nil {
		return nil, err
	}

	return &summary, nil
}

func (repo *ReportRepository) GetSalesSummaryByDate(
	startDate, endDate string,
) (*models.SalesSummary, error) {

	var summary models.SalesSummary

	// total revenue & total transaksi
	err := repo.db.QueryRow(`
		SELECT 
			COALESCE(SUM(total_amount), 0),
			COUNT(*)
		FROM transactions
		WHERE DATE(created_at) BETWEEN $1 AND $2
	`, startDate, endDate).
		Scan(&summary.TotalRevenue, &summary.TotalTransaksi)

	if err != nil {
		return nil, err
	}

	// produk terlaris
	err = repo.db.QueryRow(`
		SELECT p.name, COALESCE(SUM(td.quantity), 0) AS qty
		FROM transaction_details td
		JOIN products p ON p.id = td.product_id
		JOIN transactions t ON t.id = td.transaction_id
		WHERE DATE(t.created_at) BETWEEN $1 AND $2
		GROUP BY p.name
		ORDER BY qty DESC
		LIMIT 1
	`, startDate, endDate).
		Scan(&summary.ProdukTerlaris.Nama, &summary.ProdukTerlaris.QtyTerjual)

	if err == sql.ErrNoRows {
		summary.ProdukTerlaris = models.BestProduct{}
		return &summary, nil
	}
	if err != nil {
		return nil, err
	}

	return &summary, nil
}
