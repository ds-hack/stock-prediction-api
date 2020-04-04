package stockpricesservice

import (
	"log"
	"time"

	"github.com/ds-hack/stock-prediction-api/database"
	"github.com/ds-hack/stock-prediction-api/models"
)

// GetRawPrices 指定の企業について、株価の生データ(open/high/low/close/volume)を取得する
func GetRawPrices(stockCode, startDate, endDate string) (*models.RawStockPriceWrapper, error) {
	var (
		rawStockPriceWrapper *models.RawStockPriceWrapper
		rawStockPrices       []models.RawStockPrice
		companyDB            *database.Company
		stockPricesDB        []*database.StockPrice
		err                  error
	)

	// StockCodeにより企業情報を検索・取得する
	companyDB, err = database.GetCompanyByStockCode(stockCode)
	if err != nil {
		log.Printf("database fetch err: %v", err)
		return nil, err
	}

	// 企業IDと期間を指定して、株価の生データを取得する
	stockPricesDB, err = database.GetStockPrices(companyDB.CompanyID, startDate, endDate)
	if err != nil {
		log.Printf("database fetch err: %v", err)
		return nil, err
	}

	for _, stockpricedb := range stockPricesDB {
		stockprice := models.RawStockPrice{
			Date:   stockpricedb.Date,
			Open:   stockpricedb.OpenPrice,
			High:   stockpricedb.HighPrice,
			Low:    stockpricedb.LowPrice,
			Close:  stockpricedb.ClosePrice,
			Volume: stockpricedb.Volume,
		}
		rawStockPrices = append(rawStockPrices, stockprice)
	}

	rawStockPriceWrapper = &models.RawStockPriceWrapper{
		Timestamp:   time.Now(),
		CompanyID:   companyDB.CompanyID,
		CompanyName: companyDB.CompanyName,
		StockCode:   companyDB.StockCode,
		Data:        rawStockPrices,
	}

	return rawStockPriceWrapper, nil
}
