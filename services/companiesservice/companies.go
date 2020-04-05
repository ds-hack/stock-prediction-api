package companiesservice

import (
	"log"
	"time"

	"github.com/ds-hack/stock-prediction-api/database"
	"github.com/ds-hack/stock-prediction-api/models"
)

// GetAll Dashboardアプリに登録されている企業情報一覧を取得
func GetAll() (*models.CompanyWrapper, error) {
	var (
		companyWrapper *models.CompanyWrapper
		companies      []models.Company
		companiesDB    []*database.Company
	)

	companiesDB, err := database.GetCompanies()
	if err != nil {
		log.Printf("database fetch err: %v", err)
		return nil, err
	}

	for _, companydb := range companiesDB {
		company := models.Company{
			CompanyID:      companydb.CompanyID,
			CompanyName:    companydb.CompanyName,
			StockCode:      companydb.StockCode,
			CountryCode:    companydb.CountryCode,
			ListedMarket:   companydb.ListedMarket,
			FoundationDate: companydb.FoundationDate,
			Longitude:      companydb.Longitude,
			Latitude:       companydb.Latitude,
		}
		companies = append(companies, company)
	}
	companyWrapper = &models.CompanyWrapper{
		Timestamp: time.Now(),
		Companies: companies,
	}

	return companyWrapper, nil
}

// GetCompanyByStockCode CompaniesテーブルからStockCodeをキーとして企業情報を取得する
func GetCompanyByStockCode(stockCode string) (*models.CompanyDetail, error) {
	var (
		companyDetail *models.CompanyDetail
		companyDB     database.Company
	)

	companyDB, err := database.GetCompanyByStockCode(stockCode)
	if err != nil {
		log.Printf("database fetch err: %v", err)
		return nil, err
	}

	companyDetail = &models.CompanyDetail{
		Timestamp:      time.Now(),
		CompanyID:      companyDB.CompanyID,
		CompanyName:    companyDB.CompanyName,
		StockCode:      companyDB.StockCode,
		CountryCode:    companyDB.CountryCode,
		ListedMarket:   companyDB.ListedMarket,
		FoundationDate: companyDB.FoundationDate,
		Longitude:      companyDB.Longitude,
		Latitude:       companyDB.Latitude,
	}

	return companyDetail, nil
}
