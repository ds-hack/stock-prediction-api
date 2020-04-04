package database

import (
	"github.com/jinzhu/gorm"
)

// StockPrice StockPriceテーブルのエンティティモデル
type StockPrice struct {
	Model
	// システム内で設定する企業毎に一意となるコード
	CompanyID string `gorm:"column:company_id;type:varchar(16);not null;primary_key"`

	// 日付
	Date string `gorm:"column:date;type:date;not null"`

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

// GetStockPrices StockPricesテーブルから株価を取得する
func GetStockPrices(companyID, startDate, endDate string) ([]*StockPrice, error) {
	var stockprices []*StockPrice

	err := db.Table("stockprice").
		Where("company_id = ? AND date BETWEEN ? AND ?", companyID, startDate, endDate).
		Order("date").
		Find(&stockprices).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return stockprices, nil
}
