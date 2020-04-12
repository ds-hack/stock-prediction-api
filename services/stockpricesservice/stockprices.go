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
		companyDB            database.Company
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

// GetMAPrices 指定の企業について、指定の計算方法の移動平均株価を取得する
func GetMAPrices(stockCode, startDate, endDate string, maTypes []string) (*models.StockPriceMAWrapper, error) {
	var (
		stockPriceMAWrapper *models.StockPriceMAWrapper
		stockPricesMA       []models.StockPriceMA
		companyDB           database.Company
		stockPricesMADB     []*database.StockPriceMA
		err                 error
	)

	// StockCodeにより企業情報を検索・取得する
	companyDB, err = database.GetCompanyByStockCode(stockCode)
	if err != nil {
		log.Printf("database fetch err: %v", err)
		return nil, err
	}

	// 企業ID・期間・計算方法を指定して、移動平均株価データを取得する
	stockPricesMADB, err = database.GetStockPricesMA(companyDB.CompanyID, startDate, endDate, maTypes)
	if err != nil {
		log.Printf("database fetch err: %v", err)
		return nil, err
	}

	for _, madb := range stockPricesMADB {
		ma := models.StockPriceMA{
			Date:    madb.Date,
			MAType:  madb.MAType,
			MAValue: madb.MAValue,
		}
		stockPricesMA = append(stockPricesMA, ma)
	}

	stockPriceMAWrapper = &models.StockPriceMAWrapper{
		Timestamp:   time.Now(),
		CompanyID:   companyDB.CompanyID,
		CompanyName: companyDB.CompanyName,
		StockCode:   companyDB.StockCode,
		Data:        stockPricesMA,
	}
	return stockPriceMAWrapper, nil
}
