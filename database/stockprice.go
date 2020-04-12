package database

import (
	"github.com/jinzhu/gorm"
)

// StockPrice stockpriceテーブルのエンティティモデル
type StockPrice struct {
	Model
	// システム内で設定する企業毎に一意となるコード
	CompanyID string `gorm:"column:company_id;type:varchar(16);not null;primary_key"`

	// 日付
	Date string `gorm:"column:date;type:date;not null;primary_key"`

	// 株価(始値)
	OpenPrice float32 `gorm:"column:open_price;type:number"`

	// 株価(高値)
	HighPrice float32 `gorm:"column:high_price;type:number"`

	// 株価(安値)
	LowPrice float32 `gorm:"column:low_price;type:number"`

	// 株価(終値)
	ClosePrice float32 `gorm:"column:close_price;type:number"`

	// 出来高
	Volume float32 `gorm:"column:volume;type:number"`
}

// StockPriceMA stockprice_maテーブルのエンティティモデル
type StockPriceMA struct {
	Model
	// システム内で設定する企業毎に一意となるコード
	CompanyID string `gorm:"column:company_id;type:varchar(16);not null;primary_key"`

	// 日付
	Date string `gorm:"column:date;type:date;not null;primary_key"`

	// 移動平均の計算条件(ex. sma5, ema25, ...)
	MAType string `gorm:"column:ma_type;type:varchar(16);not null;primary_key"`

	// ma_typeの条件下の移動平均株価
	MAValue float32 `gorm:"column:ma_value;type:number"`
}

// GetStockPrices stockpriceテーブルから株価を取得する
func GetStockPrices(companyID, startDate, endDate string) ([]*StockPrice, error) {
	var stockPrices []*StockPrice

	err := db.Table("stockprice").
		Where("company_id = ? AND date BETWEEN ? AND ?", companyID, startDate, endDate).
		Order("date").
		Find(&stockPrices).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return stockPrices, nil
}

// GetStockPricesMA stockprice_maテーブルから移動平均株価を取得する
func GetStockPricesMA(companyID, startDate, endDate string, maTypes []string) ([]*StockPriceMA, error) {
	var stockPricesMA []*StockPriceMA

	err := db.Table("stockprice_ma").
		Where("company_id = ? AND date BETWEEN ? AND ? AND ma_type IN (?)",
			companyID, startDate, endDate, maTypes).
		Order("ma_type, date").
		Find(&stockPricesMA).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return stockPricesMA, nil
}
